package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV5BindBigPictureWithIPv4(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 0, t)
}

func TestV5BindBigPictureWithDomain(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 1, t)
}

func TestV5BindBigPictureWithIPv6(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 2, t)
}

func TestV5BindMiddlePictureWithIPv4(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 3, t)
}

func TestV5BindMiddlePictureWithDomain(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 4, t)
}

func TestV5BindMiddlePictureWithIPv6(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 5, t)
}

func TestV5BindSmallPictureWithIPv4(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 6, t)
}

func TestV5BindSmallPictureWithDomain(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 7, t)
}

func TestV5BindSmallPictureWithIPv6(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v5", "bind", 8, t)
}
