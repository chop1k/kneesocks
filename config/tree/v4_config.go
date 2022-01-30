package tree

type SocksV4Config struct {
	AllowConnect    bool
	AllowBind       bool
	ConnectDeadline uint         `validate:"required"`
	BindDeadline    uint         `validate:"required"`
	Restrictions    Restrictions `validate:"required"`
}
