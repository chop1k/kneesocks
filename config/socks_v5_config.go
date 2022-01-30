package config

import "socks/config/tree"

type SocksV5Config interface {
	IsConnectAllowed() bool
	IsBindAllowed() bool
	GetConnectDeadline() uint
	GetBindDeadline() uint
	IsUdpAssociationAllowed() bool
	IsIPv4Allowed() bool
	IsIPv6Allowed() bool
	IsDomainAllowed() bool
	GetAuthenticationMethods() []string
	GetUsers() map[string]tree.User
}

type BaseSocksV5Config struct {
	config tree.Config
}

func NewBaseSocksV5Config(config tree.Config) BaseSocksV5Config {
	return BaseSocksV5Config{config: config}
}

func (b BaseSocksV5Config) IsConnectAllowed() bool {
	return b.config.SocksV5.AllowConnect
}

func (b BaseSocksV5Config) IsBindAllowed() bool {
	return b.config.SocksV5.AllowBind
}

func (b BaseSocksV5Config) GetConnectDeadline() uint {
	return b.config.SocksV5.ConnectDeadline
}

func (b BaseSocksV5Config) GetBindDeadline() uint {
	return b.config.SocksV5.BindDeadline
}

func (b BaseSocksV5Config) IsUdpAssociationAllowed() bool {
	return b.config.SocksV5.AllowUdpAssociation
}

func (b BaseSocksV5Config) IsIPv4Allowed() bool {
	return b.config.SocksV5.AllowIPv4
}

func (b BaseSocksV5Config) IsIPv6Allowed() bool {
	return b.config.SocksV5.AllowIPv6
}

func (b BaseSocksV5Config) IsDomainAllowed() bool {
	return b.config.SocksV5.AllowDomain
}

func (b BaseSocksV5Config) GetAuthenticationMethods() []string {
	return b.config.SocksV5.AuthenticationMethodsAllowed
}

func (b BaseSocksV5Config) GetUsers() map[string]tree.User {
	return b.config.SocksV5.Users
}
