package v5

import (
	"socks/internal/kneesocks/config/tree"
	"time"
)

type Handler struct {
}

func NewHandler() (Handler, error) {
	return Handler{}, nil
}

func (h Handler) Handle(raw tree.SocksV5Config) Config {
	users := make(map[string]User)

	for name, user := range raw.Users {
		users[name] = User{
			Password:     user.Password,
			Restrictions: user.Restrictions,
		}
	}

	return Config{
		AllowConnect:                 raw.AllowConnect,
		AllowBind:                    raw.AllowBind,
		AllowUdpAssociation:          raw.AllowUdpAssociation,
		AllowIPv4:                    raw.AllowIPv4,
		AllowIPv6:                    raw.AllowIPv6,
		AllowDomain:                  raw.AllowDomain,
		AuthenticationMethodsAllowed: raw.AuthenticationMethodsAllowed,
		Deadline: DeadlineConfig{
			Selection:        time.Second * time.Duration(raw.Deadline.Selection),
			Password:         time.Second * time.Duration(raw.Deadline.Password),
			PasswordResponse: time.Second * time.Duration(raw.Deadline.PasswordResponse),
			Request:          time.Second * time.Duration(raw.Deadline.Request),
			Response:         time.Second * time.Duration(raw.Deadline.Response),
			Connect:          time.Second * time.Duration(raw.Deadline.Connect),
			Bind:             time.Second * time.Duration(raw.Deadline.Bind),
		},
		Users: users,
	}
}
