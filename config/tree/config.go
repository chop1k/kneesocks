package tree

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"os"
)

type Config struct {
	Tcp      TcpConfig `validate:"required"`
	Udp      UdpConfig `validate:"required"`
	SocksV4  *SocksV4Config
	SocksV4a *SocksV4aConfig
	SocksV5  *SocksV5Config
	Log      LogConfig `validate:"required"`
}

func NewConfig(validate validator.Validate, path string) (Config, error) {
	return readConfig(validate, path)
}

func readConfig(validate validator.Validate, path string) (Config, error) {
	file, err := os.Open(path)

	if err != nil {
		return Config{}, err
	}

	return decodeConfig(validate, file)
}

func decodeConfig(validate validator.Validate, file *os.File) (Config, error) {
	decoder := *json.NewDecoder(file)

	config := Config{}

	err := decoder.Decode(&config)

	if err != nil {
		return Config{}, err
	}

	return validateConfig(validate, config)
}

func validateConfig(validate validator.Validate, config Config) (Config, error) {
	err := validate.Struct(config)

	if err != nil {
		return Config{}, err
	}

	return config, nil
}
