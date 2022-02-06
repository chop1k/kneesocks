package tree

type Restrictions struct {
	WhiteList []string         `validate:"required"`
	BlackList []string         `validate:"required"`
	Rate      RateRestrictions `validate:"required"`
}

type RateRestrictions struct {
	MaxSimultaneousConnections  int `validate:"required"`
	HostReadBuffersPerSecond    int `validate:"required"`
	HostWriteBuffersPerSecond   int `validate:"required"`
	ClientReadBuffersPerSecond  int `validate:"required"`
	ClientWriteBuffersPerSecond int `validate:"required"`
}
