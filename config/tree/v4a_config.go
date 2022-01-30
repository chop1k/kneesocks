package tree

type SocksV4aConfig struct {
	AllowConnect    bool
	AllowBind       bool
	AllowIPv4       bool
	ConnectDeadline uint         `validate:"required"`
	BindDeadline    uint         `validate:"required"`
	Restrictions    Restrictions `validate:"required"`
}
