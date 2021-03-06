package core

import (
	"context"
	walletbase "github.com/cpacia/multiwallet/base"
	"github.com/cpacia/openbazaar3.0/database"
	"github.com/cpacia/openbazaar3.0/events"
	"github.com/cpacia/openbazaar3.0/models"
	"github.com/cpacia/openbazaar3.0/models/factory"
	"github.com/cpacia/openbazaar3.0/net"
	"github.com/cpacia/openbazaar3.0/orders/pb"
	"github.com/cpacia/openbazaar3.0/orders/utils"
	iwallet "github.com/cpacia/wallet-interface"
	"testing"
	"time"
)

func TestOpenBazaarNode_CancelOrder(t *testing.T) {
	network, err := NewMocknet(2)
	if err != nil {
		t.Fatal(err)
	}

	defer network.TearDown()

	go network.StartWalletNetwork()

	for _, node := range network.Nodes() {
		go node.orderProcessor.Start()
	}

	orderSub0, err := network.Nodes()[0].eventBus.Subscribe(&events.NewOrder{})
	if err != nil {
		t.Fatal(err)
	}
	orderAckSub0, err := network.Nodes()[1].eventBus.Subscribe(&events.MessageACK{})
	if err != nil {
		t.Fatal(err)
	}

	listing := factory.NewPhysicalListing("tshirt")

	done := make(chan struct{})
	if err := network.Nodes()[0].SaveListing(listing, done); err != nil {
		t.Fatal(err)
	}
	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("Timeout waiting on channel")
	}

	index, err := network.Nodes()[0].GetMyListings()
	if err != nil {
		t.Fatal(err)
	}

	purchase := factory.NewPurchase()
	purchase.Items[0].ListingHash = index[0].CID
	if err := network.Nodes()[1].PingNode(context.Background(), network.Nodes()[0].Identity()); err != nil {
		t.Fatal(err)
	}

	// Cancelable order
	// We're going to disconnect the nodes, make the purchase, and then reconnect. This should cause node 1
	// to resend the order upon reconnection.
	network.Nodes()[0].networkService.Close()
	go network.Nodes()[1].syncMessages()
	if err := network.ipfsNet.UnlinkPeers(network.Nodes()[0].Identity(), network.Nodes()[1].Identity()); err != nil {
		t.Fatal(err)
	}
	orderID2, paymentAddress, paymentAmount, err := network.Nodes()[1].PurchaseListing(context.Background(), purchase)
	if err != nil {
		t.Fatal(err)
	}

	wallet1, err := network.Nodes()[1].multiwallet.WalletForCurrencyCode(iwallet.CtMock)
	if err != nil {
		t.Fatal(err)
	}

	txSub0, err := network.Nodes()[0].eventBus.Subscribe(&events.TransactionReceived{})
	if err != nil {
		t.Fatal(err)
	}

	txSub1, err := network.Nodes()[1].eventBus.Subscribe(&events.TransactionReceived{})
	if err != nil {
		t.Fatal(err)
	}

	addr, err := wallet1.CurrentAddress()
	if err != nil {
		t.Fatal(err)
	}

	if err := network.wn.GenerateToAddress(addr, iwallet.NewAmount(10000000000000)); err != nil {
		t.Fatal(err)
	}

	select {
	case <-txSub1.Out():
	case <-time.After(time.Second * 10):
		t.Fatal("Timeout waiting on channel")
	}

	// Reconnecting nodes should trigger node 1 to send the order to node 0 again.
	time.Sleep(1)
	network.Nodes()[0].networkService = net.NewNetworkService(network.Nodes()[0].ipfsNode.PeerHost, net.NewBanManager(nil), true)
	network.Nodes()[0].registerHandlers()

	if _, err := network.ipfsNet.LinkPeers(network.Nodes()[0].Identity(), network.Nodes()[1].Identity()); err != nil {
		t.Fatal(err)
	}

	select {
	case <-orderSub0.Out():
	case <-time.After(time.Second * 10):
		t.Fatal("Timeout waiting on channel")
	}

	select {
	case <-orderAckSub0.Out():
	case <-time.After(time.Second * 10):
		t.Fatal("Timeout waiting on channel")
	}

	wTx, err := wallet1.Begin()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := wallet1.Spend(wTx, paymentAddress, paymentAmount.Amount, iwallet.FlNormal); err != nil {
		t.Fatal(err)
	}

	if err := wTx.Commit(); err != nil {
		t.Fatal(err)
	}

	select {
	case <-txSub0.Out():
	case <-time.After(time.Second * 10):
		t.Fatal("Timeout waiting on channel")
	}

	select {
	case <-txSub1.Out():
	case <-time.After(time.Second * 10):
		t.Fatal("Timeout waiting on channel")
	}

	var order3 models.Order
	err = network.Nodes()[0].repo.DB().View(func(tx database.Tx) error {
		return tx.Read().Where("id = ?", orderID2.String()).Last(&order3).Error
	})
	if err != nil {
		t.Fatal(err)
	}

	if order3.SerializedOrderOpen == nil {
		t.Error("Node 0 failed to save order")
	}

	var order4 models.Order
	err = network.Nodes()[1].repo.DB().View(func(tx database.Tx) error {
		return tx.Read().Where("id = ?", orderID2.String()).Last(&order4).Error
	})
	if err != nil {
		t.Fatal(err)
	}

	if order4.SerializedOrderOpen == nil {
		t.Error("Node 1 failed to save order")
	}
	if !order4.OrderOpenAcked {
		t.Error("Node 1 failed to record order open ACK")
	}

	orderOpen, err := order4.OrderOpenMessage()
	if err != nil {
		t.Fatal(err)
	}

	if orderOpen.Payment.Method != pb.OrderOpen_Payment_CANCELABLE {
		t.Fatal("Expected CANCELABLE order")
	}

	releaseSub, err := network.Nodes()[1].eventBus.Subscribe(&events.SpendFromPaymentAddress{})
	if err != nil {
		t.Fatal(err)
	}

	done5 := make(chan struct{})
	if err := network.Nodes()[1].CancelOrder(orderID2, done5); err != nil {
		t.Fatal(err)
	}
	select {
	case <-done5:
	case <-time.After(time.Second * 10):
		t.Fatal("Timeout waiting on channel")
	}

	select {
	case <-releaseSub.Out():
	case <-time.After(time.Second * 10):
		t.Fatal("Timeout waiting on channel")
	}

	var order6 models.Order
	err = network.Nodes()[1].repo.DB().View(func(tx database.Tx) error {
		return tx.Read().Where("id = ?", orderID2.String()).Last(&order6).Error
	})
	if err != nil {
		t.Fatal(err)
	}

	if order6.SerializedOrderCancel == nil {
		t.Error("Node 0 failed to save order cancel")
	}

	txs, err := order6.GetTransactions()
	if err != nil {
		t.Fatal(err)
	}
	if len(txs) != 2 {
		t.Errorf("Expected 2 transactions, got %d", len(txs))
	}

	var order7 models.Order
	err = network.Nodes()[0].repo.DB().View(func(tx database.Tx) error {
		return tx.Read().Where("id = ?", orderID2.String()).Last(&order7).Error
	})
	if err != nil {
		t.Fatal(err)
	}

	if order7.SerializedOrderCancel == nil {
		t.Error("Node 0 failed to save order cancel")
	}
}

