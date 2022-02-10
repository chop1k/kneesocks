package tree

type UdpConfig struct {
	Bind     UdpBindConfig     `validate:"required"`
	Buffer   UdpBufferConfig   `validate:"required"`
	Deadline UdpDeadlineConfig `validate:"required"`
}

type UdpBindConfig struct {
	Address string `validate:"required"`
	Port    uint16 `validate:"required"`
}

type UdpBufferConfig struct {
	PacketSize uint `validate:"required"`
}

type UdpDeadlineConfig struct {
	Read uint `validate:"required"`
}
