package v4a

import "socks/config/tree"

type RestrictionsConfig interface {
	GetWhitelist() []string
	GetBlacklist() []string
	GetRate() tree.RateRestrictions
}

type BaseRestrictionsConfig struct {
	tree tree.Config
}

func NewBaseRestrictionsConfig(tree tree.Config) (BaseRestrictionsConfig, error) {
	return BaseRestrictionsConfig{tree: tree}, nil
}

func (b BaseRestrictionsConfig) GetWhitelist() []string {
	return b.tree.SocksV4a.Restrictions.WhiteList
}

func (b BaseRestrictionsConfig) GetBlacklist() []string {
	return b.tree.SocksV4a.Restrictions.BlackList
}

func (b BaseRestrictionsConfig) GetRate() tree.RateRestrictions {
	return b.tree.SocksV4a.Restrictions.Rate
}
