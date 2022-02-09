package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV5UdpAssociationByDomainWithSmallPicture(t *testing.T) {
	stand.New().Execute("v5", "associate", 0, t)
}

func TestV5UdpAssociationByIPv4WithSmallPicture(t *testing.T) {
	stand.New().Execute("v5", "associate", 1, t)
}

func TestV5UdpAssociationByIPv6WithSmallPicture(t *testing.T) {
	stand.New().Execute("v5", "associate", 2, t)
}
