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

type Restrictions struct {
	WhiteList []string `validate:"required"`
	BlackList []string `validate:"required"`
}

type SocksV5DeadlineConfig struct {
	Methods          uint `validate:"required"`
	Selection        uint `validate:"required"`
	Password         uint `validate:"required"`
	PasswordResponse uint `validate:"required"`
	Request          uint `validate:"required"`
	Response         uint `validate:"required"`
	Connect          uint `validate:"required"`
	Bind             uint `validate:"required"`
	Transfer         uint `validate:"required"`
}
