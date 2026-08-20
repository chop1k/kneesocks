package tcp

type ConfigReplicator struct {
	deadline DeadlineConfig
}

func NewConfigReplicator(deadline DeadlineConfig) ConfigReplicator {
	return ConfigReplicator{deadline: deadline}
}

func (c ConfigReplicator) CopyDeadline() DeadlineConfig {
	return DeadlineConfig{
		Welcome:  c.deadline.Welcome,
		Exchange: c.deadline.Exchange,
	}
}
