package main

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"os"
)

type Config struct {
	Socks   SocksConfig   `validate:"required"`
	Tcp     TcpConfig     `validate:"required"`
	Udp     UdpConfig     `validate:"required"`
	Picture PictureConfig `validate:"required"`
	Log     LogConfig     `validate:"required"`
	Http    HttpConfig    `validate:"required"`
}

type SocksConfig struct {
	IPv4 string `validate:"required,ipv4"`
	IPv6 string `validate:"required,ipv6"`
	Port uint16 `validate:"required"`
	Zone string
}

type TcpConfig struct {
	BindIPv4 string `validate:"required,ipv4"`
	BindIPv6 string `validate:"required,ipv6"`
	BindPort uint16 `validate:"required"`
	BindZone string
}

type UdpConfig struct {
	BindIPv4 string `validate:"required,ipv4"`
	BindIPv6 string `validate:"required,ipv6"`
	BindPort uint16 `validate:"required"`
	BindZone string
}

type PictureConfig struct {
	BigPicturePath    string `validate:"required,uri"`
	MiddlePicturePath string `validate:"required,uri"`
	SmallPicturePath  string `validate:"required,uri"`
}

type HttpConfig struct {
	Address string `validate:"required"`
	Port    uint16 `validate:"required"`
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
