package tree

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"os"
)

type Config struct {
	Tcp       TcpConfig `validate:"required"`
	Udp       UdpConfig `validate:"required"`
	SocksV4   *SocksV4Config
	SocksV4a  *SocksV4aConfig
	SocksV5   *SocksV5Config
	Log       LogConfig `validate:"required"`
	WhiteList []string  `validate:"required,dive,ip|hostname"`
	BlackList []string  `validate:"required,dive,ip|hostname"`
	Users     []User    `validate:"required"`
	Unix      *UnixConfig
}

type TcpConfig struct {
	BindIp           string `validate:"required,ip"`
	BindPort         uint   `validate:"required"`
	BindZone         string
	ClientBufferSize uint `validate:"required"`
	HostBufferSize   uint `validate:"required"`
}

type UdpConfig struct {
	BindIp     string `validate:"required,ip"`
	BindPort   uint   `validate:"required"`
	BindZone   string
	BufferSize uint `validate:"required"`
}

type SocksV4Config struct {
	AllowConnect    bool
	AllowBind       bool
	ConnectDeadline uint `validate:"required"`
	BindDeadline    uint `validate:"required"`
}

type SocksV4aConfig struct {
	AllowConnect    bool
	AllowBind       bool
	ConnectDeadline uint `validate:"required"`
	BindDeadline    uint `validate:"required"`
}

type SocksV5Config struct {
	AllowConnect                 bool
	AllowBind                    bool
	AllowUdpAssociation          bool
	AllowIPv4                    bool
	AllowIPv6                    bool
	AllowDomain                  bool
	AuthenticationMethodsAllowed []string `validate:"required"`
	ConnectDeadline              uint     `validate:"required"`
	BindDeadline                 uint     `validate:"required"`
}

type LogConfig struct {
	Formats FormatsConfig
	Loggers LoggersConfig
}

type FormatsConfig struct {
	Replacer string `validate:"required"`
}

type LoggersConfig struct {
	Tcp      *TcpLoggerConfig
	Udp      *UdpLoggerConfig
	SocksV4  *SocksV4LoggerConfig
	SocksV4a *SocksV4aLoggerConfig
	SocksV5  *SocksV5LoggerConfig
	Unix     *UnixLoggerConfig
	Errors   *ErrorsLoggerConfig
}

type LogOutputsConfig struct {
	Console *ConsoleOutputConfig
	File    *FileOutputConfig
}

type ConsoleOutputConfig struct {
}

type FileOutputConfig struct {
	Path string
}

type TcpLoggerConfig struct {
	Formats TcpLoggerFormats
	Outputs LogOutputsConfig
}

type TcpLoggerFormats struct {
	ConnectionAccepted           string `validate:"required"`
	ConnectionDenied             string `validate:"required"`
	ConnectionProtocolDetermined string `validate:"required"`
	ConnectionBound              string `validate:"required"`
	Listen                       string `validate:"required"`
}

type UdpLoggerConfig struct {
	Formats UdpLoggerFormats
	Outputs LogOutputsConfig
}

type UdpLoggerFormats struct {
	PacketReceived string `validate:"required"`
	PacketDenied   string `validate:"required"`
	PacketSent     string `validate:"required"`
}

type SocksV4LoggerConfig struct {
	Formats SocksV4LoggerFormats
	Outputs LogOutputsConfig
}

type SocksV4LoggerFormats struct {
	ConnectRequest    string `validate:"required"`
	ConnectFailed     string `validate:"required"`
	ConnectSuccessful string `validate:"required"`
	BindRequest       string `validate:"required"`
	BindFailed        string `validate:"required"`
	BindSuccessful    string `validate:"required"`
	Bound             string `validate:"required"`
}

type SocksV4aLoggerConfig struct {
	Formats SocksV4aLoggerFormats
	Outputs LogOutputsConfig
}

type SocksV4aLoggerFormats struct {
	SocksV4LoggerFormats
}

type SocksV5LoggerConfig struct {
	Formats SocksV5LoggerFormats
	Outputs LogOutputsConfig
}

type SocksV5LoggerFormats struct {
	SocksV4LoggerFormats
	AuthenticationSuccessful string `validate:"required"`
	AuthenticationFailed     string `validate:"required"`
	UdpAssociationRequest    string `validate:"required"`
	UdpAssociationSuccessful string `validate:"required"`
	UdpAssociationFailed     string `validate:"required"`
}

type UnixLoggerConfig struct {
	Formats UnixLoggerFormats
	Outputs LogOutputsConfig
}

type UnixLoggerFormats struct {
}

type ErrorsLoggerConfig struct {
	Formats ErrorsLoggerFormats
	Outputs LogOutputsConfig
}

type ErrorsLoggerFormats struct {
	ReadFromClientFailed string `validate:"required"`
	WriteToClientFailed  string `validate:"required"`
	ReadFromHostFailed   string `validate:"required"`
	WriteToHostFailed    string `validate:"required"`
	AddressParsingFailed string `validate:"required"`
}

type User struct {
	Name     string `validate:"required,max=255"`
	Password string `validate:"required,max=255"`
}

type UnixConfig struct {
	Enable                       bool
	AuthenticationMethodsAllowed []string `validate:"required"`
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
