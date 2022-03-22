package v4

import (
	"socks/internal/kneesocks/config/tree"
	"time"
)

type Handler struct {
}

func NewHandler() (Handler, error) {
	return Handler{}, nil
}

func (h Handler) Handle(raw tree.SocksV4Config) Config {
	return Config{
		AllowConnect: raw.AllowConnect,
		AllowBind:    raw.AllowBind,
		Deadline: DeadlineConfig{
			Response: time.Second * time.Duration(raw.Deadline.Response),
			Connect:  time.Second * time.Duration(raw.Deadline.Connect),
			Bind:     time.Second * time.Duration(raw.Deadline.Bind),
		},
		Restrictions: tree.Restrictions{
			WhiteList: raw.Restrictions.WhiteList,
			BlackList: raw.Restrictions.BlackList,
			Rate: tree.RateRestrictions{
				MaxSimultaneousConnections:  raw.Restrictions.Rate.MaxSimultaneousConnections,
				HostReadBuffersPerSecond:    raw.Restrictions.Rate.HostReadBuffersPerSecond,
				HostWriteBuffersPerSecond:   raw.Restrictions.Rate.HostWriteBuffersPerSecond,
				ClientReadBuffersPerSecond:  raw.Restrictions.Rate.ClientReadBuffersPerSecond,
				ClientWriteBuffersPerSecond: raw.Restrictions.Rate.ClientWriteBuffersPerSecond,
			},
		},
	}
}
