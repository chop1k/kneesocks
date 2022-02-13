package main

import (
	"github.com/sarulabs/di"
	"socks/config/tcp"
	udp2 "socks/config/udp"
	v43 "socks/config/v4"
	v4a3 "socks/config/v4a"
	v53 "socks/config/v5"
	"socks/handlers"
	v42 "socks/handlers/v4"
	helpers5 "socks/handlers/v4/helpers"
	v4a2 "socks/handlers/v4a"
	helpers2 "socks/handlers/v4a/helpers"
	v52 "socks/handlers/v5"
	"socks/handlers/v5/authenticator"
	"socks/handlers/v5/helpers"
	tcp2 "socks/logger/tcp"
	"socks/logger/udp"
	v44 "socks/logger/v4"
	v4a4 "socks/logger/v4a"
	v54 "socks/logger/v5"
	"socks/managers"
	helpers3 "socks/protocol"
	"socks/protocol/auth/password"
	v4 "socks/protocol/v4"
	"socks/protocol/v4a"
	v5 "socks/protocol/v5"
	"socks/transfer"
	"socks/utils"
)

func registerHandlers(builder di.Builder) {
	bindHandlerDef := di.Def{
		Name:  "bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			tcpLogger := ctn.Get("tcp_logger").(tcp2.Logger)
			bindHandler := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)

			return handlers.NewBaseBindHandler(addressUtils, tcpLogger, bindHandler, bindManager)
		},
	}

	connectionHandlerDef := di.Def{
		Name:  "connection_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			v4Handler := ctn.Get("v4_handler").(v42.Handler)
			v4aHandler := ctn.Get("v4a_handler").(v4a2.Handler)
			v5Handler := ctn.Get("v5_handler").(v52.Handler)
			tcpLogger := ctn.Get("tcp_logger").(tcp2.Logger)
			receiver := ctn.Get("receiver").(helpers3.Receiver)
			bindHandler := ctn.Get("bind_handler").(handlers.BindHandler)
			replicator := ctn.Get("tcp_config_replicator").(tcp.ConfigReplicator)

			return handlers.NewBaseConnectionHandler(
				v4Handler,
				v4aHandler,
				v5Handler,
				tcpLogger,
				receiver,
				bindHandler,
				replicator,
			)
		},
	}

	packetHandlerDef := di.Def{
		Name:  "packet_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("v5_parser").(v5.Parser)
			builder := ctn.Get("v5_builder").(v5.Builder)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			clients := ctn.Get("udp_client_manager").(managers.UdpClientManager)
			hosts := ctn.Get("udp_host_manager").(managers.UdpHostManager)
			logger := ctn.Get("udp_logger").(udp.Logger)
			replicator := ctn.Get("udp_config_replicator").(udp2.ConfigReplicator)

			return handlers.NewBasePacketHandler(
				parser,
				builder,
				addressUtils,
				clients,
				hosts,
				logger,
				replicator,
			), nil
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			v4Logger := ctn.Get("v4_logger").(v44.Logger)
			sender := ctn.Get("v4_sender").(v4.Sender)
			errorHandler := ctn.Get("v4_error_handler").(v42.ErrorHandler)
			transmitter := ctn.Get("v4_transmitter").(helpers5.Transmitter)

			return v42.NewBaseConnectHandler(
				v4Logger,
				sender,
				errorHandler,
				transmitter,
			)
		},
	}

	bindHandlerDef := di.Def{
		Name:  "v4_bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			v4Logger := ctn.Get("v4_logger").(v44.Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v4_sender").(v4.Sender)
			errorHandler := ctn.Get("v4_error_handler").(v42.ErrorHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			transmitter := ctn.Get("v4_transmitter").(helpers5.Transmitter)

			return v42.NewBaseBindHandler(
				v4Logger,
				addressUtils,
				sender,
				errorHandler,
				bindManager,
				transmitter,
			)
		},
	}

	handlerDef := di.Def{
		Name:  "v4_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("v4_parser").(v4.Parser)
			v4Logger := ctn.Get("v4_logger").(v44.Logger)
			connectHandler := ctn.Get("v4_connect_handler").(v42.ConnectHandler)
			bindHandler := ctn.Get("v4_bind_handler").(v42.BindHandler)
			sender := ctn.Get("v4_sender").(v4.Sender)
			errorHandler := ctn.Get("v4_error_handler").(v42.ErrorHandler)
			validator := ctn.Get("v4_validator").(helpers5.Validator)
			cleaner := ctn.Get("v4_cleaner").(helpers5.Cleaner)
			replicator := ctn.Get("v4_config_replicator").(v43.ConfigReplicator)

			return v42.NewBaseHandler(
				parser,
				v4Logger,
				connectHandler,
				bindHandler,
				sender,
				errorHandler,
				validator,
				cleaner,
				replicator,
			)
		},
	}

	errorHandlerDef := di.Def{
		Name:  "v4_error_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			sender := ctn.Get("v4_sender").(v4.Sender)
			v4Logger := ctn.Get("v4_logger").(v44.Logger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return v42.NewBaseErrorHandler(
				v4Logger,
				sender,
				errorUtils,
			)
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

			return helpers5.NewBaseCleaner(manager)
		},
	}

	limiterDef := di.Def{
		Name:  "v4_limiter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

			return helpers5.NewBaseLimiter(manager)
		},
	}

	transmitterDef := di.Def{
		Name:  "v4_transmitter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
			bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
			bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

			return helpers5.NewBaseTransmitter(connect, bind, bindRate)
		},
	}

	validatorDef := di.Def{
		Name:  "v4_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(managers.BlacklistManager)
			sender := ctn.Get("v4_sender").(v4.Sender)
			logger := ctn.Get("v4_logger").(v44.Logger)
			limiter := ctn.Get("v4_limiter").(helpers5.Limiter)

			return helpers5.NewBaseValidator(whitelist, blacklist, sender, logger, limiter)
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			v4aLogger := ctn.Get("v4a_logger").(v4a4.Logger)
			sender := ctn.Get("v4a_sender").(v4a.Sender)
			errorHandler := ctn.Get("v4a_error_handler").(v4a2.ErrorHandler)
			transmitter := ctn.Get("v4a_transmitter").(helpers2.Transmitter)

			return v4a2.NewBaseConnectHandler(
				v4aLogger,
				sender,
				errorHandler,
				transmitter,
			)
		},
	}

	bindHandlerDef := di.Def{
		Name:  "v4a_bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			v4aLogger := ctn.Get("v4a_logger").(v4a4.Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v4a_sender").(v4a.Sender)
			errorHandler := ctn.Get("v4a_error_handler").(v4a2.ErrorHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			transmitter := ctn.Get("v4a_transmitter").(helpers2.Transmitter)

			return v4a2.NewBaseBindHandler(
				v4aLogger,
				addressUtils,
				sender,
				errorHandler,
				bindManager,
				transmitter,
			)
		},
	}

	handlerDef := di.Def{
		Name:  "v4a_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("v4a_parser").(v4a.Parser)
			v4aLogger := ctn.Get("v4a_logger").(v4a4.Logger)
			connectHandler := ctn.Get("v4a_connect_handler").(v4a2.ConnectHandler)
			bindHandler := ctn.Get("v4a_bind_handler").(v4a2.BindHandler)
			sender := ctn.Get("v4a_sender").(v4a.Sender)
			errorHandler := ctn.Get("v4a_error_handler").(v4a2.ErrorHandler)
			validator := ctn.Get("v4a_validator").(helpers2.Validator)
			cleaner := ctn.Get("v4a_cleaner").(helpers2.Cleaner)
			replicator := ctn.Get("v4a_config_replicator").(v4a3.ConfigReplicator)

			return v4a2.NewBaseHandler(
				parser,
				v4aLogger,
				connectHandler,
				bindHandler,
				sender,
				errorHandler,
				validator,
				cleaner,
				replicator,
			)
		},
	}

	errorHandlerDef := di.Def{
		Name:  "v4a_error_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			sender := ctn.Get("v4a_sender").(v4a.Sender)
			v4Logger := ctn.Get("v4a_logger").(v4a4.Logger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return v4a2.NewBaseErrorHandler(
				v4Logger,
				sender,
				errorUtils,
			)
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

			return helpers2.NewBaseCleaner(manager)
		},
	}

	limiterDef := di.Def{
		Name:  "v4a_limiter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

			return helpers2.NewBaseLimiter(manager)
		},
	}

	transmitterDef := di.Def{
		Name:  "v4a_transmitter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
			bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
			bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

			return helpers2.NewBaseTransmitter(connect, bind, bindRate)
		},
	}

	validatorDef := di.Def{
		Name:  "v4a_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(managers.BlacklistManager)
			sender := ctn.Get("v4a_sender").(v4a.Sender)
			logger := ctn.Get("v4a_logger").(v4a4.Logger)
			limiter := ctn.Get("v4a_limiter").(helpers2.Limiter)

			return helpers2.NewBaseValidator(whitelist, blacklist, sender, logger, limiter)
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			passwordAuthenticator := ctn.Get("password_authenticator").(v52.Authenticator)
			noAuthAuthenticator := ctn.Get("no_auth_authenticator").(v52.Authenticator)
			sender := ctn.Get("v5_sender").(v5.Sender)

			return v52.NewBaseAuthenticationHandler(
				errorHandler,
				passwordAuthenticator,
				noAuthAuthenticator,
				sender,
			), nil
		},
	}

	connectHandlerDef := di.Def{
		Name:  "v5_connect_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			v5Logger := ctn.Get("v5_logger").(v54.Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v5_sender").(v5.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			transmitter := ctn.Get("v5_transmitter").(helpers.Transmitter)

			return v52.NewBaseConnectHandler(
				v5Logger,
				addressUtils,
				sender,
				errorHandler,
				transmitter,
			)
		},
	}

	bindHandlerDef := di.Def{
		Name:  "v5_bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			v5Logger := ctn.Get("v5_logger").(v54.Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v5_sender").(v5.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			transmitter := ctn.Get("v5_transmitter").(helpers.Transmitter)

			return v52.NewBaseBindHandler(
				addressUtils,
				v5Logger,
				sender,
				errorHandler,
				bindManager,
				transmitter,
			)
		},
	}

	udpAssociationHandler := di.Def{
		Name:  "v5_udp_association_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			v5Logger := ctn.Get("v5_logger").(v54.Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			udpClientManager := ctn.Get("udp_client_manager").(managers.UdpClientManager)
			sender := ctn.Get("v5_sender").(v5.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)

			return v52.NewBaseUdpAssociationHandler(
				addressUtils,
				udpClientManager,
				v5Logger,
				sender,
				errorHandler,
			)
		},
	}

	handlerDef := di.Def{
		Name:  "v5_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("v5_parser").(v5.Parser)
			authenticationHandler := ctn.Get("authentication_handler").(v52.AuthenticationHandler)
			v5Logger := ctn.Get("v5_logger").(v54.Logger)
			connectHandler := ctn.Get("v5_connect_handler").(v52.ConnectHandler)
			bindHandler := ctn.Get("v5_bind_handler").(v52.BindHandler)
			associationHandler := ctn.Get("v5_udp_association_handler").(v52.UdpAssociationHandler)
			sender := ctn.Get("v5_sender").(v5.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			validator := ctn.Get("v5_validator").(helpers.Validator)
			receiver := ctn.Get("v5_receiver").(v5.Receiver)
			cleaner := ctn.Get("v5_cleaner").(helpers.Cleaner)
			replicator := ctn.Get("v5_config_replicator").(v53.ConfigReplicator)

			return v52.NewBaseHandler(
				parser,
				authenticationHandler,
				v5Logger,
				connectHandler,
				bindHandler,
				associationHandler,
				errorHandler,
				sender,
				receiver,
				validator,
				cleaner,
				replicator,
			)
		},
	}

	errorHandlerDef := di.Def{
		Name:  "v5_error_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			sender := ctn.Get("v5_sender").(v5.Sender)
			v5Logger := ctn.Get("v5_logger").(v54.Logger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return v52.NewBaseErrorHandler(
				v5Logger,
				sender,
				errorUtils,
			)
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			receiver := ctn.Get("auth_password_receiver").(password.Receiver)
			sender := ctn.Get("auth_password_sender").(password.Sender)

			return authenticator.NewBasePasswordAuthenticator(
				errorHandler,
				sender,
				receiver,
			)
		},
	}

	noAuthAuthenticatorDef := di.Def{
		Name:  "no_auth_authenticator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return authenticator.NewBaseNoAuthAuthenticator()
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

			return helpers.NewBaseCleaner(manager)
		},
	}

	limiterDef := di.Def{
		Name:  "v5_limiter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

			return helpers.NewBaseLimiter(manager)
		},
	}

	transmitterDef := di.Def{
		Name:  "v5_transmitter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
			bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
			bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

			return helpers.NewBaseTransmitter(connect, bind, bindRate)
		},
	}

	validatorDef := di.Def{
		Name:  "v5_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(managers.BlacklistManager)
			sender := ctn.Get("v5_sender").(v5.Sender)
			logger := ctn.Get("v5_logger").(v54.Logger)
			limiter := ctn.Get("v5_limiter").(helpers.Limiter)

			return helpers.NewBaseValidator(whitelist, blacklist, sender, logger, limiter)
		},
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
