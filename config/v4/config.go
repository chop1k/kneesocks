package v4

import (
	"socks/config/tree"
	"time"
)

type Config struct {
	AllowConnect bool
	AllowBind    bool
	Deadline     DeadlineConfig
	Restrictions tree.Restrictions
}

type DeadlineConfig struct {
	Response time.Duration
	Connect  time.Duration
	Bind     time.Duration
}
