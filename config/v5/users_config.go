package v5

import (
	"errors"
	"socks/config/tree"
)

var (
	UserNotExistsError = errors.New("User not exists. ")
)

type UsersConfig interface {
	GetUser(name string) (tree.User, error)
	GetRestrictions(name string) (tree.Restrictions, error)
	GetRate(name string) (tree.RateRestrictions, error)
}

type BaseUsersConfig struct {
	tree tree.Config
}

func NewBaseUsersConfig(tree tree.Config) (BaseUsersConfig, error) {
	return BaseUsersConfig{tree: tree}, nil
}

func (b BaseUsersConfig) GetUser(name string) (tree.User, error) {
	for username, user := range b.tree.SocksV5.Users {
		if username == name {
			return user, nil
		}
	}

	return tree.User{}, UserNotExistsError
}

func (b BaseUsersConfig) GetRestrictions(name string) (tree.Restrictions, error) {
	for username, user := range b.tree.SocksV5.Users {
		if username == name {
			return user.Restrictions, nil
		}
	}

	return tree.Restrictions{}, UserNotExistsError
}

func (b BaseUsersConfig) GetRate(name string) (tree.RateRestrictions, error) {
	for username, user := range b.tree.SocksV5.Users {
		if username == name {
			return user.Restrictions.Rate, nil
		}
	}

	return tree.RateRestrictions{}, UserNotExistsError
}
