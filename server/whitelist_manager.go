package server

import (
	"regexp"
	"socks/config"
	"socks/logger"
)

type WhitelistManager interface {
	IsWhitelisted(address string) bool
}

type BaseWhitelistManager struct {
	config config.WhitelistConfig
	logger logger.ServerLogger
}

func NewBaseWhitelistManager(
	config config.WhitelistConfig,
	logger logger.ServerLogger,
) (BaseWhitelistManager, error) {
	return BaseWhitelistManager{
		config: config,
		logger: logger,
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
			b.logger.WhitelistMatchError(address, pattern, err)

			continue
		}

		if matched {
			return false
		}
	}

	return true
}
