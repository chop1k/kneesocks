package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestBigPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 10, t)
}

func TestMiddlePictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 13, t)
}

func TestSmallPictureWithDomain(t *testing.T) {
	stand.New().Execute("v5", "auth", 16, t)
}
