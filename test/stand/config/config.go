package config

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"os"
	"socks/test/stand/config/v4"
	"socks/test/stand/config/v4a"
	"socks/test/stand/config/v5"
	"testing"
)

type Config struct {
	V4      v4.Config     `validate:"required"`
	V4a     v4a.Config    `validate:"required"`
	V5      v5.Config     `validate:"required"`
	Socks   SocksConfig   `validate:"required"`
	Server  ServerConfig  `validate:"required"`
	Picture PictureConfig `validate:"required"`
	Misc    MiscConfig    `validate:"required"`
	User    UserConfig    `validate:"required"`
}

type SocksConfig struct {
	IPv4    string `validate:"required,ipv4"`
	Domain  string `validate:"required,hostname"`
	IPv6    string `validate:"required,ipv6"`
	TcpPort uint16 `validate:"required"`
	UdpPort uint16 `validate:"required"`
}

type ServerConfig struct {
	IPv4        string `validate:"required,ipv4"`
	Domain      string `validate:"required,hostname"`
	IPv6        string `validate:"required,ipv6"`
	ConnectPort uint16 `validate:"required"`
	BindPort    uint16 `validate:"required"`
	UdpPort     uint16 `validate:"required"`
}

type PictureConfig struct {
	BigPictureHash    string `validate:"required"`
	MiddlePictureHash string `validate:"required"`
	SmallPictureHash  string `validate:"required"`
}

type MiscConfig struct {
	TempDirPath         string `validate:"required,uri"`
	TempFileNamePattern string `validate:"required"`
	RandomSuffixLength  uint   `validate:"required"`
}

type UserConfig struct {
	Name     string `validate:"required"`
	Password string `validate:"required"`
}

func NewConfig(validate validator.Validate, path string, t *testing.T) Config {
	return readConfig(validate, path, t)
}

func readConfig(validate validator.Validate, path string, t *testing.T) Config {
	file, err := os.Open(path)

	require.NoError(t, err)

	return decodeConfig(validate, file, t)
}

func decodeConfig(validate validator.Validate, file *os.File, t *testing.T) Config {
	decoder := *json.NewDecoder(file)

	config := Config{}

	err := decoder.Decode(&config)

	require.NoError(t, err)

	return validateConfig(validate, config, t)
}

func validateConfig(validate validator.Validate, config Config, t *testing.T) Config {
	err := validate.Struct(config)

	require.NoError(t, err)

	return config
}
