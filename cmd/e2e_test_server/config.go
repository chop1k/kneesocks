package main

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"os"
)

type Config struct {
	Socks   SocksConfig
	Tcp     TcpConfig
	Udp     UdpConfig
	Picture PictureConfig
	Log     LogConfig
}

type SocksConfig struct {
	IPv4 string `validate:"required,ipv4"`
	IPv6 string `validate:"required,ipv6"`
	Port uint16 `validate:"required"`
	Zone string
}

type TcpConfig struct {
	BindIP   string `validate:"required,ip"`
	BindPort uint16 `validate:"required"`
	BindZone string
}

type UdpConfig struct {
	BindIP   string `validate:"required,ip"`
	BindPort uint16 `validate:"required"`
	BindZone string
}

type PictureConfig struct {
	BigPicturePath    string `validate:"required,uri"`
	MiddlePicturePath string `validate:"required,uri"`
	SmallPicturePath  string `validate:"required,uri"`
}

type LogConfig struct {
	Path string `validate:"required"`
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
