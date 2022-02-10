package v5

import (
	"errors"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/mitchellh/mapstructure"
	"regexp"
	"socks/config/tree"
)

const pattern = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"

var (
	UserNotExistsError   = errors.New("User not exists. ")
	InvalidUsernameError = errors.New("Invalid username. ")
)

type UsersConfig interface {
	GetUser(name string) (tree.User, error)
	GetRestrictions(name string) (tree.Restrictions, error)
	GetRate(name string) (tree.RateRestrictions, error)
}

type BaseUsersConfig struct {
	config gabs.Container
}

func NewBaseUsersConfig(config gabs.Container) (BaseUsersConfig, error) {
	return BaseUsersConfig{config: config}, nil
}

func (b BaseUsersConfig) GetUser(name string) (tree.User, error) {
	matched, err := regexp.MatchString(pattern, name)

	if err != nil {
		return tree.User{}, err
	}

	if !matched {
		return tree.User{}, InvalidUsernameError
	}

	user, ok := b.config.Path(fmt.Sprintf("SocksV5.Users.%s", name)).Data().(map[string]interface{})

	if !ok {
		return tree.User{}, UserNotExistsError
	}

	_user := tree.User{}

	return _user, mapstructure.Decode(user, &_user)
}

func (b BaseUsersConfig) GetRestrictions(name string) (tree.Restrictions, error) {
	matched, err := regexp.MatchString(pattern, name)

	if err != nil {
		return tree.Restrictions{}, err
	}

	if !matched {
		return tree.Restrictions{}, InvalidUsernameError
	}

	restrictions, ok := b.config.Path(fmt.Sprintf("SocksV5.Users.%s.Restrictions", name)).Data().(map[string]interface{})

	if !ok {
		return tree.Restrictions{}, UserNotExistsError
	}

	_restrictions := tree.Restrictions{}

	return _restrictions, mapstructure.Decode(restrictions, &_restrictions)
}

func (b BaseUsersConfig) GetRate(name string) (tree.RateRestrictions, error) {
	matched, err := regexp.MatchString(pattern, name)

	if err != nil {
		return tree.RateRestrictions{}, err
	}

	if !matched {
		return tree.RateRestrictions{}, InvalidUsernameError
	}

	rate, ok := b.config.Path(fmt.Sprintf("SocksV5.Users.%s.Restrictions.Rate", name)).Data().(map[string]interface{})

	if !ok {
		return tree.RateRestrictions{}, UserNotExistsError
	}

	_rate := tree.RateRestrictions{}

	return _rate, mapstructure.Decode(rate, &_rate)
}
