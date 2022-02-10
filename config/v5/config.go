package v5

import (
	"errors"
	"fmt"
	"github.com/Jeffail/gabs"
)

type Config interface {
	IsConnectAllowed() (bool, error)
	IsBindAllowed() (bool, error)
	IsUdpAssociationAllowed() (bool, error)
	IsIPv4Allowed() (bool, error)
	IsIPv6Allowed() (bool, error)
	IsDomainAllowed() (bool, error)
	GetAuthenticationMethods() ([]string, error)
}

type BaseConfig struct {
	config gabs.Container
}

func NewBaseConfig(config gabs.Container) (BaseConfig, error) {
	return BaseConfig{config: config}, nil
}

func (b BaseConfig) IsConnectAllowed() (bool, error) {
	allowed, ok := b.config.Path("SocksV5.AllowConnect").Data().(bool)

	if !ok {
		return false, errors.New("SocksV5.AllowConnect: Not specified or have invalid type. ")
	}

	return allowed, nil
}

func (b BaseConfig) IsBindAllowed() (bool, error) {
	allowed, ok := b.config.Path("SocksV5.AllowBind").Data().(bool)

	if !ok {
		return false, errors.New("SocksV5.AllowBind: Not specified or have invalid type. ")
	}

	return allowed, nil
}

func (b BaseConfig) IsUdpAssociationAllowed() (bool, error) {
	allowed, ok := b.config.Path("SocksV5.AllowUdpAssociation").Data().(bool)

	if !ok {
		return false, errors.New("SocksV5.AllowUdpAssociation: Not specified or have invalid type. ")
	}

	return allowed, nil
}

func (b BaseConfig) IsIPv4Allowed() (bool, error) {
	allowed, ok := b.config.Path("SocksV5.AllowIPv4").Data().(bool)

	if !ok {
		return false, errors.New("SocksV5.AllowIPv4: Not specified or have invalid type. ")
	}

	return allowed, nil
}

func (b BaseConfig) IsIPv6Allowed() (bool, error) {
	allowed, ok := b.config.Path("SocksV5.AllowIPv6").Data().(bool)

	if !ok {
		return false, errors.New("SocksV5.AllowIPv6: Not specified or have invalid type. ")
	}

	return allowed, nil
}

func (b BaseConfig) IsDomainAllowed() (bool, error) {
	allowed, ok := b.config.Path("SocksV5.AllowDomain").Data().(bool)

	if !ok {
		return false, errors.New("SocksV5.AllowDomain: Not specified or have invalid type. ")
	}

	return allowed, nil
}

func (b BaseConfig) GetAuthenticationMethods() ([]string, error) {
	methods, ok := b.config.Path("SocksV5.AuthenticationMethodsAllowed").Data().([]interface{})

	if !ok {
		return nil, errors.New("SocksV5.AuthenticationMethodsAllowed: Not specified or have invalid type. ")
	}

	var _methods []string

	for i, m := range methods {
		_v, ok := m.(string)

		if !ok {
			return nil, errors.New(fmt.Sprintf("SocksV5.AuthenticationMethodsAllowed.%d: Invalid type, must be string. ", i))
		}

		_methods = append(_methods, _v)
	}

	return _methods, nil
}
