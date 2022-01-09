package tree

type UdpConfig struct {
	BindIp     string `validate:"required,ip"`
	BindPort   uint   `validate:"required"`
	BindZone   string
	BufferSize uint `validate:"required"`
}
