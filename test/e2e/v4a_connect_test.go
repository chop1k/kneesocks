package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV4aConnectWithBigPicture(t *testing.T) {
	stand.New().Execute("v4a", "connect", 0, t)
}

func TestV4aConnectWithMiddlePicture(t *testing.T) {
	stand.New().Execute("v4a", "connect", 1, t)
}

func TestV4aConnectWithSmallPicture(t *testing.T) {
	stand.New().Execute("v4a", "connect", 2, t)
}
