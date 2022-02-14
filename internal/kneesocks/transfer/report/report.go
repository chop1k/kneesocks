package report

import "time"

type Report struct {
	Start     time.Time
	End       time.Time
	Bandwidth Bandwidth
	Rate      Rate
}

type Bandwidth struct {
}

type Rate struct {
	AverageSpeed uint
	MaxSpeed     uint
	MinSpeed     uint
	HostLimit    int
	ClientLimit  int
}
