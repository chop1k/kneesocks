package v4

type ConfigReplicator struct {
	config *Config
}

func NewConfigReplicator(config *Config) (ConfigReplicator, error) {
	return ConfigReplicator{config: config}, nil
}

func (b ConfigReplicator) Copy() *Config {
	if b.config == nil {
		return nil
	}

	config := *b.config

	blacklist := make([]string, len(b.config.Restrictions.BlackList))
	whitelist := make([]string, len(b.config.Restrictions.WhiteList))

	copy(blacklist, b.config.Restrictions.BlackList)
	copy(whitelist, b.config.Restrictions.WhiteList)

	config.Restrictions.BlackList = blacklist
	config.Restrictions.WhiteList = whitelist

	return &config
}
