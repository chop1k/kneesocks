package managers

import (
	"regexp"
)

type BlacklistManager interface {
	IsBlacklisted(list []string, address string) bool
}

type BaseBlacklistManager struct {
}

func NewBaseBlacklistManager() (BaseBlacklistManager, error) {
	return BaseBlacklistManager{}, nil
}

func (b BaseBlacklistManager) IsBlacklisted(list []string, address string) bool {
	for _, pattern := range list {
		matched, err := regexp.MatchString(pattern, address)

		if err != nil {
			continue
		}

		if matched {
			return true
		}
	}

	return false
}
