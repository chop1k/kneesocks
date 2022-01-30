package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV5BindByDomainWithBigPicture(t *testing.T) {
	//t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 0, t)
}

func TestV5BindByDomainWithMiddlePicture(t *testing.T) {
	//t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 1, t)
}

func TestV5BindByDomainWithSmallPicture(t *testing.T) {
	//t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 2, t)
}

func TestV5BindByIPv4WithBigPicture(t *testing.T) {
	//t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 3, t)
}

func TestV5BindByIPv4WithMiddlePicture(t *testing.T) {
	//t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 4, t)
}

func TestV5BindByIPv4WithSmallPicture(t *testing.T) {
	//t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 5, t)
}

func TestV5BindByIPv6WithBigPicture(t *testing.T) {
	//t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 6, t)
}

func TestV5BindByIPv6WithMiddlePicture(t *testing.T) {
	//t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 7, t)
}

func TestV5BindByIPv6WithSmallPicture(t *testing.T) {
	//t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 8, t)
}
