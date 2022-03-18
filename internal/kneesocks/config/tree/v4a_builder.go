package tree

import (
	"encoding/json"
	"os"

	"github.com/go-playground/validator/v10"
)

type SocksV4aBuilder struct {
	validate validator.Validate
}

func NewSocksV4aBuilder(validate validator.Validate) (SocksV4aBuilder, error) {
	return SocksV4aBuilder{
		validate: validate,
	}, nil
}

func (b SocksV4aBuilder) Build(file *os.File) (SocksV4aConfig, error) {
	decoder := *json.NewDecoder(file)

	config := SocksV4aConfig{}

	err := decoder.Decode(&config)

	if err != nil {
		return SocksV4aConfig{}, err
	}

	err = b.validate.Struct(config)

	if err != nil {
		return SocksV4aConfig{}, err
	}

	return config, nil
}
