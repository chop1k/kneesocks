package v5

type Logger struct {
	Association  AssociationLogger
	Auth         AuthLogger
	Bind         BindLogger
	Connect      ConnectLogger
	Errors       ErrorsLogger
	Restrictions RestrictionsLogger
	Transfer     TransferLogger
}

func NewLogger(
	association AssociationLogger,
	auth AuthLogger,
	bind BindLogger,
	connect ConnectLogger,
	errors ErrorsLogger,
	restrictions RestrictionsLogger,
	transfer TransferLogger,
) (Logger, error) {
	return Logger{
		Association:  association,
		Auth:         auth,
		Bind:         bind,
		Connect:      connect,
		Errors:       errors,
		Restrictions: restrictions,
		Transfer:     transfer,
	}, nil
}
