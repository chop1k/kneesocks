package v5

import (
	"github.com/stretchr/testify/require"
	"socks/test/stand/config"
	"testing"
)

type Test struct {
	_case config.Case
	t     *testing.T
}

func NewTest(_case config.Case, t *testing.T) (Test, error) {
	return Test{
		_case: _case,
		t:     t,
	}, nil
}

func (t Test) Start() {
	if t._case.Command == "connect" {

	} else if t._case.Command == "bind" {

	} else if t._case.Command == "auth" {

	} else if t._case.Command == "associate" {

	} else {
		require.Fail(t.t, "Unsupported command \"%s\".", t._case.Command)
	}
}
