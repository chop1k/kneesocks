package tcp

import "socks/config/tree"

type DeadlineConfig interface {
	GetExchangeDeadline() uint
	GetWelcomeDeadline() uint
}

type BaseDeadlineConfig struct {
	tree tree.Config
}

func NewBaseDeadlineConfig(tree tree.Config) (BaseDeadlineConfig, error) {
	return BaseDeadlineConfig{tree: tree}, nil
}

func (b BaseDeadlineConfig) GetExchangeDeadline() uint {
	return b.tree.Tcp.Deadline.Exchange
}

func (b BaseDeadlineConfig) GetWelcomeDeadline() uint {
	return b.tree.Tcp.Deadline.Welcome
}
