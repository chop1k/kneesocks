package server

import (
	"regexp"
	"socks/config"
)

type BlacklistManager interface {
	IsBlacklisted(address string) bool
}

type BaseBlacklistManager struct {
	config config.BlacklistConfig
}

func NewBaseBlacklistManager(config config.BlacklistConfig) (BaseBlacklistManager, error) {
	return BaseBlacklistManager{
		config: config,
	}, nil
}

func (b BaseBlacklistManager) IsBlacklisted(address string) bool {
	list := b.config.GetBlacklist()

	for _, pattern := range list {
		matched, err := regexp.MatchString(pattern, address)

		if err != nil {
			// b.errors.Unexpected(pattern, address, err)

			continue
		}

		if matched {
			return true
		}
	}

	return false
}
