package tree

type TcpConfig struct {
	BindIp           string `validate:"required,ip"`
	BindPort         uint16 `validate:"required"`
	BindZone         string
	ClientBufferSize uint        `validate:"required"`
	HostBufferSize   uint        `validate:"required"`
	Deadline         TcpDeadline `validate:"required"`
}

type TcpDeadline struct {
	Welcome  uint `validate:"required"`
	Exchange uint `validate:"required"`
}
