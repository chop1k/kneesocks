package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV4BindBigPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")
	
	stand.New().Execute("v4", "bind", 0, t)
}

func TestV4BindMiddlePicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v4", "bind", 1, t)
}

func TestV4BindSmallPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v4", "bind", 2, t)
}
