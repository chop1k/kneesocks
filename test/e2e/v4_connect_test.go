package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV4ConnectWithBigPicture(t *testing.T) {
	stand.New().Execute("v4", "connect", 1, t)
}

func TestV4ConnectWithMiddlePicture(t *testing.T) {
	stand.New().Execute("v4", "connect", 2, t)
}

func TestV4ConnectWithSmallPicture(t *testing.T) {
	stand.New().Execute("v4", "connect", 3, t)
}
