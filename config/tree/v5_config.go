package tree

type SocksV5Config struct {
	AllowConnect                 bool
	AllowBind                    bool
	AllowUdpAssociation          bool
	AllowIPv4                    bool
	AllowIPv6                    bool
	AllowDomain                  bool
	AuthenticationMethodsAllowed []string        `validate:"required"`
	ConnectDeadline              uint            `validate:"required"`
	BindDeadline                 uint            `validate:"required"`
	Users                        map[string]User `validate:"required,dive,keys,max=255,endkeys"`
}

type User struct {
	Password     string       `validate:"required,max=255"`
	Restrictions Restrictions `validate:"required"`
}

type Restrictions struct {
	WhiteList []string `validate:"required"`
	BlackList []string `validate:"required"`
}
