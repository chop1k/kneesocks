package v4

import (
	v4 "socks/protocol/v4"
	"socks/test/stand/config"
	"socks/test/stand/picture"
	"testing"
)

type BindTester struct {
	config  config.Config
	t       *testing.T
	build   v4.Builder
	picture picture.Picture
}

func NewBindTester(
	config config.Config,
	t *testing.T,
	build v4.Builder,
	picture picture.Picture,
) (BindTester, error) {
	return BindTester{
		config:  config,
		t:       t,
		build:   build,
		picture: picture,
	}, nil
}

func (t BindTester) Test(number int) {

}
