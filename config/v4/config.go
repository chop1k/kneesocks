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

//
//type Config interface {
//	IsConnectAllowed() (bool, error)
//	IsBindAllowed() (bool, error)
//}
//
//type BaseConfig struct {
//	config gabs.Container
//}
//
//func NewBaseConfig(config gabs.Container) (BaseConfig, error) {
//	return BaseConfig{config: config}, nil
//}
//
//func (b BaseConfig) IsConnectAllowed() (bool, error) {
//	allowed, ok := b.config.Path("SocksV4.AllowConnect").Data().(bool)
//
//	if !ok {
//		return false, errors.New("SocksV4.AllowConnect: Not specified or have invalid type. ")
//	}
//
//	return allowed, nil
//}
//
//func (b BaseConfig) IsBindAllowed() (bool, error) {
//	allowed, ok := b.config.Path("SocksV4.AllowBind").Data().(bool)
//
//	if !ok {
//		return false, errors.New("SocksV4.AllowBind: Not specified or have invalid type. ")
//	}
//
//	return allowed, nil
//}
