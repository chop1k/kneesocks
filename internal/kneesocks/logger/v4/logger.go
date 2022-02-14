package v4

type Logger struct {
	Bind         BindLogger
	Connect      ConnectLogger
	Errors       ErrorsLogger
	Restrictions RestrictionsLogger
	Transfer     TransferLogger
}

func NewLogger(
	bind BindLogger,
	connect ConnectLogger,
	errors ErrorsLogger,
	restrictions RestrictionsLogger,
	transfer TransferLogger,
) (Logger, error) {
	return Logger{
		Bind:         bind,
		Connect:      connect,
		Errors:       errors,
		Restrictions: restrictions,
		Transfer:     transfer,
	}, nil
}
