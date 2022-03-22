package dependency

import (
	"socks/internal/kneesocks/dependency/build"

	"github.com/sarulabs/di"
)

func registerHandlers(builder di.Builder) {
	bindHandlerDef := di.Def{
		Name:  "bind_handler",
		Scope: di.App,
		Build: build.BindHandler,
	}

	connectionHandlerDef := di.Def{
		Name:  "connection_handler",
		Scope: di.App,
		Build: build.ConnectionHandler,
	}

	packetHandlerDef := di.Def{
		Name:  "packet_handler",
		Scope: di.App,
		Build: build.PacketHandler,
	}

	err := builder.Add(
		bindHandlerDef,
		connectionHandlerDef,
		packetHandlerDef,
	)

	if err != nil {
		panic(err)
	}

	registerV4Handlers(builder)
	registerV4aHandlers(builder)
	registerV5Handlers(builder)
}

func registerV4Handlers(builder di.Builder) {
	connectHandlerDef := di.Def{
		Name:  "v4_connect_handler",
		Scope: di.App,
		Build: build.V4ConnectHandler,
	}

	bindHandlerDef := di.Def{
		Name:  "v4_bind_handler",
		Scope: di.App,
		Build: build.V4BindHandler,
	}

	handlerDef := di.Def{
		Name:  "v4_handler",
		Scope: di.App,
		Build: build.V4Handler,
	}

	errorHandlerDef := di.Def{
		Name:  "v4_error_handler",
		Scope: di.App,
		Build: build.V4ErrorHandler,
	}

	err := builder.Add(
		connectHandlerDef,
		bindHandlerDef,
		handlerDef,
		errorHandlerDef,
	)

	if err != nil {
		panic(err)
	}

	registerV4Helpers(builder)
}

func registerV4Helpers(builder di.Builder) {
	cleanerDef := di.Def{
		Name:  "v4_cleaner",
		Scope: di.App,
		Build: build.V4Cleaner,
	}

	limiterDef := di.Def{
		Name:  "v4_limiter",
		Scope: di.App,
		Build: build.V4Limiter,
	}

	transmitterDef := di.Def{
		Name:  "v4_transmitter",
		Scope: di.App,
		Build: build.V4Transmitter,
	}

	validatorDef := di.Def{
		Name:  "v4_validator",
		Scope: di.App,
		Build: build.V4Validator,
	}

	err := builder.Add(
		cleanerDef,
		limiterDef,
		transmitterDef,
		validatorDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV4aHandlers(builder di.Builder) {
	connectHandlerDef := di.Def{
		Name:  "v4a_connect_handler",
		Scope: di.App,
		Build: build.V4aConnectHandler,
	}

	bindHandlerDef := di.Def{
		Name:  "v4a_bind_handler",
		Scope: di.App,
		Build: build.V4aBindHandler,
	}

	handlerDef := di.Def{
		Name:  "v4a_handler",
		Scope: di.App,
		Build: build.V4aHandler,
	}

	errorHandlerDef := di.Def{
		Name:  "v4a_error_handler",
		Scope: di.App,
		Build: build.V4aErrorHandler,
	}

	err := builder.Add(
		connectHandlerDef,
		bindHandlerDef,
		handlerDef,
		errorHandlerDef,
	)

	if err != nil {
		panic(err)
	}

	registerV4aHelpers(builder)
}

func registerV4aHelpers(builder di.Builder) {
	cleanerDef := di.Def{
		Name:  "v4a_cleaner",
		Scope: di.App,
		Build: build.V4aCleaner,
	}

	limiterDef := di.Def{
		Name:  "v4a_limiter",
		Scope: di.App,
		Build: build.V4aLimiter,
	}

	transmitterDef := di.Def{
		Name:  "v4a_transmitter",
		Scope: di.App,
		Build: build.V4aTransmitter,
	}

	validatorDef := di.Def{
		Name:  "v4a_validator",
		Scope: di.App,
		Build: build.V4aValidator,
	}

	err := builder.Add(
		cleanerDef,
		limiterDef,
		transmitterDef,
		validatorDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV5Handlers(builder di.Builder) {
	authenticationHandlerDef := di.Def{
		Name:  "authentication_handler",
		Scope: di.App,
		Build: build.AuthenticationHandler,
	}

	connectHandlerDef := di.Def{
		Name:  "v5_connect_handler",
		Scope: di.App,
		Build: build.V5ConnectHandler,
	}

	bindHandlerDef := di.Def{
		Name:  "v5_bind_handler",
		Scope: di.App,
		Build: build.V5BindHandler,
	}

	udpAssociationHandler := di.Def{
		Name:  "v5_udp_association_handler",
		Scope: di.App,
		Build: build.V5UdpAssociationHandler,
	}

	handlerDef := di.Def{
		Name:  "v5_handler",
		Scope: di.App,
		Build: build.V5Handler,
	}

	errorHandlerDef := di.Def{
		Name:  "v5_error_handler",
		Scope: di.App,
		Build: build.V5ErrorHandler,
	}

	err := builder.Add(
		authenticationHandlerDef,
		connectHandlerDef,
		bindHandlerDef,
		handlerDef,
		errorHandlerDef,
		udpAssociationHandler,
	)

	if err != nil {
		panic(err)
	}

	registerAuthenticators(builder)
	registerV5Helpers(builder)
}

func registerAuthenticators(builder di.Builder) {
	passwordAuthenticatorDef := di.Def{
		Name:  "password_authenticator",
		Scope: di.App,
		Build: build.PasswordAuthenticator,
	}

	noAuthAuthenticatorDef := di.Def{
		Name:  "no_auth_authenticator",
		Scope: di.App,
		Build: build.NoAuthAuthenticator,
	}

	err := builder.Add(
		passwordAuthenticatorDef,
		noAuthAuthenticatorDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV5Helpers(builder di.Builder) {
	cleanerDef := di.Def{
		Name:  "v5_cleaner",
		Scope: di.App,
		Build: build.V5Cleaner,
	}

	limiterDef := di.Def{
		Name:  "v5_limiter",
		Scope: di.App,
		Build: build.V5Limiter,
	}

	transmitterDef := di.Def{
		Name:  "v5_transmitter",
		Scope: di.App,
		Build: build.V5Transmitter,
	}

	validatorDef := di.Def{
		Name:  "v5_validator",
		Scope: di.App,
		Build: build.V5Validator,
	}

	err := builder.Add(
		cleanerDef,
		limiterDef,
		transmitterDef,
		validatorDef,
	)

	if err != nil {
		panic(err)
	}
}
