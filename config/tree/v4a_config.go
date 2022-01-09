package tree

type SocksV4aConfig struct {
	AllowConnect    bool
	AllowBind       bool
	ConnectDeadline uint `validate:"required"`
	BindDeadline    uint `validate:"required"`
}
