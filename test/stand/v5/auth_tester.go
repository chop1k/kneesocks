package v5

import (
	v5 "socks/protocol/v5"
	"socks/test/stand/config"
	"testing"
)

type AuthTester struct {
	t       *testing.T
	config  config.Config
	builder v5.Builder
}
