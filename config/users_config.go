package config

import "socks/config/tree"

type UsersConfig interface {
	GetUsers() []tree.User
}

type BaseUsersConfig struct {
	config tree.Config
}

func NewBaseUsersConfig(config tree.Config) (BaseUsersConfig, error) {
	return BaseUsersConfig{
		config: config,
	}, nil
}

func (b BaseUsersConfig) GetUsers() []tree.User {
	return b.config.Users
}
