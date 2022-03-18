package tree

import (
	"encoding/json"
	"os"

	"github.com/go-playground/validator/v10"
)

type SocksV5Builder struct {
	validate validator.Validate
}

func NewSocksV5Builder(validate validator.Validate) (SocksV5Builder, error) {
	return SocksV5Builder{
		validate: validate,
	}, nil
}

func (b SocksV5Builder) Build(file *os.File) (SocksV5Config, error) {
	decoder := *json.NewDecoder(file)

	config := SocksV5Config{}

	err := decoder.Decode(&config)

	if err != nil {
		return SocksV5Config{}, err
	}

	err = b.validate.Struct(config)

	if err != nil {
		return SocksV5Config{}, err
	}

	return config, nil
}
