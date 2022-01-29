package config

import (
	"github.com/stretchr/testify/require"
	v4 "socks/test/stand/config/v4"
	"socks/test/stand/config/v4a"
	v5 "socks/test/stand/config/v5"
	"testing"
)

type Scope struct {
	t      *testing.T
	config Config
}

func NewScope(t *testing.T, config Config) (Scope, error) {
	return Scope{t: t, config: config}, nil
}

func (s Scope) GetV4Connect(number int) v4.ConnectScope {
	for i, scope := range s.config.V4.Connect {
		if i == number {
			return scope
		}
	}

	require.Fail(s.t, "Cannot find scope for v4 connect. ", number)

	return v4.ConnectScope{}
}

func (s Scope) GetV4Bind(number int) v4.BindScope {
	for i, scope := range s.config.V4.Bind {
		if i == number {
			return scope
		}
	}

	require.Fail(s.t, "Cannot find scope for v4 bind. ", number)

	return v4.BindScope{}
}

func (s Scope) GetV4aBind(number int) v4a.BindScope {
	for i, scope := range s.config.V4a.Bind {
		if i == number {
			return scope
		}
	}

	require.Fail(s.t, "Cannot find scope for v4a bind. ", number)

	return v4a.BindScope{}
}

func (s Scope) GetV4aConnect(number int) v4a.ConnectScope {
	for i, scope := range s.config.V4a.Connect {
		if i == number {
			return scope
		}
	}

	require.Fail(s.t, "Cannot find scope for v4a connect. ", number)

	return v4a.ConnectScope{}
}

func (s Scope) GetV5Connect(number int) v5.ConnectScope {
	for i, scope := range s.config.V5.Connect {
		if i == number {
			return scope
		}
	}

	require.Fail(s.t, "Cannot find scope for v5 connect. ", number)

	return v5.ConnectScope{}
}

func (s Scope) GetV5Auth(number int) v5.AuthScope {
	for i, scope := range s.config.V5.Auth {
		if i == number {
			return scope
		}
	}

	require.Fail(s.t, "Cannot find scope for v5 auth. ", number)

	return v5.AuthScope{}
}

func (s Scope) GetV5Bind(number int) v5.BindScope {
	for i, scope := range s.config.V5.Bind {
		if i == number {
			return scope
		}
	}

	require.Fail(s.t, "Cannot find scope for v5 bind. ", number)

	return v5.BindScope{}
}

func (s Scope) GetV5Associate(number int) v5.AssociateScope {
	for i, scope := range s.config.V5.Associate {
		if i == number {
			return scope
		}
	}

	require.Fail(s.t, "Cannot find scope for v5 associate. ", number)

	return v5.AssociateScope{}
}
