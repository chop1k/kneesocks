package v4a

import (
	"github.com/stretchr/testify/require"
	"socks/test/stand/config"
	"testing"
)

type Test struct {
	_case   config.Case
	t       *testing.T
	connect ConnectTester
	bind    BindTester
}

func NewTest(_case config.Case, t *testing.T, connect ConnectTester, bind BindTester) (Test, error) {
	return Test{
		_case:   _case,
		t:       t,
		connect: connect,
		bind:    bind,
	}, nil
}

func (t Test) Start() {
	switch t._case.Command {
	case "connect":
		t.connect.Test(t._case.Number)
	case "bind":
		t.bind.Test(t._case.Number)
	default:
		require.Fail(t.t, "Unsupported command \"%s\".", t._case.Command)
	}
}
