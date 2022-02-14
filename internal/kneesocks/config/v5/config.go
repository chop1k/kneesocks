package v5

import (
	"socks/internal/kneesocks/config/tree"
	"time"
)

type Config struct {
	AllowConnect                 bool
	AllowBind                    bool
	AllowUdpAssociation          bool
	AllowIPv4                    bool
	AllowIPv6                    bool
	AllowDomain                  bool
	AuthenticationMethodsAllowed []string
	Deadline                     DeadlineConfig
	Users                        map[string]User
}

type User struct {
	Password     string
	Restrictions tree.Restrictions
}

type DeadlineConfig struct {
	Selection        time.Duration
	Password         time.Duration
	PasswordResponse time.Duration
	Request          time.Duration
	Response         time.Duration
	Connect          time.Duration
	Bind             time.Duration
}
