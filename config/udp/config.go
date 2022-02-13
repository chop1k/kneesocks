package udp

import "time"

type Config struct {
	Bind     BindConfig
	Buffer   BufferConfig
	Deadline DeadlineConfig
}

type BindConfig struct {
	Address string
	Port    uint16
}

type BufferConfig struct {
	PacketSize uint
}

type DeadlineConfig struct {
	Read time.Duration
}
