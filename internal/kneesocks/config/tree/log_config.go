package tree

type LogConfig struct {
	Tcp      *TcpLoggerConfig
	Udp      *UdpLoggerConfig
	SocksV4  *SocksV4LoggerConfig
	SocksV4a *SocksV4aLoggerConfig
	SocksV5  *SocksV5LoggerConfig
}

type ConsoleOutputConfig struct {
	TimeFormat string `validate:"required"`
}

type FileOutputConfig struct {
	Path string `validate:"uri"`
}

type TcpLoggerConfig struct {
	Level   int
	Console *ConsoleOutputConfig
	File    *FileOutputConfig
}

type UdpLoggerConfig struct {
	Level   int
	Console *ConsoleOutputConfig
	File    *FileOutputConfig
}

type SocksV4LoggerConfig struct {
	Level   int
	Console *ConsoleOutputConfig
	File    *FileOutputConfig
}

type SocksV4aLoggerConfig struct {
	Level   int
	Console *ConsoleOutputConfig
	File    *FileOutputConfig
}

type SocksV5LoggerConfig struct {
	Level   int
	Console *ConsoleOutputConfig
	File    *FileOutputConfig
}
