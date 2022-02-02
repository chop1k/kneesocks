package tree

type SocksV4Config struct {
	AllowConnect bool
	AllowBind    bool
	Deadline     SocksV4DeadlineConfig `validate:"required"`
	Restrictions Restrictions          `validate:"required"`
}

type SocksV4DeadlineConfig struct {
	Response uint `validate:"required"`
	Connect  uint `validate:"required"`
	Bind     uint `validate:"required"`
	Transfer uint `validate:"required"`
}
