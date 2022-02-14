package tree

type SocksV4Config struct {
	AllowConnect bool
	AllowBind    bool
	Deadline     SocksV4DeadlineConfig `validate:"required"`
	Restrictions Restrictions          `validate:"required"`
}

type SocksV4DeadlineConfig struct {
	Response int `validate:"required"`
	Connect  int `validate:"required"`
	Bind     int `validate:"required"`
}
