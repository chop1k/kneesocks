package v5

import (
	"github.com/stretchr/testify/require"
	"socks/test/stand/config"
	"testing"
)

type Test struct {
	_case   config.Case
	t       *testing.T
	auth    AuthTester
	connect ConnectTester
}

func NewTest(
	_case config.Case,
	t *testing.T,
	auth AuthTester,
	connect ConnectTester,
) (Test, error) {
	return Test{
		_case:   _case,
		t:       t,
		auth:    auth,
		connect: connect,
	}, nil
}

func (t Test) Start() {
	if t._case.Command == "connect" {
		t.connect.Test(t._case.Number)
	} else if t._case.Command == "bind" {

	} else if t._case.Command == "auth" {
		t.auth.Test(t._case.Number)
	} else if t._case.Command == "associate" {

	} else {
		require.Fail(t.t, "Unsupported command \"%s\".", t._case.Command)
	}
}
