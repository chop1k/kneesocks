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
	if t._case.Command == "connect" {
		t.connect.Test(t._case.Number)
	} else if t._case.Command == "bind" {
		t.bind.Test(t._case.Number)
	} else {
		require.Fail(t.t, "Unsupported command \"%s\".", t._case.Command)
	}
}
