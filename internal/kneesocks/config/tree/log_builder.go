package tree

import (
	"encoding/json"
	"os"

	"github.com/go-playground/validator/v10"
)

type LogBuilder struct {
	validate validator.Validate
}

func NewLogBuilder(validator validator.Validate) (LogBuilder, error) {
	return LogBuilder{
		validate: validator,
	}, nil
}

func (b LogBuilder) Build(file *os.File) (LogConfig, error) {
	decoder := *json.NewDecoder(file)

	config := LogConfig{}

	err := decoder.Decode(&config)

	if err != nil {
		return LogConfig{}, err
	}

	err = b.validate.Struct(config)

	if err != nil {
		return LogConfig{}, err
	}

	return config, nil
}
