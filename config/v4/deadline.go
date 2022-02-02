package v4

import "socks/config/tree"

type DeadlineConfig interface {
	GetConnectDeadline() uint
	GetBindDeadline() uint
}

type BaseDeadlineConfig struct {
	tree tree.Config
}

func NewBaseDeadlineConfig(tree tree.Config) (BaseDeadlineConfig, error) {
	return BaseDeadlineConfig{tree: tree}, nil
}

func (b BaseDeadlineConfig) GetConnectDeadline() uint {
	return b.tree.SocksV4.Deadline.Connect
}

func (b BaseDeadlineConfig) GetBindDeadline() uint {
	return b.tree.SocksV4.Deadline.Bind
}
