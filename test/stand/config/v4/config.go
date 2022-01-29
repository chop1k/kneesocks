package v4

type Config struct {
	Connect []ConnectScope `validate:"required"`
	Bind    []BindScope    `validate:"required"`
}

type ConnectScope struct {
	Picture byte `validate:"required"`
}

type BindScope struct {
	Port    uint16 `validate:"required"`
	Picture byte   `validate:"required"`
}
