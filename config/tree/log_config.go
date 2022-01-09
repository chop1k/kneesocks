package tree

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
