package v4a

type Config struct {
	Connect []ConnectScope `validate:"required"`
	Bind    []BindScope    `validate:"required"`
}

type ConnectScope struct {
}

type BindScope struct {
}
