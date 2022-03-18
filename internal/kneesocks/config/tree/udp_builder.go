package tree

import (
	"encoding/json"
	"os"

	"github.com/go-playground/validator/v10"
)

type UdpBuilder struct {
	validate validator.Validate
}

func NewUdpBuilder(validate validator.Validate) (UdpBuilder, error) {
	return UdpBuilder{
		validate: validate,
	}, nil
}

func (b UdpBuilder) Build(file *os.File) (UdpConfig, error) {
	decoder := *json.NewDecoder(file)

	config := UdpConfig{}

	err := decoder.Decode(&config)

	if err != nil {
		return UdpConfig{}, err
	}

	err = b.validate.Struct(config)

	if err != nil {
		return UdpConfig{}, err
	}

	return config, nil
}
