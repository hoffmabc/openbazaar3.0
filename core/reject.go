package core

import (
	"errors"
	"github.com/cpacia/openbazaar3.0/database"
	"github.com/cpacia/openbazaar3.0/models"
	npb "github.com/cpacia/openbazaar3.0/net/pb"
	"github.com/cpacia/openbazaar3.0/orders/pb"
	"github.com/golang/protobuf/ptypes"
)

// RejectOrder sends a ORDER_REJECT message to the remote peer and updates the node's
// order state. Only a vendor can call this method and only if the order has been opened
// and no other actions have been taken.
func (n *OpenBazaarNode) RejectOrder(orderID models.OrderID, reason string, done chan struct{}) error {
	var order models.Order
	err := n.repo.DB().View(func(tx database.Tx) error {
		return tx.Read().Where("id = ?", orderID.String()).First(&order).Error
	})
	if err != nil {
		return err
	}

	if !order.CanReject(n.Identity()) {
		return errors.New("order is not in a state where it can be rejected")
	}

	buyer, err := order.Buyer()
	if err != nil {
		return err
	}
	vendor, err := order.Vendor()
	if err != nil {
		return err
	}

	reject := pb.OrderReject{
		Type:   pb.OrderReject_USER_REJECT,
		Reason: reason,
	}

	rejectAny, err := ptypes.MarshalAny(&reject)
	if err != nil {
		return err
	}

	resp := npb.OrderMessage{
		OrderID:     order.ID.String(),
		MessageType: npb.OrderMessage_ORDER_REJECT,
		Message:     rejectAny,
	}

	payload, err := ptypes.MarshalAny(&resp)
	if err != nil {
		return err
	}

	message := newMessageWithID()
	message.MessageType = npb.Message_ORDER
	message.Payload = payload

	return n.repo.DB().Update(func(tx database.Tx) error {
		_, err := n.orderProcessor.ProcessMessage(tx, vendor, &resp)
		if err != nil {
			return err
		}

		return n.messenger.ReliablySendMessage(tx, buyer, message, done)
	})
}
