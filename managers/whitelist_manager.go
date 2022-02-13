package managers

import (
	"regexp"
)

type WhitelistManager interface {
	IsWhitelisted(list []string, address string) bool
}

type BaseWhitelistManager struct {
}

func NewBaseWhitelistManager() (BaseWhitelistManager, error) {
	return BaseWhitelistManager{}, nil
}

func (b BaseWhitelistManager) IsWhitelisted(list []string, address string) bool {
	if len(list) <= 0 {
		return false
	}

	for _, pattern := range list {
		matched, err := regexp.MatchString(pattern, address)

		if err != nil {
			continue
		}

		if matched {
			return false
		}
	}

	return true
}
