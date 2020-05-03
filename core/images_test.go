package core

import (
	"context"
	"github.com/cpacia/openbazaar3.0/models"
	"github.com/ipfs/go-cid"
	ipath "github.com/ipfs/interface-go-ipfs-core/path"
	"testing"
	"time"
)

var (
	jpgImageB64 = "/9j/4AAQSkZJRgABAQAAAQABAAD//gA7Q1JFQVRPUjogZ2QtanBlZyB2MS4wICh1c2luZyBJSkcgSlBFRyB2NjIpLCBxdWFsaXR5ID0gNjUK/9sAQwALCAgKCAcLCgkKDQwLDREcEhEPDxEiGRoUHCkkKyooJCcnLTJANy0wPTAnJzhMOT1DRUhJSCs2T1VORlRAR0hF/9sAQwEMDQ0RDxEhEhIhRS4nLkVFRUVFRUVFRUVFRUVFRUVFRUVFRUVFRUVFRUVFRUVFRUVFRUVFRUVFRUVFRUVFRUVF/8AAEQgAMgAyAwEiAAIRAQMRAf/EAB8AAAEFAQEBAQEBAAAAAAAAAAABAgMEBQYHCAkKC//EALUQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+v/EAB8BAAMBAQEBAQEBAQEAAAAAAAABAgMEBQYHCAkKC//EALURAAIBAgQEAwQHBQQEAAECdwABAgMRBAUhMQYSQVEHYXETIjKBCBRCkaGxwQkjM1LwFWJy0QoWJDThJfEXGBkaJicoKSo1Njc4OTpDREVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5eoKDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uLj5OXm5+jp6vLz9PX29/j5+v/aAAwDAQACEQMRAD8A840awhv5zFKWDYyMHrVvWtE/szynj3GJ+MnsaoWFw1ndxTr1Rskeor0+70uPXNBYQ4JkQSRH36iiXw3CO9meWxxNJIqICWY4AHeu5g8C232aMztL5pUFtpGM/lUXgPw+13qD3lwhEdscAEdX/wDrVseNddl0l4bSxcLcN8zHAOB6c1UnyJLqxRTlJ9kY83guzQcNN/30P8KwNY0W206AvufceFBPWvRtMtrw6RHLqUm+dxvOVA2j04rzjxJqAv8AUXEZ/cxHavv71M20+UcbNc3Q5/bRUu2igCVRXpfw51MXFtJp0rfPD88ee6nr+R/nXmq13fw40xpL6TUXyEiGxfcnrVwV7kSdrHo7C10ixnn2rFEu6R8cZPU15r4espfFviua/uQTbxvvbPT/AGVrX+IetMyQ6PakmSUhpAv6Cuh0DT4PC/hsGbCsE82ZvfHSs4O160umiLmtFTW73/rzMjx7rC6Zp32WFgJ7gY4/hXua8nbmtTXtWk1rVZruQnDHCL/dXtWW1TBPd7suVl7q6EeKKKKsgkt42nlSNBlmIAr1rTZINA0MAkBYU3MfU15joEsEN5588irs+6GPetfX9cW8iis7eVSjHLsDxRJ+7yx3YRV5XeyNjwlbPrviCbWL0bkjbcoPQt2H4Vf+IOsTyxpplrHIyt80rKpwfQU3R9W0vS9PitkvIBtHzHeOT3q+3ibTiP8Aj9g/77FE+R2itkEXK7m92eXm2uB1gk/74NRPDKoOY3H1U16RceI7FgcXkJ/4GKwdT1q2lgkVJ0YlSOGoco20CzONzRUe6igBq1KtFFAD6KKKAGmo26UUUAR0UUUAf//Z"

	gifImageB64 = "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"

	pngImageB64 = "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"
)

