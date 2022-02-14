package tree

type SocksV4aConfig struct {
	AllowConnect bool
	AllowBind    bool
	Deadline     SocksV4aDeadlineConfig `validate:"required"`
	Restrictions Restrictions           `validate:"required"`
}

type SocksV4aDeadlineConfig struct {
	Response int `validate:"required"`
	Connect  int `validate:"required"`
	Bind     int `validate:"required"`
}
