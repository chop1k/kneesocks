package e2e

import (
	"socks/test/stand"
	"testing"
)

func TestV4aBindWithBigPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v4a", "bind", 0, t)
}

func TestV4aBindWithMiddlePicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v4a", "bind", 1, t)
}

func TestV4aBindWithSmallPicture(t *testing.T) {
	t.Skip("Random behavior.¯\\_( ͡° ͜ʖ ͡°)_/¯")

	stand.New().Execute("v4a", "bind", 2, t)
}
