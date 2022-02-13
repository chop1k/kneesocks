package tcp

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
	ClientSize uint
	HostSize   uint
}

type DeadlineConfig struct {
	Welcome  time.Duration
	Exchange time.Duration
}
