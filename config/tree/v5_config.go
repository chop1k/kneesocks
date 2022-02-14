package tree

type SocksV5Config struct {
	AllowConnect                 bool
	AllowBind                    bool
	AllowUdpAssociation          bool
	AllowIPv4                    bool
	AllowIPv6                    bool
	AllowDomain                  bool
	AuthenticationMethodsAllowed []string              `validate:"required"`
	Deadline                     SocksV5DeadlineConfig `validate:"required"`
	Users                        map[string]User       `validate:"required,dive,keys,max=255,endkeys"`
}

type User struct {
	Password     string       `validate:"required,max=255"`
	Restrictions Restrictions `validate:"required"`
}

type SocksV5DeadlineConfig struct {
	Selection        int `validate:"required"`
	Password         int `validate:"required"`
	PasswordResponse int `validate:"required"`
	Request          int `validate:"required"`
	Response         int `validate:"required"`
	Connect          int `validate:"required"`
	Bind             int `validate:"required"`
}
