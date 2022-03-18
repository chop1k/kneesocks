package tree

import (
	"encoding/json"
	"os"

	"github.com/go-playground/validator/v10"
)

type SocksV4Builder struct {
	validate validator.Validate
}

func NewSocksV4Builder(validate validator.Validate) (SocksV4Builder, error) {
	return SocksV4Builder{
		validate: validate,
	}, nil
}

func (b SocksV4Builder) Build(file *os.File) (SocksV4Config, error) {
	decoder := *json.NewDecoder(file)

	config := SocksV4Config{}

	err := decoder.Decode(&config)

	if err != nil {
		return SocksV4Config{}, err
	}

	err = b.validate.Struct(config)

	if err != nil {
		return SocksV4Config{}, err
	}

	return config, nil
}
