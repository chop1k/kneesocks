package tree

type TcpConfig struct {
	BindIp           string `validate:"required,ip"`
	BindPort         uint   `validate:"required"`
	BindZone         string
	ClientBufferSize uint `validate:"required"`
	HostBufferSize   uint `validate:"required"`
}
