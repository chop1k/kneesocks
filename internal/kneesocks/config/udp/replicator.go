package udp

type ConfigReplicator struct {
	buffer   BufferConfig
	deadline DeadlineConfig
}

func NewConfigReplicator(
	buffer BufferConfig,
	deadline DeadlineConfig,
) (ConfigReplicator, error) {
	return ConfigReplicator{
		buffer:   buffer,
		deadline: deadline,
	}, nil
}

func (c ConfigReplicator) CopyBuffer() BufferConfig {
	return BufferConfig{
		PacketSize: c.buffer.PacketSize,
	}
}

func (c ConfigReplicator) CopyDeadline() DeadlineConfig {
	return DeadlineConfig{
		Read: c.deadline.Read,
	}
}
