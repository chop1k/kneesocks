package tree

type UdpConfig struct {
	BindIp     string `validate:"required,ip"`
	BindPort   uint16 `validate:"required"`
	BindZone   string
	BufferSize uint `validate:"required"`
}
