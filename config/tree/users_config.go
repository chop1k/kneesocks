package tree

type User struct {
	Name     string `validate:"required,max=255"`
	Password string `validate:"required,max=255"`
}
