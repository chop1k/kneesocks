package main

import (
	"github.com/sarulabs/di"
	"socks/config/tcp"
	udp2 "socks/config/udp"
	v43 "socks/config/v4"
	v4a3 "socks/config/v4a"
	v53 "socks/config/v5"
	"socks/handlers"
	helpers4 "socks/handlers/helpers"
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
			binder := ctn.Get("binder").(helpers4.Binder)
			bindHandler := ctn.Get("transfer_bind_handler").(transfer.BindHandler)

			return handlers.NewBaseBindHandler(addressUtils, binder, tcpLogger, bindHandler)
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

			return handlers.NewBaseConnectionHandler(
				v4Handler,
				v4aHandler,
				v5Handler,
				tcpLogger,
				receiver,
				bindHandler,
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
			cfg := ctn.Get("udp_config").(udp2.BindConfig)
			buffer := ctn.Get("udp_buffer_config").(udp2.BufferConfig)

			return handlers.NewBasePacketHandler(
				parser,
				builder,
				addressUtils,
				clients,
				hosts,
				logger,
				cfg,
				buffer,
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

	registerHelpers(builder)
	registerV4Handlers(builder)
	registerV4aHandlers(builder)
	registerV5Handlers(builder)
}

func registerHelpers(builder di.Builder) {
	binderDef := di.Def{
		Name:  "binder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("tcp_deadline_config").(tcp.DeadlineConfig)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)

			return helpers4.NewBaseBinder(cfg, bindManager)
		},
	}

	err := builder.Add(
		binderDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV4Handlers(builder di.Builder) {
	connectHandlerDef := di.Def{
		Name:  "v4_connect_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			v4Logger := ctn.Get("v4_logger").(v44.Logger)
			sender := ctn.Get("v4_sender").(v4.Sender)
			errorHandler := ctn.Get("v4_error_handler").(v42.ErrorHandler)
			dialer := ctn.Get("v4_dialer").(helpers5.Dialer)
			transmitter := ctn.Get("v4_transmitter").(helpers5.Transmitter)

			return v42.NewBaseConnectHandler(
				v4Logger,
				sender,
				errorHandler,
				dialer,
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
			binder := ctn.Get("v4_binder").(helpers5.Binder)
			transmitter := ctn.Get("v4_transmitter").(helpers5.Transmitter)

			return v42.NewBaseBindHandler(
				v4Logger,
				addressUtils,
				sender,
				errorHandler,
				binder,
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

			return v42.NewBaseHandler(
				parser,
				v4Logger,
				connectHandler,
				bindHandler,
				sender,
				errorHandler,
				validator,
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
	binderDef := di.Def{
		Name:  "v4_binder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_deadline_config").(v4a3.DeadlineConfig)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)

			return helpers5.NewBaseBinder(cfg, bindManager)
		},
	}

	blacklistDef := di.Def{
		Name:  "v4_blacklist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_restrictions_config").(v43.RestrictionsConfig)
			whitelist := ctn.Get("blacklist_manager").(managers.BlacklistManager)

			return helpers5.NewBaseBlacklist(cfg, whitelist)
		},
	}

	dialerDef := di.Def{
		Name:  "v4_dialer",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_deadline_config").(v43.DeadlineConfig)

			return helpers5.NewBaseDialer(cfg)
		},
	}

	transmitterDef := di.Def{
		Name:  "v4_transmitter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_restrictions_config").(v43.RestrictionsConfig)
			connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
			bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
			bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

			return helpers5.NewBaseTransmitter(cfg, connect, bind, bindRate)
		},
	}

	validatorDef := di.Def{
		Name:  "v4_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_config").(v43.Config)
			whitelist := ctn.Get("v4_whitelist").(helpers5.Whitelist)
			blacklist := ctn.Get("v4_blacklist").(helpers5.Blacklist)
			sender := ctn.Get("v4_sender").(v4.Sender)
			logger := ctn.Get("v4_logger").(v44.Logger)

			return helpers5.NewBaseValidator(cfg, whitelist, blacklist, sender, logger)
		},
	}

	whitelistDef := di.Def{
		Name:  "v4_whitelist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_restrictions_config").(v43.RestrictionsConfig)
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)

			return helpers5.NewBaseWhitelist(cfg, whitelist)
		},
	}

	err := builder.Add(
		binderDef,
		blacklistDef,
		dialerDef,
		transmitterDef,
		validatorDef,
		whitelistDef,
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
			dialer := ctn.Get("v4a_dialer").(helpers2.Dialer)
			transmitter := ctn.Get("v4a_transmitter").(helpers2.Transmitter)

			return v4a2.NewBaseConnectHandler(
				v4aLogger,
				sender,
				errorHandler,
				dialer,
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
			binder := ctn.Get("v4a_binder").(helpers2.Binder)
			transmitter := ctn.Get("v4a_transmitter").(helpers2.Transmitter)

			return v4a2.NewBaseBindHandler(
				v4aLogger,
				addressUtils,
				sender,
				errorHandler,
				binder,
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

			return v4a2.NewBaseHandler(
				parser,
				v4aLogger,
				connectHandler,
				bindHandler,
				sender,
				errorHandler,
				validator,
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
	binderDef := di.Def{
		Name:  "v4a_binder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_deadline_config").(v4a3.DeadlineConfig)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)

			return helpers2.NewBaseBinder(cfg, bindManager)
		},
	}

	blacklistDef := di.Def{
		Name:  "v4a_blacklist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_restrictions_config").(v4a3.RestrictionsConfig)
			whitelist := ctn.Get("blacklist_manager").(managers.BlacklistManager)

			return helpers2.NewBaseBlacklist(cfg, whitelist)
		},
	}

	dialerDef := di.Def{
		Name:  "v4a_dialer",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_deadline_config").(v4a3.DeadlineConfig)

			return helpers2.NewBaseDialer(cfg)
		},
	}

	transmitterDef := di.Def{
		Name:  "v4a_transmitter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_restrictions_config").(v4a3.RestrictionsConfig)
			connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
			bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
			bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

			return helpers2.NewBaseTransmitter(cfg, connect, bind, bindRate)
		},
	}

	validatorDef := di.Def{
		Name:  "v4a_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_config").(v4a3.Config)
			whitelist := ctn.Get("v4a_whitelist").(helpers2.Whitelist)
			blacklist := ctn.Get("v4a_blacklist").(helpers2.Blacklist)
			sender := ctn.Get("v4a_sender").(v4a.Sender)
			logger := ctn.Get("v4a_logger").(v4a4.Logger)

			return helpers2.NewBaseValidator(cfg, whitelist, blacklist, sender, logger)
		},
	}

	whitelistDef := di.Def{
		Name:  "v4a_whitelist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_restrictions_config").(v4a3.RestrictionsConfig)
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)

			return helpers2.NewBaseWhitelist(cfg, whitelist)
		},
	}

	err := builder.Add(
		binderDef,
		blacklistDef,
		dialerDef,
		transmitterDef,
		validatorDef,
		whitelistDef,
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
			cfg := ctn.Get("v5_config").(v53.Config)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			passwordAuthenticator := ctn.Get("password_authenticator").(v52.Authenticator)
			noAuthAuthenticator := ctn.Get("no_auth_authenticator").(v52.Authenticator)
			sender := ctn.Get("v5_sender").(v5.Sender)

			return v52.NewBaseAuthenticationHandler(
				cfg,
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
			dialer := ctn.Get("v5_dialer").(helpers.Dialer)
			transmitter := ctn.Get("v5_transmitter").(helpers.Transmitter)

			return v52.NewBaseConnectHandler(
				v5Logger,
				addressUtils,
				sender,
				errorHandler,
				dialer,
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
			binder := ctn.Get("v5_binder").(helpers.Binder)
			transmitter := ctn.Get("v5_transmitter").(helpers.Transmitter)

			return v52.NewBaseBindHandler(
				addressUtils,
				v5Logger,
				sender,
				errorHandler,
				binder,
				transmitter,
			)
		},
	}

	udpAssociationHandler := di.Def{
		Name:  "v5_udp_association_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_config").(v53.Config)
			v5Logger := ctn.Get("v5_logger").(v54.Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			udpClientManager := ctn.Get("udp_client_manager").(managers.UdpClientManager)
			sender := ctn.Get("v5_sender").(v5.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)

			return v52.NewBaseUdpAssociationHandler(
				cfg,
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
			cfg := ctn.Get("users_config").(v53.UsersConfig)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			receiver := ctn.Get("auth_password_receiver").(password.Receiver)
			sender := ctn.Get("auth_password_sender").(password.Sender)

			return authenticator.NewBasePasswordAuthenticator(
				cfg,
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
	binderDef := di.Def{
		Name:  "v5_binder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_deadline_config").(v53.DeadlineConfig)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)

			return helpers.NewBaseBinder(cfg, bindManager)
		},
	}

	blacklistDef := di.Def{
		Name:  "v5_blacklist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("users_config").(v53.UsersConfig)
			whitelist := ctn.Get("blacklist_manager").(managers.BlacklistManager)

			return helpers.NewBaseBlacklist(cfg, whitelist)
		},
	}

	dialerDef := di.Def{
		Name:  "v5_dialer",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_deadline_config").(v53.DeadlineConfig)

			return helpers.NewBaseDialer(cfg)
		},
	}

	transmitterDef := di.Def{
		Name:  "v5_transmitter",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("users_config").(v53.UsersConfig)
			connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
			bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
			bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

			return helpers.NewBaseTransmitter(cfg, connect, bind, bindRate)
		},
	}

	validatorDef := di.Def{
		Name:  "v5_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_config").(v53.Config)
			whitelist := ctn.Get("v5_whitelist").(helpers.Whitelist)
			blacklist := ctn.Get("v5_blacklist").(helpers.Blacklist)
			sender := ctn.Get("v5_sender").(v5.Sender)
			logger := ctn.Get("v5_logger").(v54.Logger)

			return helpers.NewBaseValidator(cfg, whitelist, blacklist, sender, logger)
		},
	}

	whitelistDef := di.Def{
		Name:  "v5_whitelist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("users_config").(v53.UsersConfig)
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)

			return helpers.NewBaseWhitelist(cfg, whitelist)
		},
	}

	err := builder.Add(
		binderDef,
		blacklistDef,
		dialerDef,
		transmitterDef,
		validatorDef,
		whitelistDef,
	)

	if err != nil {
		panic(err)
	}
}
