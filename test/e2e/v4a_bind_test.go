package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV4aBindBigPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v4a", "bind", 0, t)
}

func TestV4aBindMiddlePicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v4a", "bind", 1, t)
}

func TestV4aBindSmallPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v4a", "bind", 2, t)
}
