package tree

type TcpConfig struct {
	Bind     TcpBindConfig   `validate:"required"`
	Buffer   TcpBufferConfig `validate:"required"`
	Deadline TcpDeadline     `validate:"required"`
}

type TcpBindConfig struct {
	Address string `validate:"required"`
	Port    uint16 `validate:"required"`
}

type TcpBufferConfig struct {
	ClientSize uint `validate:"required"`
	HostSize   uint `validate:"required"`
}

type TcpDeadline struct {
	Welcome  uint `validate:"required"`
	Exchange uint `validate:"required"`
}
