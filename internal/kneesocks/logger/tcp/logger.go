package tcp

type Logger struct {
	Connection ConnectionLogger
	Errors     ErrorsLogger
	Listen     ListenLogger
}

func NewLogger(
	connection ConnectionLogger,
	errors ErrorsLogger,
	listen ListenLogger,
) (Logger, error) {
	return Logger{
		Connection: connection,
		Errors:     errors,
		Listen:     listen,
	}, nil
}
