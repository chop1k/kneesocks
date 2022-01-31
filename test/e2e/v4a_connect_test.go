package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV4aConnectBigPicture(t *testing.T) {
	stand.New().Execute("v4a", "connect", 0, t)
}

func TestV4aConnectMiddlePicture(t *testing.T) {
	stand.New().Execute("v4a", "connect", 1, t)
}

func TestV4aConnectSmallPicture(t *testing.T) {
	stand.New().Execute("v4a", "connect", 2, t)
}
