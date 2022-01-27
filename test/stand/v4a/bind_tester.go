package v4a

import (
	"socks/protocol/v4a"
	"socks/test/stand/config"
	"socks/test/stand/picture"
	"testing"
)

type BindTester struct {
	config  config.Config
	t       *testing.T
	build   v4a.Builder
	picture picture.Picture
}

func NewBindTester(
	config config.Config,
	t *testing.T,
	build v4a.Builder,
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
