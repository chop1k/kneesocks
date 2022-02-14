package tcp

type ConfigReplicator struct {
	deadline DeadlineConfig
}

func NewConfigReplicator(deadline DeadlineConfig) (ConfigReplicator, error) {
	return ConfigReplicator{deadline: deadline}, nil
}

func (c ConfigReplicator) CopyDeadline() DeadlineConfig {
	return DeadlineConfig{
		Welcome:  c.deadline.Welcome,
		Exchange: c.deadline.Exchange,
	}
}
