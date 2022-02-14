package managers

import (
	"errors"
	"socks/internal/kneesocks/config/tree"
)

var (
	ClientAlreadyExistsError = errors.New("Client already exists. ")
)

type BindRateManager struct {
	rate map[string]tree.RateRestrictions
}

func NewBindRateManager() (BindRateManager, error) {
	return BindRateManager{
		rate: make(map[string]tree.RateRestrictions),
	}, nil
}

func (b BindRateManager) Add(client string, rate tree.RateRestrictions) error {
	_, ok := b.rate[client]

	if ok {
		return ClientAlreadyExistsError
	} else {
		b.rate[client] = rate
	}

	return nil
}

func (b BindRateManager) Get(client string) (tree.RateRestrictions, error) {
	restrictions, ok := b.rate[client]

	if !ok {
		return tree.RateRestrictions{}, ClientNotExistsError
	}

	return restrictions, nil
}

func (b BindRateManager) Remove(client string) {
	delete(b.rate, client)
}
