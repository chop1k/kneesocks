package dependency

import (
	"socks/internal/kneesocks/config/tcp"
	udpConfig "socks/internal/kneesocks/config/udp"
	v4Config "socks/internal/kneesocks/config/v4"
	v4aConfig "socks/internal/kneesocks/config/v4a"
	v5Config "socks/internal/kneesocks/config/v5"
	"socks/internal/kneesocks/handlers"
	v4Handlers "socks/internal/kneesocks/handlers/v4"
	"socks/internal/kneesocks/handlers/v4/helpers"
	v4aHandler "socks/internal/kneesocks/handlers/v4a"
	v4aHelpers "socks/internal/kneesocks/handlers/v4a/helpers"
	v5Handlers "socks/internal/kneesocks/handlers/v5"
	"socks/internal/kneesocks/handlers/v5/authenticator"
	v5Helpers "socks/internal/kneesocks/handlers/v5/helpers"
	tcpLogger "socks/internal/kneesocks/logger/tcp"
	"socks/internal/kneesocks/logger/udp"
	v4Logger "socks/internal/kneesocks/logger/v4"
	v4aLogger "socks/internal/kneesocks/logger/v4a"
	v5Logger "socks/internal/kneesocks/logger/v5"
	"socks/internal/kneesocks/managers"
	"socks/internal/kneesocks/transfer"
	protocolHelpers "socks/pkg/protocol"
	"socks/pkg/protocol/auth/password"
	v4Protocol "socks/pkg/protocol/v4"
	v4aProtocol "socks/pkg/protocol/v4a"
	v5Protocol "socks/pkg/protocol/v5"
	"socks/pkg/utils"

	"github.com/sarulabs/di"
)

