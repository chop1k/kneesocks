package stand

import (
	"github.com/stretchr/testify/require"
	"socks/test/stand/config"
	"socks/test/stand/v4"
	"socks/test/stand/v4a"
	"socks/test/stand/v5"
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
	switch t._case.Protocol {
	case "v4":
		t.v4.Start()
	case "v4a":
		t.v4a.Start()
	case "v5":
		t.v5.Start()
	default:
		require.Fail(t.t, "Unsupported protocol \"%s\".", t._case.Protocol)
	}
}
