package main

import (
	"github.com/sarulabs/di"
	"socks/config/tcp"
	"socks/config/udp"
	v43 "socks/config/v4"
	v4a3 "socks/config/v4a"
	v53 "socks/config/v5"
	"socks/handlers"
	helpers3 "socks/handlers/helpers"
	v42 "socks/handlers/v4"
	helpers5 "socks/handlers/v4/helpers"
	v4a2 "socks/handlers/v4a"
	helpers2 "socks/handlers/v4a/helpers"
	v52 "socks/handlers/v5"
	"socks/handlers/v5/authenticator"
	helpers4 "socks/handlers/v5/authenticator/helpers"
	"socks/handlers/v5/helpers"
	tcp2 "socks/logger/tcp"
	v44 "socks/logger/v4"
	v4a4 "socks/logger/v4a"
	v54 "socks/logger/v5"
	"socks/managers"
	"socks/protocol/auth/password"
	v4 "socks/protocol/v4"
	"socks/protocol/v4a"
	v5 "socks/protocol/v5"
	"socks/transfer"
	"socks/utils"
)

func registerHandlers(builder di.Builder) {
	connectionHandlerDef := di.Def{
		Name:  "connection_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			v4Handler := ctn.Get("v4_handler").(v42.Handler)
			v4aHandler := ctn.Get("v4a_handler").(v4a2.Handler)
			v5Handler := ctn.Get("v5_handler").(v52.Handler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			tcpLogger := ctn.Get("tcp_logger").(tcp2.Logger)
			receiver := ctn.Get("receiver").(helpers3.Receiver)

			return handlers.NewBaseConnectionHandler(
				streamHandler,
				v4Handler,
				v4aHandler,
				v5Handler,
				bindManager,
				addressUtils,
				tcpLogger,
				receiver,
			)
		},
	}

	packetHandlerDef := di.Def{
		Name:  "packet_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("v5_parser").(v5.Parser)
			builder := ctn.Get("v5_builder").(v5.Builder)
			udpAssociationManager := ctn.Get("udp_association_manager").(managers.UdpAssociationManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)

			return handlers.NewBasePacketHandler(
				parser,
				builder,
				udpAssociationManager,
				addressUtils,
			), nil
		},
	}

	err := builder.Add(
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
	receiverDef := di.Def{
		Name:  "receiver",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("tcp_deadline_config").(tcp.DeadlineConfig)
			deadlineManager := ctn.Get("deadline_manager").(managers.DeadlineManager)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)

			return helpers3.NewBaseReceiver(cfg, deadlineManager, bindManager)
		},
	}

	err := builder.Add(
		receiverDef,
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
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			sender := ctn.Get("v4_sender").(helpers5.Sender)
			errorHandler := ctn.Get("v4_error_handler").(v42.ErrorHandler)
			dialer := ctn.Get("v4_dialer").(helpers5.Dialer)

			return v42.NewBaseConnectHandler(
				streamHandler,
				v4Logger,
				sender,
				errorHandler,
				dialer,
			)
		},
	}

	bindHandlerDef := di.Def{
		Name:  "v4_bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			v4Logger := ctn.Get("v4_logger").(v44.Logger)
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v4_sender").(helpers5.Sender)
			errorHandler := ctn.Get("v4_error_handler").(v42.ErrorHandler)
			receiver := ctn.Get("v4_receiver").(helpers5.Receiver)

			return v42.NewBaseBindHandler(
				v4Logger,
				streamHandler,
				bindManager,
				addressUtils,
				sender,
				errorHandler,
				receiver,
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
			sender := ctn.Get("v4_sender").(helpers5.Sender)
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
			sender := ctn.Get("v4_sender").(helpers5.Sender)
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
	blacklistDef := di.Def{
		Name:  "v4_blacklist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_config").(v43.Config)
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

	receiverDef := di.Def{
		Name:  "v4_receiver",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_deadline_config").(v43.DeadlineConfig)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)

			return helpers5.NewBaseReceiver(cfg, bindManager)
		},
	}

	senderDef := di.Def{
		Name:  "v4_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			protocol := ctn.Get("v4").(v4.Protocol)
			tcpConfig := ctn.Get("tcp_config").(tcp.Config)

			return helpers5.NewBaseSender(
				protocol,
				tcpConfig,
			)
		},
	}

	validatorDef := di.Def{
		Name:  "v4_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_config").(v43.Config)
			whitelist := ctn.Get("v4_whitelist").(helpers5.Whitelist)
			blacklist := ctn.Get("v4_blacklist").(helpers5.Blacklist)
			sender := ctn.Get("v4_sender").(helpers5.Sender)
			logger := ctn.Get("v4_logger").(v44.Logger)

			return helpers5.NewBaseValidator(cfg, whitelist, blacklist, sender, logger)
		},
	}

	whitelistDef := di.Def{
		Name:  "v4_whitelist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_config").(v43.Config)
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)

			return helpers5.NewBaseWhitelist(cfg, whitelist)
		},
	}

	err := builder.Add(
		blacklistDef,
		dialerDef,
		receiverDef,
		senderDef,
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
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			sender := ctn.Get("v4a_sender").(helpers2.Sender)
			errorHandler := ctn.Get("v4a_error_handler").(v4a2.ErrorHandler)
			dialer := ctn.Get("v4a_dialer").(helpers2.Dialer)

			return v4a2.NewBaseConnectHandler(
				streamHandler,
				v4aLogger,
				sender,
				errorHandler,
				dialer,
			)
		},
	}

	bindHandlerDef := di.Def{
		Name:  "v4a_bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			v4aLogger := ctn.Get("v4a_logger").(v4a4.Logger)
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v4a_sender").(helpers2.Sender)
			errorHandler := ctn.Get("v4a_error_handler").(v4a2.ErrorHandler)
			receiver := ctn.Get("v4a_receiver").(helpers2.Receiver)

			return v4a2.NewBaseBindHandler(
				v4aLogger,
				streamHandler,
				bindManager,
				addressUtils,
				sender,
				errorHandler,
				receiver,
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
			sender := ctn.Get("v4a_sender").(helpers2.Sender)
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
			sender := ctn.Get("v4a_sender").(helpers2.Sender)
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
	blacklistDef := di.Def{
		Name:  "v4a_blacklist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_config").(v4a3.Config)
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

	receiverDef := di.Def{
		Name:  "v4a_receiver",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_deadline_config").(v4a3.DeadlineConfig)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)

			return helpers2.NewBaseReceiver(cfg, bindManager)
		},
	}

	senderDef := di.Def{
		Name:  "v4a_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			protocol := ctn.Get("v4a").(v4a.Protocol)
			tcpConfig := ctn.Get("tcp_config").(tcp.Config)

			return helpers2.NewBaseSender(
				protocol,
				tcpConfig,
			)
		},
	}

	validatorDef := di.Def{
		Name:  "v4a_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_config").(v4a3.Config)
			whitelist := ctn.Get("v4a_whitelist").(helpers2.Whitelist)
			blacklist := ctn.Get("v4a_blacklist").(helpers2.Blacklist)
			sender := ctn.Get("v4a_sender").(helpers2.Sender)
			logger := ctn.Get("v4a_logger").(v4a4.Logger)

			return helpers2.NewBaseValidator(cfg, whitelist, blacklist, sender, logger)
		},
	}

	whitelistDef := di.Def{
		Name:  "v4a_whitelist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_config").(v4a3.Config)
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)

			return helpers2.NewBaseWhitelist(cfg, whitelist)
		},
	}

	err := builder.Add(
		blacklistDef,
		dialerDef,
		receiverDef,
		senderDef,
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
			protocol := ctn.Get("v5").(v5.Protocol)
			cfg := ctn.Get("v5_config").(v53.Config)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			passwordAuthenticator := ctn.Get("password_authenticator").(v52.Authenticator)
			noAuthAuthenticator := ctn.Get("no_auth_authenticator").(v52.Authenticator)

			return v52.NewBaseAuthenticationHandler(
				protocol,
				cfg,
				errorHandler,
				passwordAuthenticator,
				noAuthAuthenticator,
			), nil
		},
	}

	connectHandlerDef := di.Def{
		Name:  "v5_connect_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			v5Logger := ctn.Get("v5_logger").(v54.Logger)
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v5_sender").(helpers.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			whitelist := ctn.Get("v5_whitelist").(helpers.Whitelist)
			blacklist := ctn.Get("v5_blacklist").(helpers.Blacklist)
			dialer := ctn.Get("v5_dialer").(helpers.Dialer)

			return v52.NewBaseConnectHandler(
				streamHandler,
				v5Logger,
				addressUtils,
				sender,
				errorHandler,
				whitelist,
				blacklist,
				dialer,
			)
		},
	}

	bindHandlerDef := di.Def{
		Name:  "v5_bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			v5Logger := ctn.Get("v5_logger").(v54.Logger)
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v5_sender").(helpers.Sender)
			whitelist := ctn.Get("v5_whitelist").(helpers.Whitelist)
			blacklist := ctn.Get("v5_blacklist").(helpers.Blacklist)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			receiver := ctn.Get("v5_receiver").(helpers.Receiver)

			return v52.NewBaseBindHandler(
				bindManager,
				streamHandler,
				addressUtils,
				v5Logger,
				sender,
				whitelist,
				blacklist,
				errorHandler,
				receiver,
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
			udpAssociationManager := ctn.Get("udp_association_manager").(managers.UdpAssociationManager)
			sender := ctn.Get("v5_sender").(helpers.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)

			return v52.NewBaseUdpAssociationHandler(
				cfg,
				addressUtils,
				udpAssociationManager,
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
			sender := ctn.Get("v5_sender").(helpers.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			validator := ctn.Get("v5_validator").(helpers.Validator)
			receiver := ctn.Get("v5_receiver").(helpers.Receiver)

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
			sender := ctn.Get("v5_sender").(helpers.Sender)
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

func registerAuthenticatorHelpers(builder di.Builder) {
	receiverDef := di.Def{
		Name:  "v5_auth_receiver",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_deadline_config").(v53.DeadlineConfig)
			deadlineManager := ctn.Get("deadline_manager").(managers.DeadlineManager)
			parser := ctn.Get("auth_password_parser").(password.Parser)

			return helpers4.NewBaseReceiver(cfg, deadlineManager, parser)
		},
	}

	err := builder.Add(
		receiverDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerAuthenticators(builder di.Builder) {
	passwordAuthenticatorDef := di.Def{
		Name:  "password_authenticator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			passwd := ctn.Get("auth_password").(password.Password)
			cfg := ctn.Get("v5_config").(v53.Config)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			receiver := ctn.Get("v5_auth_receiver").(helpers4.Receiver)

			return authenticator.NewBasePasswordAuthenticator(
				passwd,
				cfg,
				errorHandler,
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

	registerAuthenticatorHelpers(builder)
}

func registerV5Helpers(builder di.Builder) {
	blacklistDef := di.Def{
		Name:  "v5_blacklist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_config").(v53.Config)
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

	receiverDef := di.Def{
		Name:  "v5_receiver",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_deadline_config").(v53.DeadlineConfig)
			deadlineManager := ctn.Get("deadline_manager").(managers.DeadlineManager)
			parser := ctn.Get("v5_parser").(v5.Parser)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)

			return helpers.NewBaseReceiver(cfg, deadlineManager, parser, bindManager)
		},
	}

	senderDef := di.Def{
		Name:  "v5_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			protocol := ctn.Get("v5").(v5.Protocol)
			tcpConfig := ctn.Get("tcp_config").(tcp.Config)
			udpConfig := ctn.Get("udp_config").(udp.Config)

			return helpers.NewBaseSender(
				protocol,
				tcpConfig,
				udpConfig,
			)
		},
	}

	validatorDef := di.Def{
		Name:  "v5_validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_config").(v53.Config)
			whitelist := ctn.Get("v5_whitelist").(helpers.Whitelist)
			blacklist := ctn.Get("v5_blacklist").(helpers.Blacklist)
			sender := ctn.Get("v5_sender").(helpers.Sender)
			logger := ctn.Get("v5_logger").(v54.Logger)

			return helpers.NewBaseValidator(cfg, whitelist, blacklist, sender, logger)
		},
	}

	whitelistDef := di.Def{
		Name:  "v5_whitelist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_config").(v53.Config)
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)

			return helpers.NewBaseWhitelist(cfg, whitelist)
		},
	}

	err := builder.Add(
		blacklistDef,
		dialerDef,
		receiverDef,
		senderDef,
		validatorDef,
		whitelistDef,
	)

	if err != nil {
		panic(err)
	}
}
