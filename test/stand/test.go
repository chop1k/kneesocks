package stand

import (
	"github.com/stretchr/testify/require"
	"socks/test/stand/config"
	v4 "socks/test/stand/v4"
	"socks/test/stand/v4a"
	v5 "socks/test/stand/v5"
	"testing"
)

type Test struct {
	v4    v4.Test
	v4a   v4a.Test
	v5    v5.Test
	t     *testing.T
	_case config.Case
}

func NewTest(
	t *testing.T,
	_case config.Case,
	v4 v4.Test,
	v4a v4a.Test,
	v5 v5.Test,
) (Test, error) {
	return Test{
		t:     t,
		_case: _case,
		v4:    v4,
		v4a:   v4a,
		v5:    v5,
	}, nil
}

func (t Test) Start() {
	if t._case.Protocol == "v4" {
		t.v4.Start()
	} else if t._case.Protocol == "v4a" {
		t.v4a.Start()
	} else if t._case.Protocol == "v4a" {
		t.v5.Start()
	} else {
		require.Fail(t.t, "Unsupported protocol \"%s\".", t._case.Protocol)
	}
}
