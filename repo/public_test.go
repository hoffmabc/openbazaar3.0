package repo

import (
	"context"
	"github.com/cpacia/openbazaar3.0/models"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/bootstrap"
	"github.com/ipfs/go-ipfs/core/mock"
	"github.com/ipfs/go-ipfs/namesys"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	mocknet "github.com/libp2p/go-libp2p/p2p/net/mock"
	"os"
	"path"
	"testing"
)

func TestNewPublicData(t *testing.T) {

	dir := path.Join(os.TempDir(), "openbazaar", "public_test")
	_, err := NewPublicData(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	directories := []string{
		path.Join(dir, "listings"),
		path.Join(dir, "ratings"),
		path.Join(dir, "images"),
		path.Join(dir, "images", "tiny"),
		path.Join(dir, "images", "small"),
		path.Join(dir, "images", "medium"),
		path.Join(dir, "images", "large"),
		path.Join(dir, "images", "original"),
		path.Join(dir, "posts"),
		path.Join(dir, "channel"),
		path.Join(dir, "files"),
	}

	for _, p := range directories {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			t.Errorf("Failed to create directory %s", p)
		}
	}
}

func TestPublicData_Profile(t *testing.T) {
	dir := path.Join(os.TempDir(), "openbazaar", "profile_test")
	pd, err := NewPublicData(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	name := "Ron Swanson"
	err = pd.SetProfile(&models.Profile{
		Name: name,
	})
	if err != nil {
		t.Fatal(err)
	}

	pro, err := pd.GetProfile()
	if err != nil {
		t.Fatal(err)
	}

	if pro.Name != name {
		t.Errorf("Incorrect name returned. Expected %s got %s", name, pro.Name)
	}
}

func TestPublicData_Followers(t *testing.T) {
	dir := path.Join(os.TempDir(), "openbazaar", "followers_test")
	pd, err := NewPublicData(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	l := models.Followers{
		"QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
		"Qmd9hFFuueFrSR7YwUuAfirXXJ7ANZAMc5sx4HFxn7mPkc",
	}
	err = pd.SetFollowers(l)
	if err != nil {
		t.Fatal(err)
	}

	followers, err := pd.GetFollowers()
	if err != nil {
		t.Fatal(err)
	}

	if followers[0] != l[0] {
		t.Errorf("Incorrect peerID returned. Expected %s got %s", l[0], followers[0])
	}
	if followers[1] != l[1] {
		t.Errorf("Incorrect peerID returned. Expected %s got %s", l[1], followers[1])
	}
}

func TestPublicData_Following(t *testing.T) {
	dir := path.Join(os.TempDir(), "openbazaar", "following_test")
	pd, err := NewPublicData(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	l := models.Following{
		"Qmd9hFFuueFrSR7YwUuAfirXXJ7ANZAMc5sx4HFxn7mPkc",
		"QmfQkD8pBSBCBxWEwFSu4XaDVSWK6bjnNuaWZjMyQbyDub",
	}
	err = pd.SetFollowing(l)
	if err != nil {
		t.Fatal(err)
	}

	following, err := pd.GetFollowing()
	if err != nil {
		t.Fatal(err)
	}

	if following[0] != l[0] {
		t.Errorf("Incorrect peerID returned. Expected %s got %s", l[0], following[0])
	}
	if following[1] != l[1] {
		t.Errorf("Incorrect peerID returned. Expected %s got %s", l[1], following[1])
	}
}

func TestPublicData_Publish(t *testing.T) {
	dir := path.Join(os.TempDir(), "openbazaar", "public_test")
	pd, err := NewPublicData(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	name := "Ron Swanson"
	err = pd.SetProfile(&models.Profile{
		Name: name,
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create network
	mn := mocknet.New(ctx)

	var nodes []*core.IpfsNode
	for i := 0; i < 2; i++ {
		nd, err := core.NewNode(ctx, &core.BuildCfg{
			Online: true,
			Host:   coremock.MockHostOption(mn),
		})
		if err != nil {
			t.Fatal(err)
		}

		nd.Namesys = namesys.NewNameSystem(nd.Routing, nd.Repo.Datastore(), 0)

		nodes = append(nodes, nd)
	}

	if err := mn.LinkAll(); err != nil {
		t.Fatal(err)
	}

	bsinf := bootstrap.BootstrapConfigWithPeers(
		[]peerstore.PeerInfo{
			nodes[0].Peerstore.PeerInfo(nodes[0].Identity),
		},
	)

	for _, n := range nodes[1:] {
		if err := n.Bootstrap(bsinf); err != nil {
			t.Fatal(err)
		}
	}

	if err := pd.Publish(context.Background(), nodes[0]); err != nil {
		t.Fatal(err)
	}
}
