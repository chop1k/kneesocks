package v5

type ConfigReplicator struct {
	config Config
}

func (b ConfigReplicator) Copy() Config {
	config := b.config

	methods := make([]string, len(b.config.AuthenticationMethodsAllowed))

	copy(methods, b.config.AuthenticationMethodsAllowed)

	config.AuthenticationMethodsAllowed = methods

	users := make(map[string]User)

	for name, user := range b.config.Users {
		_user := user

		blacklist := make([]string, len(user.Restrictions.BlackList))
		whitelist := make([]string, len(user.Restrictions.WhiteList))

		copy(blacklist, user.Restrictions.BlackList)
		copy(whitelist, user.Restrictions.WhiteList)

		_user.Restrictions.BlackList = blacklist
		_user.Restrictions.WhiteList = whitelist

		users[name] = _user
	}

	config.Users = users

	return config
}
