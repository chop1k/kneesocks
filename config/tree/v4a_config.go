package tree

type SocksV4aConfig struct {
	AllowConnect bool
	AllowBind    bool
	Deadline     SocksV4aDeadlineConfig `validate:"required"`
	Restrictions Restrictions           `validate:"required"`
}

type SocksV4aDeadlineConfig struct {
	Response uint `validate:"required"`
	Connect  uint `validate:"required"`
	Bind     uint `validate:"required"`
	Transfer uint `validate:"required"`
}
