package server

import (
	"regexp"
	"socks/config"
)

type WhitelistManager interface {
	IsWhitelisted(address string) bool
}

type BaseWhitelistManager struct {
	config config.WhitelistConfig
}

func NewBaseWhitelistManager(config config.WhitelistConfig) (BaseWhitelistManager, error) {
	return BaseWhitelistManager{
		config: config,
	}, nil
}

func (b BaseWhitelistManager) IsWhitelisted(address string) bool {
	list := b.config.GetWhitelist()

	if len(list) <= 0 {
		return false
	}

	for _, pattern := range list {
		matched, err := regexp.MatchString(pattern, address)

		if err != nil {
			// b.errors.Unexpected(pattern, address, err)

			continue
		}

		if matched {
			return false
		}
	}

	return true
}
