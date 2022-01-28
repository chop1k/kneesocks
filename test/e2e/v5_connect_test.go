package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV5ConnectByDomainWithBigPicture(t *testing.T) {
	stand.New().Execute("v5", "connect", 19, t)
}

func TestV5ConnectByDomainWithMiddlePicture(t *testing.T) {
	stand.New().Execute("v5", "connect", 22, t)
}

func TestV5ConnectByDomainWithSmallPicture(t *testing.T) {
	stand.New().Execute("v5", "connect", 25, t)
}

func TestV5ConnectByIPv4WithBigPicture(t *testing.T) {
	stand.New().Execute("v5", "connect", 18, t)
}

func TestV5ConnectByIPv4WithMiddlePicture(t *testing.T) {
	stand.New().Execute("v5", "connect", 21, t)
}

func TestV5ConnectByIPv4WithSmallPicture(t *testing.T) {
	stand.New().Execute("v5", "connect", 24, t)
}

func TestV5ConnectByIPv6WithBigPicture(t *testing.T) {
	stand.New().Execute("v5", "connect", 20, t)
}

func TestV5ConnectByIPv6WithMiddlePicture(t *testing.T) {
	stand.New().Execute("v5", "connect", 23, t)
}

func TestV5ConnectByIPv6WithSmallPicture(t *testing.T) {
	stand.New().Execute("v5", "connect", 26, t)
}
