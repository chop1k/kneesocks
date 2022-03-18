package tree

import (
	"encoding/json"
	"os"

	"github.com/go-playground/validator/v10"
)

type TcpBuilder struct {
	validate validator.Validate
}

func NewTcpBuilder(validate validator.Validate) (TcpBuilder, error) {
	return TcpBuilder{
		validate: validate,
	}, nil
}

func (b TcpBuilder) Build(file *os.File) (TcpConfig, error) {
	decoder := *json.NewDecoder(file)

	config := TcpConfig{}

	err := decoder.Decode(&config)

	if err != nil {
		return TcpConfig{}, err
	}

	err = b.validate.Struct(config)

	if err != nil {
		return TcpConfig{}, err
	}

	return config, nil
}