func TestOpenBazaarNode_SetImages(t *testing.T) {
	m, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})

	err = m.SetProfile(&models.Profile{
		Name: "Ron Paul",
	}, done)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("timed out waiting on channel")
	}

	done2 := make(chan struct{})
	hashes, err := m.SetAvatarImage(jpgImageB64, done2)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-done2:
	case <-time.After(time.Second * 10):
		t.Fatal("timed out waiting on channel")
	}

	var (
		expectedTiny     = `QmNyXxH9KLa71QT8mgSyoq9kK8M8ns765nyVndzJD4gxeu`
		expectedSmall    = `QmWSUorwepjpELNvJfq8abQHdXUC78eiYmCeh5uMbb6Xuq`
		expectedMedium   = `QmdYmbKG5e7Y9f8FGRWo3CjuGJgptsSRKFBjNr8qzupqyN`
		expectedLarge    = `QmQE8S5VbUL9bhKkTSPiTcsCBZYA4kYpboq6cyJXCzPLrf`
		expectedOriginal = `QmZq8EGL18wA9tAwAa2EvqMW5GSEJyzkLDw21XmnkFkEVA`
	)

	if hashes.Tiny != expectedTiny {
		t.Errorf("Expected %s, got %s", expectedTiny, hashes.Tiny)
	}
	if hashes.Small != expectedSmall {
		t.Errorf("Expected %s, got %s", expectedSmall, hashes.Small)
	}
	if hashes.Medium != expectedMedium {
		t.Errorf("Expected %s, got %s", expectedMedium, hashes.Medium)
	}
	if hashes.Large != expectedLarge {
		t.Errorf("Expected %s, got %s", expectedLarge, hashes.Large)
	}
	if hashes.Original != expectedOriginal {
		t.Errorf("Expected %s, got %s", expectedOriginal, hashes.Original)
	}

	id, err := cid.Decode(hashes.Tiny)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Small)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Medium)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Large)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Original)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenBazaarNode_SetHeader(t *testing.T) {
	m, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})

	err = m.SetProfile(&models.Profile{
		Name: "Ron Paul",
	}, done)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("timed out waiting on channel")
	}

	done2 := make(chan struct{})
	hashes, err := m.SetHeaderImage(jpgImageB64, done2)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-done2:
	case <-time.After(time.Second * 10):
		t.Fatal("timed out waiting on channel")
	}

	var (
		expectedTiny     = `QmP6Sr6aiJhEHseLVi5tBSgekqTgySGp3cLdHFtEHE8oxS`
		expectedSmall    = `QmdDjvzzJ7ifvkBPbhgxyqYTvWmKaWjgctFcHV94dDW4qG`
		expectedMedium   = `QmTcj5UKy4EzMpFJR7GmLHPCgF8kfgabS78AfAA41xkGQ5`
		expectedLarge    = `Qmd8smgsc57L8SXHX2Dkgq53HVrJ2XQAVqEm91nPHaDdn7`
		expectedOriginal = `QmZq8EGL18wA9tAwAa2EvqMW5GSEJyzkLDw21XmnkFkEVA`
	)

	if hashes.Tiny != expectedTiny {
		t.Errorf("Expected %s, got %s", expectedTiny, hashes.Tiny)
	}
	if hashes.Small != expectedSmall {
		t.Errorf("Expected %s, got %s", expectedSmall, hashes.Small)
	}
	if hashes.Medium != expectedMedium {
		t.Errorf("Expected %s, got %s", expectedMedium, hashes.Medium)
	}
	if hashes.Large != expectedLarge {
		t.Errorf("Expected %s, got %s", expectedLarge, hashes.Large)
	}
	if hashes.Original != expectedOriginal {
		t.Errorf("Expected %s, got %s", expectedOriginal, hashes.Original)
	}

	id, err := cid.Decode(hashes.Tiny)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Small)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Medium)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Large)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Original)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenBazaarNode_SetProductImage(t *testing.T) {
	m, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})

	err = m.SetProfile(&models.Profile{
		Name: "Ron Paul",
	}, done)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("timed out waiting on channel")
	}

	hashes, err := m.SetProductImage(jpgImageB64, "testfile")
	if err != nil {
		t.Fatal(err)
	}

	done2 := make(chan struct{})
	m.Publish(done2)

	select {
	case <-done2:
	case <-time.After(time.Second * 10):
		t.Fatal("timed out waiting on channel")
	}

	var (
		expectedTiny     = `QmWSUorwepjpELNvJfq8abQHdXUC78eiYmCeh5uMbb6Xuq`
		expectedSmall    = `QmdYmbKG5e7Y9f8FGRWo3CjuGJgptsSRKFBjNr8qzupqyN`
		expectedMedium   = `QmQE8S5VbUL9bhKkTSPiTcsCBZYA4kYpboq6cyJXCzPLrf`
		expectedLarge    = `QmdFm3TnRGYNxKECQBFy6vLCPDaNT26ZUfReD64v7RVEVe`
		expectedOriginal = `QmZq8EGL18wA9tAwAa2EvqMW5GSEJyzkLDw21XmnkFkEVA`
	)

	if hashes.Tiny != expectedTiny {
		t.Errorf("Expected %s, got %s", expectedTiny, hashes.Tiny)
	}
	if hashes.Small != expectedSmall {
		t.Errorf("Expected %s, got %s", expectedSmall, hashes.Small)
	}
	if hashes.Medium != expectedMedium {
		t.Errorf("Expected %s, got %s", expectedMedium, hashes.Medium)
	}
	if hashes.Large != expectedLarge {
		t.Errorf("Expected %s, got %s", expectedLarge, hashes.Large)
	}
	if hashes.Original != expectedOriginal {
		t.Errorf("Expected %s, got %s", expectedOriginal, hashes.Original)
	}

	id, err := cid.Decode(hashes.Tiny)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Small)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Medium)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Large)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}

	id, err = cid.Decode(hashes.Original)
	if err != nil {
		t.Fatal(err)
	}
	_, err = m.cat(context.Background(), ipath.IpfsPath(id))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGIFAndPNG(t *testing.T) {
	m, err := MockNode()
	if err != nil {
		t.Fatal(err)
	}

	done := make(chan struct{})

	err = m.SetProfile(&models.Profile{
		Name: "Ron Paul",
	}, done)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case <-done:
	case <-time.After(time.Second * 10):
		t.Fatal("timed out waiting on channel")
	}

	_, err = m.SetAvatarImage(gifImageB64, nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = m.SetAvatarImage(pngImageB64, nil)
	if err != nil {
		t.Fatal(err)
	}
}
