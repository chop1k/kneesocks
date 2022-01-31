package managers

import (
	"regexp"
	"socks/logger"
)

type BlacklistManager interface {
	IsBlacklisted(list []string, address string) bool
}

type BaseBlacklistManager struct {
	logger logger.ServerLogger
}

func NewBaseBlacklistManager(
	logger logger.ServerLogger,
) (BaseBlacklistManager, error) {
	return BaseBlacklistManager{
		logger: logger,
	}, nil
}

func (b BaseBlacklistManager) IsBlacklisted(list []string, address string) bool {
	for _, pattern := range list {
		matched, err := regexp.MatchString(pattern, address)

		if err != nil {
			b.logger.BlacklistMatchError(address, pattern, err)

			continue
		}

		if matched {
			return true
		}
	}

	return false
}
