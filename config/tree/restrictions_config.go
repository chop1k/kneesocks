package tree

type Restrictions struct {
	WhiteList []string         `validate:"required"`
	BlackList []string         `validate:"required"`
	Rate      RateRestrictions `validate:"required"`
}

type RateRestrictions struct {
	MaxSimultaneousConnections  uint `validate:"required"`
	HostReadBuffersPerSecond    uint `validate:"required"`
	HostWriteBuffersPerSecond   uint `validate:"required"`
	ClientReadBuffersPerSecond  uint `validate:"required"`
	ClientWriteBuffersPerSecond uint `validate:"required"`
}
