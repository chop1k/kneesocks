package v5

type Config struct {
	Connect   []ConnectScope   `validate:"required"`
	Bind      []BindScope      `validate:"required"`
	Auth      []AuthScope      `validate:"required"`
	Associate []AssociateScope `validate:"required"`
}

type ConnectScope struct {
}

type BindScope struct {
}

type AuthScope struct {
}

type AssociateScope struct {
}
