package core

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/cpacia/multiwallet"
	"github.com/cpacia/openbazaar3.0/api"
	"github.com/cpacia/openbazaar3.0/core/coreiface"
	"github.com/cpacia/openbazaar3.0/events"
	"github.com/cpacia/openbazaar3.0/net"
	"github.com/cpacia/openbazaar3.0/notifications"
	"github.com/cpacia/openbazaar3.0/orders"
	"github.com/cpacia/openbazaar3.0/repo"
	"github.com/cpacia/openbazaar3.0/wallet"
	"github.com/ipfs/go-ipfs/core"
	peer "github.com/libp2p/go-libp2p-peer"
	"os"
	"sync/atomic"
	"time"
)

// OpenBazaarNode holds all the components that make up a network node
// on the OpenBazaar network. It also exposes an exported API which can
// be used to control the node.
type OpenBazaarNode struct {

	// ipfsNode is the IPFS instance that powers this node.
	ipfsNode *core.IpfsNode

	// repo holds the database and public data directory.
	repo *repo.Repo

	// escrowMasterKey represents an secp256k1 private key, the
	// public key of which is advertised by the node in its profile
	// and in listings to be used when building escrow transactions.
	escrowMasterKey *btcec.PrivateKey

	// ratingMasterKey represents an secp256k1 private key that
	// we used to generate rating keys to sign ratings with.
	ratingMasterKey *btcec.PrivateKey

	// ipnsQuorum is the size of the IPNS quorum to use. Smaller quorums
	// resolve faster but run the risk of getting back older records.
	ipnsQuorum uint

	// messenger is the primary object used to send messages to other peers.
	// It ensures reliable delivery by persisting messages and retrying them.
	// Generally you should always send messages using this and not the
	// NetworkService as the later will only attempt to send direct messages
	// and will not retry.
	messenger *net.Messenger

	// networkService manages the sending and receiving of messages
	// on the OpenBazaar protocol.
	networkService *net.NetworkService

	// banManager holds a list of peers that have been banned by this node.
	banManager *net.BanManager

	// eventBus allows a subscriber to receive event notifications from the node.
	eventBus events.Bus

	// followerTracker tries to maintain connections to a minimum number of our
	// followers so that we can use them to push data for redundancy.
	followerTracker *FollowerTracker

	// multiwallet is a map of cyptocurrency wallets.
	multiwallet multiwallet.Multiwallet

	// orderProcessor is the engine we use for processing all orders.
	orderProcessor *orders.OrderProcessor

	// exchangeRates is a provider of exchange rate data for various currencies.
	exchangeRates *wallet.ExchangeRateProvider

	// notifier listens to events coming off the bus, marshals them to notifications
	// and sends them off to the websocket.
	notifier *notifications.Notifier

	// gateway is the openbazaar API.
	gateway *api.Gateway

	// testnet is whether the this node is configured to use the test network.
	testnet bool

	// publishActive is an atomic integer that represents the number of inflight
	// publishes.
	publishActive int32

	// publishChan is used to signal to the republish loop that a publish
	// has just completed and it should update it's last published time.
	publishChan chan pubCloser

	// ipfsOnlyMode signals that the node is running in IPFS only mode.
	ipfsOnlyMode bool

	// shutdown is closed when the node is stopped. Any listening
	// goroutines can use this to terminate.
	shutdown chan struct{}
}

// Start gets the node up and running and listens for a signal interrupt.
func (n *OpenBazaarNode) Start() {
	if !n.ipfsOnlyMode {
		go n.messenger.Start()
		go n.followerTracker.Start()
		go n.orderProcessor.Start()
		go n.syncMessages()
		go n.publishHandler()
		go n.multiwallet.Start()
		go n.gateway.Serve()
		go n.notifier.Start()
		if err := n.removeDisabledCoinsFromListings(); err != nil && !os.IsNotExist(err) {
			log.Errorf("Error removing disabled coins from listings: %s", err)
		}
	}
}

// Stop cleanly shutsdown the OpenBazaarNode and signals to any
// listening goroutines that it's time to stop.
func (n *OpenBazaarNode) Stop(force bool) error {
	if atomic.LoadInt32(&n.publishActive) > 0 && !force {
		return coreiface.ErrPublishingActive
	}

	close(n.shutdown)
	if !n.ipfsOnlyMode {
		n.networkService.Close()
		n.messenger.Stop()
		n.orderProcessor.Stop()
		n.followerTracker.Close()
		if n.gateway != nil {
			n.gateway.Close()
		}
		n.multiwallet.Close()
	}
	n.repo.Close()

	stop := make(chan struct{})
	go func() {
		n.ipfsNode.Context().Done()
		n.ipfsNode.Close()
		n.eventBus.Emit(&events.IPFSShutdown{})
		close(stop)
	}()
	select {
	case <-time.After(time.Second * 2):
		return coreiface.ErrIPFSDelayedShutdown
	case <-stop:

	}
	return nil
}

// UsingTestnet returns whether or not this node is running on
// the test network.
func (n *OpenBazaarNode) UsingTestnet() bool {
	return n.testnet
}

// DestroyNode shutsdown the node and deletes the entire data directory.
// This should only be used during testing as destroying a live node will
// result in data loss.
func (n *OpenBazaarNode) DestroyNode() {
	n.Stop(true)
	n.repo.DestroyRepo()
}

// IPFSNode returns the underlying IPFS node instance.
func (n *OpenBazaarNode) IPFSNode() *core.IpfsNode {
	return n.ipfsNode
}

// Multiwallet returns the underlying Multiwallet instance.
func (n *OpenBazaarNode) Multiwallet() multiwallet.Multiwallet {
	return n.multiwallet
}

// Identity returns the peer ID for this node.
func (n *OpenBazaarNode) Identity() peer.ID {
	return n.ipfsNode.Identity
}

// SubscribeEvent returns a subscription to the provided event. The event argument
// may be an interface slice.
func (n *OpenBazaarNode) SubscribeEvent(event interface{}) (events.Subscription, error) {
	return n.eventBus.Subscribe(event)
}
