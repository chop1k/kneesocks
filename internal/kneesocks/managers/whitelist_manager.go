package managers

import (
	"regexp"
)

type WhitelistManager struct {
}

func NewWhitelistManager() (WhitelistManager, error) {
	return WhitelistManager{}, nil
}

func (b WhitelistManager) IsWhitelisted(list []string, address string) bool {
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
