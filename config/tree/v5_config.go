package tree

type SocksV5Config struct {
	AllowConnect                 bool
	AllowBind                    bool
	AllowUdpAssociation          bool
	AllowIPv4                    bool
	AllowIPv6                    bool
	AllowDomain                  bool
	AuthenticationMethodsAllowed []string `validate:"required"`
	ConnectDeadline              uint     `validate:"required"`
	BindDeadline                 uint     `validate:"required"`
}
