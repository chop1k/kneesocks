package v5

import "socks/config/tree"

type Config interface {
	IsConnectAllowed() bool
	IsBindAllowed() bool
	IsUdpAssociationAllowed() bool
	IsIPv4Allowed() bool
	IsIPv6Allowed() bool
	IsDomainAllowed() bool
	GetAuthenticationMethods() []string
	GetUsers() map[string]tree.User
}

type BaseConfig struct {
	config tree.Config
}

func NewBaseConfig(config tree.Config) (BaseConfig, error) {
	return BaseConfig{config: config}, nil
}

func (b BaseConfig) IsConnectAllowed() bool {
	return b.config.SocksV5.AllowConnect
}

func (b BaseConfig) IsBindAllowed() bool {
	return b.config.SocksV5.AllowBind
}

func (b BaseConfig) IsUdpAssociationAllowed() bool {
	return b.config.SocksV5.AllowUdpAssociation
}

func (b BaseConfig) IsIPv4Allowed() bool {
	return b.config.SocksV5.AllowIPv4
}

func (b BaseConfig) IsIPv6Allowed() bool {
	return b.config.SocksV5.AllowIPv6
}

func (b BaseConfig) IsDomainAllowed() bool {
	return b.config.SocksV5.AllowDomain
}

func (b BaseConfig) GetAuthenticationMethods() []string {
	return b.config.SocksV5.AuthenticationMethodsAllowed
}

func (b BaseConfig) GetUsers() map[string]tree.User {
	return b.config.SocksV5.Users
}
