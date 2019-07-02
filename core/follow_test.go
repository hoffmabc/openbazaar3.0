package core

import (
	"github.com/cpacia/openbazaar3.0/events"
	"github.com/cpacia/openbazaar3.0/models"
	peer "github.com/libp2p/go-libp2p-peer"
	"testing"
)

func TestOpenBazaarNode_Follow(t *testing.T) {
	node, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}
	defer node.repo.DestroyRepo()

	p, err := peer.IDB58Decode("QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub")
	if err != nil {
		t.Fatal(err)
	}

	if err := node.SetProfile(&models.Profile{Name: "Ron Paul"}, nil); err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	if err := node.FollowNode(p, done); err != nil {
		t.Fatal(err)
	}

	<-done

	following, err := node.GetMyFollowing()
	if err != nil {
		t.Fatal(err)
	}

	if following.Count() != 1 {
		t.Errorf("Incorrect number of following returned. Expected %d, %d", 1, following.Count())
	}

	if following[0] != p.String() {
		t.Errorf("Incorrect following peer returned. Expected %s, got %s", p.String(), following[0])
	}

	p2, err := peer.IDB58Decode("QmcUDmZK8PsPYWw5FRHKNZFjszm2K6e68BQSTpnJYUsML7")
	if err != nil {
		t.Fatal(err)
	}

	done = make(chan struct{})
	if err := node.FollowNode(p2, done); err != nil {
		t.Fatal(err)
	}

	<-done

	following, err = node.GetMyFollowing()
	if err != nil {
		t.Fatal(err)
	}

	if following.Count() != 2 {
		t.Errorf("Incorrect number of following returned. Expected %d, %d", 2, following.Count())
	}

	if following[1] != p2.String() {
		t.Errorf("Incorrect following peer returned. Expected %s, got %s", p.String(), following[1])
	}

	profile, err := node.GetMyProfile()
	if err != nil {
		t.Fatal(err)
	}

	if profile.Stats == nil {
		t.Fatal("Profile stats is nil")
	}

	if profile.Stats.FollowingCount != uint32(following.Count()) {
		t.Errorf("Following count in profile incorrect. Expected %d, got %d", following.Count(), profile.Stats.FollowingCount)
	}
}

func TestOpenBazaarNode_GetFollowing(t *testing.T) {
	mocknet, err := NewMocknet(2)
	if err != nil {
		t.Fatal(err)
	}
	defer mocknet.TearDown()

	p, err := peer.IDB58Decode("QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub")
	if err != nil {
		t.Fatal(err)
	}

	if err := mocknet.Nodes()[0].FollowNode(p, nil); err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	mocknet.Nodes()[0].Publish(done)
	<-done

	following, err := mocknet.Nodes()[1].GetFollowing(mocknet.Nodes()[0].Identity(), false)
	if err != nil {
		t.Fatal(err)
	}

	if following.Count() != 1 {
		t.Errorf("Incorrect number of following returned. Expected %d, %d", 1, following.Count())
	}

	if following[0] != p.String() {
		t.Errorf("Incorrect following peer returned. Expected %s, got %s", p.String(), following[0])
	}
}

func TestOpenBazaarNode_GetFollowers(t *testing.T) {
	mocknet, err := NewMocknet(2)
	if err != nil {
		t.Fatal(err)
	}
	defer mocknet.TearDown()

	p, err := peer.IDB58Decode("QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub")
	if err != nil {
		t.Fatal(err)
	}

	if err := mocknet.Nodes()[0].repo.PublicData().SetFollowers(models.Followers{p.String()}); err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})
	mocknet.Nodes()[0].Publish(done)
	<-done

	followers, err := mocknet.Nodes()[1].GetFollowers(mocknet.Nodes()[0].Identity(), false)
	if err != nil {
		t.Fatal(err)
	}

	if followers.Count() != 1 {
		t.Errorf("Incorrect number of following returned. Expected %d, %d", 1, followers.Count())
	}

	if followers[0] != p.String() {
		t.Errorf("Incorrect following peer returned. Expected %s, got %s", p.String(), followers[0])
	}
}

func Test_handleFollowAndUnfollow(t *testing.T) {
	mocknet, err := NewMocknet(2)
	if err != nil {
		t.Fatal(err)
	}
	defer mocknet.TearDown()

	// Test follow
	sub, err := mocknet.Nodes()[1].eventBus.Subscribe(&events.FollowNotification{})
	if err != nil {
		t.Fatal(err)
	}

	if err := mocknet.Nodes()[0].FollowNode(mocknet.Nodes()[1].Identity(), nil); err != nil {
		t.Fatal(err)
	}

	event := <-sub.Out()

	notif, ok := event.(*events.FollowNotification)
	if !ok {
		t.Fatalf("Event type assertion failed")
	}

	if notif.PeerId != mocknet.Nodes()[0].Identity().Pretty() {
		t.Errorf("Received incorrect peer ID in follow notification. Expected %s, got %s", mocknet.Nodes()[0].Identity().Pretty(), notif.PeerId)
	}

	followers, err := mocknet.Nodes()[1].GetMyFollowers()
	if err != nil {
		t.Fatal(err)
	}

	if followers.Count() != 1 {
		t.Fatalf("Incorrect number of followers returned. Expected %d, got %d", 1, followers.Count())
	}

	if followers[0] != mocknet.Nodes()[0].Identity().Pretty() {
		t.Errorf("Incorrect follower ID. Expected %s, got %s", mocknet.Nodes()[0].Identity().Pretty(), followers[0])
	}

	// Test unfollow
	sub2, err := mocknet.Nodes()[1].eventBus.Subscribe(&events.UnfollowNotification{})
	if err != nil {
		t.Fatal(err)
	}

	if err := mocknet.Nodes()[0].UnfollowNode(mocknet.Nodes()[1].Identity(), nil); err != nil {
		t.Fatal(err)
	}

	event2 := <-sub2.Out()

	notif2, ok := event2.(*events.UnfollowNotification)
	if !ok {
		t.Fatalf("Event type assertion failed")
	}

	if notif2.PeerId != mocknet.Nodes()[0].Identity().Pretty() {
		t.Errorf("Received incorrect peer ID in unfollow notification. Expected %s, got %s", mocknet.Nodes()[0].Identity().Pretty(), notif2.PeerId)
	}

	followers, err = mocknet.Nodes()[1].GetMyFollowers()
	if err != nil {
		t.Fatal(err)
	}

	if followers.Count() != 0 {
		t.Fatalf("Incorrect number of followers returned. Expected %d, got %d", 0, followers.Count())
	}
}
