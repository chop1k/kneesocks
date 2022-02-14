package managers

import (
	"regexp"
)

type BlacklistManager struct {
}

func NewBlacklistManager() (BlacklistManager, error) {
	return BlacklistManager{}, nil
}

func (b BlacklistManager) IsBlacklisted(list []string, address string) bool {
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
