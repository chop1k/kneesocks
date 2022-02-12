package v4

type ConfigBuilder struct {
	config Config
}

func NewConfigBuilder(config Config) (ConfigBuilder, error) {
	return ConfigBuilder{config: config}, nil
}

func (b ConfigBuilder) Build() Config {
	config := b.config

	var blacklist []string
	var whitelist []string

	copy(blacklist, b.config.Restrictions.BlackList)
	copy(whitelist, b.config.Restrictions.WhiteList)

	config.Restrictions.BlackList = blacklist
	config.Restrictions.WhiteList = whitelist

	return config
}
