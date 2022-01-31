package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV5ConnectBigPictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "connect", 0, t)
}

func TestV5ConnectBigPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "connect", 1, t)
}

func TestV5ConnectBigPictureWithIPv6(t *testing.T) {
	stand.New().Execute("v5", "connect", 2, t)
}

func TestV5ConnectMiddlePictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "connect", 3, t)
}

func TestV5ConnectMiddlePictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "connect", 4, t)
}

func TestV5ConnectMiddlePictureWithIPv6(t *testing.T) {

	stand.New().Execute("v5", "connect", 5, t)
}

func TestV5ConnectSmallPictureWithIPv4(t *testing.T) {
	stand.New().Execute("v5", "connect", 6, t)
}

func TestV5ConnectSmallPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "connect", 7, t)
}

func TestV5ConnectSmallPictureWithIPv6(t *testing.T) {
	stand.New().Execute("v5", "connect", 8, t)
}