func TestOpenBazaarNode_releaseFromCancelableAddress(t *testing.T) {
	node, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}

	defer node.DestroyNode()

	orderOpen := &pb.OrderOpen{
		Payment: &pb.OrderOpen_Payment{
			Method:  pb.OrderOpen_Payment_CANCELABLE,
			Coin:    iwallet.CtMock,
			Address: "1234",
		},
	}

	order := new(models.Order)
	if err := order.PutMessage(utils.MustWrapOrderMessage(orderOpen)); err != nil {
		t.Fatal(err)
	}

	addr := iwallet.NewAddress("1234", iwallet.CtMock)
	tx := walletbase.NewMockTransaction(nil, &addr)
	if err := order.PutTransaction(tx); err != nil {
		t.Fatal(err)
	}

	wTx, txid, err := node.releaseFromCancelableAddress(order)
	if err != nil {
		t.Fatal(err)
	}

	if err := wTx.Commit(); err != nil {
		t.Fatal(err)
	}

	if txid == "" {
		t.Fatal("failed to returned a valid txid")
	}

	if err := order.PutTransaction(walletbase.NewMockTransaction(&tx.To[0], nil)); err != nil {
		t.Fatal(err)
	}

	_, _, err = node.releaseFromCancelableAddress(order)
	if err == nil {
		t.Fatal("Expected error spending non-existent coins")
	}
}
