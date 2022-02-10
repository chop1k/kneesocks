package v4

import (
	"errors"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/mitchellh/mapstructure"
	"socks/config/tree"
)

type RestrictionsConfig interface {
	GetWhitelist() ([]string, error)
	GetBlacklist() ([]string, error)
	GetRate() (tree.RateRestrictions, error)
}

type BaseRestrictionsConfig struct {
	config gabs.Container
}

func NewBaseRestrictionsConfig(config gabs.Container) (BaseRestrictionsConfig, error) {
	return BaseRestrictionsConfig{config: config}, nil
}

func (b BaseRestrictionsConfig) GetWhitelist() ([]string, error) {
	whitelist, ok := b.config.Path("SocksV4.Restrictions.Whitelist").Data().([]interface{})

	if !ok {
		return nil, errors.New("SocksV4.Restrictions.Whitelist: Not specified or have invalid type. ")
	}

	var _whitelist []string

	for i, v := range whitelist {
		_v, ok := v.(string)

		if !ok {
			return nil, errors.New(fmt.Sprintf("SocksV4.Restrictions.Whitelist.%d: Invalid type, must be string. ", i))
		}

		_whitelist = append(_whitelist, _v)
	}

	return _whitelist, nil
}

func (b BaseRestrictionsConfig) GetBlacklist() ([]string, error) {
	blacklist, ok := b.config.Path("SocksV4.Restrictions.Blacklist").Data().([]interface{})

	if !ok {
		return nil, errors.New("SocksV4.Restrictions.Blacklist: Not specified or have invalid type. ")
	}

	var _blacklist []string

	for i, v := range blacklist {
		_v, ok := v.(string)

		if !ok {
			return nil, errors.New(fmt.Sprintf("SocksV4.Restrictions.Blacklist.%d: Invalid type, must be string. ", i))
		}

		_blacklist = append(_blacklist, _v)
	}

	return _blacklist, nil
}

func (b BaseRestrictionsConfig) GetRate() (tree.RateRestrictions, error) {
	rate, ok := b.config.Path("SocksV4.Restrictions.Rate").Data().(map[string]interface{})

	if !ok {
		return tree.RateRestrictions{}, errors.New("SocksV4.Restrictions.Rate: Not specified or have invalid type. ")
	}

	_rate := tree.RateRestrictions{}

	return _rate, mapstructure.Decode(rate, &_rate)
}
