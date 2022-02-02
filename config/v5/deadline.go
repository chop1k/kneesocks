package v5

import "socks/config/tree"

type DeadlineConfig interface {
	GetRequestDeadline() uint
	GetConnectDeadline() uint
	GetBindDeadline() uint
	GetPasswordDeadline() uint
}

type BaseDeadlineConfig struct {
	tree tree.Config
}

func NewBaseDeadlineConfig(tree tree.Config) (BaseDeadlineConfig, error) {
	return BaseDeadlineConfig{tree: tree}, nil
}

func (b BaseDeadlineConfig) GetRequestDeadline() uint {
	return b.tree.SocksV5.Deadline.Request
}

func (b BaseDeadlineConfig) GetConnectDeadline() uint {
	return b.tree.SocksV5.Deadline.Connect
}

func (b BaseDeadlineConfig) GetBindDeadline() uint {
	return b.tree.SocksV5.Deadline.Bind
}

func (b BaseDeadlineConfig) GetPasswordDeadline() uint {
	return b.tree.SocksV5.Deadline.Password
}
