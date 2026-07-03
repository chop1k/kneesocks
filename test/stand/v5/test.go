package v5

import (
	"github.com/stretchr/testify/require"
	"socks/test/stand/config"
	"testing"
)

type Test struct {
	_case     config.Case
	t         *testing.T
	auth      AuthTester
	connect   ConnectTester
	bind      BindTester
	associate AssociationTester
}

func NewTest(
	_case config.Case,
	t *testing.T,
	auth AuthTester,
	connect ConnectTester,
	bind BindTester,
	associate AssociationTester,
) (Test, error) {
	return Test{
		_case:     _case,
		t:         t,
		auth:      auth,
		connect:   connect,
		bind:      bind,
		associate: associate,
	}, nil
}

func (t Test) Start() {
	switch t._case.Command {
	case "connect":
		t.connect.Test(t._case.Number)
	case "bind":
		t.bind.Test(t._case.Number)
	case "auth":
		t.auth.Test(t._case.Number)
	case "associate":
		t.associate.Test(t._case.Number)
	default:
		require.Fail(t.t, "Unsupported command \"%s\".", t._case.Command)
	}
}
