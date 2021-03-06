package tree

type TcpConfig struct {
	Bind     TcpBindConfig   `validate:"required"`
	Buffer   TcpBufferConfig `validate:"required"`
	Deadline TcpDeadline     `validate:"required"`
}

type TcpBindConfig struct {
	Address string
	Port    uint16 `validate:"required"`
}

type TcpBufferConfig struct {
	ClientSize uint `validate:"required"`
	HostSize   uint `validate:"required"`
}

type TcpDeadline struct {
	Welcome  int `validate:"required"`
	Exchange int `validate:"required"`
}
