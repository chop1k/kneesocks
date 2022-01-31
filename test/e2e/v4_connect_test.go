package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV4ConnectBigPicture(t *testing.T) {
	stand.New().Execute("v4", "connect", 0, t)
}

func TestV4ConnectMiddlePicture(t *testing.T) {
	stand.New().Execute("v4", "connect", 1, t)
}

func TestV4ConnectSmallPicture(t *testing.T) {
	stand.New().Execute("v4", "connect", 2, t)
}
