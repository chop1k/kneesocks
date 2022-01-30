package v5

type Config struct {
	Connect   []ConnectScope   `validate:"required"`
	Bind      []BindScope      `validate:"required"`
	Auth      []AuthScope      `validate:"required"`
	Associate []AssociateScope `validate:"required"`
}

type ConnectScope struct {
	Picture     byte `validate:"required"`
	AddressType byte `validate:"required"`
}

type BindScope struct {
	Port        uint16 `validate:"required"`
	Picture     byte   `validate:"required"`
	AddressType byte   `validate:"required"`
}

type AuthScope struct {
	Picture     byte `validate:"required"`
	AddressType byte `validate:"required"`
	Method      byte
}

type AssociateScope struct {
	Picture     byte `validate:"required"`
	AddressType byte `validate:"required"`
}
