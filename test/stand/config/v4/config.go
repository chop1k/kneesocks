package v4

type Config struct {
	Connect []ConnectScope `validate:"required"`
	Bind    []BindScope    `validate:"required"`
}

type ConnectScope struct {
}

type BindScope struct {
}