func registerHandlers(builder di.Builder) {
	bindHandlerDef := di.Def{
		Name:  "bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			logger := ctn.Get("tcp_logger").(tcpLogger.Logger)
			bindHandler := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)

			return handlers.NewBindHandler(addressUtils, logger, bindHandler, bindManager)
		},
	}

	connectionHandlerDef := di.Def{
		Name:  "connection_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			_v4 := ctn.Get("v4_handler")

			var v4 *v4Handlers.Handler

			if _v4 == nil {
				v4 = nil
			} else {
				v4 = _v4.(*v4Handlers.Handler)
			}

			_v4a := ctn.Get("v4a_handler")

			var v4a *v4aHandler.Handler

			if _v4 == nil {
				v4a = nil
			} else {
				v4a = _v4a.(*v4aHandler.Handler)
			}

			_v5 := ctn.Get("v5_handler")

			var v5 *v5Handlers.Handler

			if _v5 == nil {
				v5 = nil
			} else {
				v5 = _v5.(*v5Handlers.Handler)
			}

			logger := ctn.Get("tcp_logger").(tcpLogger.Logger)
			receiver := ctn.Get("receiver").(protocolHelpers.Receiver)
			bindHandler := ctn.Get("bind_handler").(handlers.BindHandler)
			replicator := ctn.Get("tcp_config_replicator").(tcp.ConfigReplicator)

			return handlers.NewConnectionHandler(
				v4,
				v4a,
				v5,
				logger,
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
			parser := ctn.Get("v5_parser").(v5Protocol.Parser)
			builder := ctn.Get("v5_builder").(v5Protocol.Builder)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			clients := ctn.Get("udp_client_manager").(managers.UdpClientManager)
			hosts := ctn.Get("udp_host_manager").(managers.UdpHostManager)
			logger := ctn.Get("udp_logger").(udp.Logger)

			_replicator := ctn.Get("udp_config_replicator")

			if _replicator == nil {
				return nil, nil
			}

			replicator := _replicator.(udpConfig.ConfigReplicator)

			return handlers.NewPacketHandler(
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
			logger := ctn.Get("v4_logger").(v4Logger.Logger)
			sender := ctn.Get("v4_sender").(v4Protocol.Sender)
			errorHandler := ctn.Get("v4_error_handler").(v4Handlers.ErrorHandler)
			transmitter := ctn.Get("v4_transmitter").(helpers.Transmitter)

			return v4Handlers.NewConnectHandler(
				logger,
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
			logger := ctn.Get("v4_logger").(v4Logger.Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v4_sender").(v4Protocol.Sender)
			errorHandler := ctn.Get("v4_error_handler").(v4Handlers.ErrorHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			transmitter := ctn.Get("v4_transmitter").(helpers.Transmitter)

			return v4Handlers.NewBindHandler(
				logger,
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
			parser := ctn.Get("v4_parser").(v4Protocol.Parser)
			logger := ctn.Get("v4_logger").(v4Logger.Logger)
			connectHandler := ctn.Get("v4_connect_handler").(v4Handlers.ConnectHandler)
			bindHandler := ctn.Get("v4_bind_handler").(v4Handlers.BindHandler)
			sender := ctn.Get("v4_sender").(v4Protocol.Sender)
			errorHandler := ctn.Get("v4_error_handler").(v4Handlers.ErrorHandler)
			validator := ctn.Get("v4_validator").(helpers.Validator)
			cleaner := ctn.Get("v4_cleaner").(helpers.Cleaner)
			_replicator := ctn.Get("v4_config_replicator")

			if _replicator == nil {
				return nil, nil
			}

			replicator := _replicator.(v4Config.ConfigReplicator)

			return v4Handlers.NewHandler(
				parser,
				logger,
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
			sender := ctn.Get("v4_sender").(v4Protocol.Sender)
			logger := ctn.Get("v4_logger").(v4Logger.Logger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return v4Handlers.NewErrorHandler(
				logger,
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

			return helpers.NewCleaner(manager)
		},
	}

	limiterDef := di.Def{
		Name:  "v4_limiter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

			return helpers.NewLimiter(manager)
		},
	}

	transmitterDef := di.Def{
		Name:  "v4_transmitter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
			bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
			bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

			return helpers.NewTransmitter(connect, bind, bindRate)
		},
	}

	validatorDef := di.Def{
		Name:  "v4_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(managers.BlacklistManager)
			sender := ctn.Get("v4_sender").(v4Protocol.Sender)
			logger := ctn.Get("v4_logger").(v4Logger.Logger)
			limiter := ctn.Get("v4_limiter").(helpers.Limiter)

			return helpers.NewValidator(whitelist, blacklist, sender, logger, limiter)
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
			logger := ctn.Get("v4a_logger").(v4aLogger.Logger)
			sender := ctn.Get("v4a_sender").(v4aProtocol.Sender)
			errorHandler := ctn.Get("v4a_error_handler").(v4aHandler.ErrorHandler)
			transmitter := ctn.Get("v4a_transmitter").(v4aHelpers.Transmitter)

			return v4aHandler.NewConnectHandler(
				logger,
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
			logger := ctn.Get("v4a_logger").(v4aLogger.Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v4a_sender").(v4aProtocol.Sender)
			errorHandler := ctn.Get("v4a_error_handler").(v4aHandler.ErrorHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			transmitter := ctn.Get("v4a_transmitter").(v4aHelpers.Transmitter)

			return v4aHandler.NewBindHandler(
				logger,
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
			parser := ctn.Get("v4a_parser").(v4aProtocol.Parser)
			logger := ctn.Get("v4a_logger").(v4aLogger.Logger)
			connectHandler := ctn.Get("v4a_connect_handler").(v4aHandler.ConnectHandler)
			bindHandler := ctn.Get("v4a_bind_handler").(v4aHandler.BindHandler)
			sender := ctn.Get("v4a_sender").(v4aProtocol.Sender)
			errorHandler := ctn.Get("v4a_error_handler").(v4aHandler.ErrorHandler)
			validator := ctn.Get("v4a_validator").(v4aHelpers.Validator)
			cleaner := ctn.Get("v4a_cleaner").(v4aHelpers.Cleaner)
			_replicator := ctn.Get("v4a_config_replicator")

			if _replicator == nil {
				return nil, nil
			}

			replicator := _replicator.(v4aConfig.ConfigReplicator)

			return v4aHandler.NewHandler(
				parser,
				logger,
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
			sender := ctn.Get("v4a_sender").(v4aProtocol.Sender)
			logger := ctn.Get("v4a_logger").(v4aLogger.Logger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return v4aHandler.NewErrorHandler(
				logger,
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

			return v4aHelpers.NewCleaner(manager)
		},
	}

	limiterDef := di.Def{
		Name:  "v4a_limiter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

			return v4aHelpers.NewLimiter(manager)
		},
	}

	transmitterDef := di.Def{
		Name:  "v4a_transmitter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
			bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
			bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

			return v4aHelpers.NewTransmitter(connect, bind, bindRate)
		},
	}

	validatorDef := di.Def{
		Name:  "v4a_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(managers.BlacklistManager)
			sender := ctn.Get("v4a_sender").(v4aProtocol.Sender)
			logger := ctn.Get("v4a_logger").(v4aLogger.Logger)
			limiter := ctn.Get("v4a_limiter").(v4aHelpers.Limiter)

			return v4aHelpers.NewValidator(whitelist, blacklist, sender, logger, limiter)
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
			errorHandler := ctn.Get("v5_error_handler").(v5Handlers.ErrorHandler)
			passwordAuthenticator := ctn.Get("password_authenticator").(v5Handlers.Authenticator)
			noAuthAuthenticator := ctn.Get("no_auth_authenticator").(v5Handlers.Authenticator)
			sender := ctn.Get("v5_sender").(v5Protocol.Sender)

			return v5Handlers.NewAuthenticationHandler(
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
			logger := ctn.Get("v5_logger").(v5Logger.Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v5_sender").(v5Protocol.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v5Handlers.ErrorHandler)
			transmitter := ctn.Get("v5_transmitter").(v5Helpers.Transmitter)

			return v5Handlers.NewConnectHandler(
				logger,
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
			logger := ctn.Get("v5_logger").(v5Logger.Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v5_sender").(v5Protocol.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v5Handlers.ErrorHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			transmitter := ctn.Get("v5_transmitter").(v5Helpers.Transmitter)

			return v5Handlers.NewBindHandler(
				addressUtils,
				logger,
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
			logger := ctn.Get("v5_logger").(v5Logger.Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			udpManager := ctn.Get("udp_client_manager").(managers.UdpClientManager)
			sender := ctn.Get("v5_sender").(v5Protocol.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v5Handlers.ErrorHandler)

			return v5Handlers.NewUdpAssociationHandler(
				addressUtils,
				udpManager,
				logger,
				sender,
				errorHandler,
			)
		},
	}

	handlerDef := di.Def{
		Name:  "v5_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("v5_parser").(v5Protocol.Parser)
			authenticationHandler := ctn.Get("authentication_handler").(v5Handlers.AuthenticationHandler)
			logger := ctn.Get("v5_logger").(v5Logger.Logger)
			connectHandler := ctn.Get("v5_connect_handler").(v5Handlers.ConnectHandler)
			bindHandler := ctn.Get("v5_bind_handler").(v5Handlers.BindHandler)
			associationHandler := ctn.Get("v5_udp_association_handler").(v5Handlers.UdpAssociationHandler)
			sender := ctn.Get("v5_sender").(v5Protocol.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v5Handlers.ErrorHandler)
			validator := ctn.Get("v5_validator").(v5Helpers.Validator)
			receiver := ctn.Get("v5_receiver").(v5Protocol.Receiver)
			cleaner := ctn.Get("v5_cleaner").(v5Helpers.Cleaner)
			_replicator := ctn.Get("v5_config_replicator")

			if _replicator == nil {
				return nil, nil
			}

			replicator := _replicator.(v5Config.ConfigReplicator)

			return v5Handlers.NewHandler(
				parser,
				authenticationHandler,
				logger,
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
			sender := ctn.Get("v5_sender").(v5Protocol.Sender)
			logger := ctn.Get("v5_logger").(v5Logger.Logger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return v5Handlers.NewErrorHandler(
				logger,
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
			errorHandler := ctn.Get("v5_error_handler").(v5Handlers.ErrorHandler)
			receiver := ctn.Get("auth_password_receiver").(password.Receiver)
			sender := ctn.Get("auth_password_sender").(password.Sender)

			return authenticator.NewPasswordAuthenticator(
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
			return authenticator.NewNoAuthAuthenticator()
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

			return v5Helpers.NewCleaner(manager)
		},
	}

	limiterDef := di.Def{
		Name:  "v5_limiter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

			return v5Helpers.NewLimiter(manager)
		},
	}

	transmitterDef := di.Def{
		Name:  "v5_transmitter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
			bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
			bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

			return v5Helpers.NewTransmitter(connect, bind, bindRate)
		},
	}

	validatorDef := di.Def{
		Name:  "v5_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(managers.BlacklistManager)
			sender := ctn.Get("v5_sender").(v5Protocol.Sender)
			logger := ctn.Get("v5_logger").(v5Logger.Logger)
			limiter := ctn.Get("v5_limiter").(v5Helpers.Limiter)

			return v5Helpers.NewValidator(whitelist, blacklist, sender, logger, limiter)
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
