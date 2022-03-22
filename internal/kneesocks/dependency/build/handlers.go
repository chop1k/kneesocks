package build

import (
	tcpConfig "socks/internal/kneesocks/config/tcp"
	udpConfig "socks/internal/kneesocks/config/udp"
	v4Config "socks/internal/kneesocks/config/v4"
	v4aConfig "socks/internal/kneesocks/config/v4a"
	v5Config "socks/internal/kneesocks/config/v5"
	"socks/internal/kneesocks/handlers"
	v4Handlers "socks/internal/kneesocks/handlers/v4"
	"socks/internal/kneesocks/handlers/v4/helpers"
	v4aHandlers "socks/internal/kneesocks/handlers/v4a"
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

func BindHandler(ctn di.Container) (interface{}, error) {
	addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
	logger := ctn.Get("tcp_logger").(tcpLogger.Logger)
	bindHandler := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
	bindManager := ctn.Get("bind_manager").(managers.BindManager)

	return handlers.NewBindHandler(addressUtils, logger, bindHandler, bindManager)
}

func ConnectionHandler(ctn di.Container) (interface{}, error) {
	_v4 := ctn.Get("v4_handler")

	var v4 *v4Handlers.Handler

	if _v4 == nil {
		v4 = nil
	} else {
		v4 = _v4.(*v4Handlers.Handler)
	}

	_v4a := ctn.Get("v4a_handler")

	var v4a *v4aHandlers.Handler

	if _v4a == nil {
		v4a = nil
	} else {
		v4a = _v4a.(*v4aHandlers.Handler)
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
	replicator := ctn.Get("tcp_config_replicator").(tcpConfig.ConfigReplicator)

	return handlers.NewConnectionHandler(
		v4,
		v4a,
		v5,
		logger,
		receiver,
		bindHandler,
		replicator,
	)
}

func PacketHandler(ctn di.Container) (interface{}, error) {
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
}

func V4ConnectHandler(ctn di.Container) (interface{}, error) {
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
}

func V4BindHandler(ctn di.Container) (interface{}, error) {
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
}

func V4Handler(ctn di.Container) (interface{}, error) {
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
}

func V4ErrorHandler(ctn di.Container) (interface{}, error) {
	sender := ctn.Get("v4_sender").(v4Protocol.Sender)
	logger := ctn.Get("v4_logger").(v4Logger.Logger)
	errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

	return v4Handlers.NewErrorHandler(
		logger,
		sender,
		errorUtils,
	)
}

func V4Cleaner(ctn di.Container) (interface{}, error) {
	manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

	return helpers.NewCleaner(manager)
}

func V4Limiter(ctn di.Container) (interface{}, error) {
	manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

	return helpers.NewLimiter(manager)
}

func V4Transmitter(ctn di.Container) (interface{}, error) {
	connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
	bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
	bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

	return helpers.NewTransmitter(connect, bind, bindRate)
}

func V4Validator(ctn di.Container) (interface{}, error) {
	whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)
	blacklist := ctn.Get("blacklist_manager").(managers.BlacklistManager)
	sender := ctn.Get("v4_sender").(v4Protocol.Sender)
	logger := ctn.Get("v4_logger").(v4Logger.Logger)
	limiter := ctn.Get("v4_limiter").(helpers.Limiter)

	return helpers.NewValidator(whitelist, blacklist, sender, logger, limiter)
}

func V4aConnectHandler(ctn di.Container) (interface{}, error) {
	logger := ctn.Get("v4a_logger").(v4aLogger.Logger)
	sender := ctn.Get("v4a_sender").(v4aProtocol.Sender)
	errorHandler := ctn.Get("v4a_error_handler").(v4aHandlers.ErrorHandler)
	transmitter := ctn.Get("v4a_transmitter").(v4aHelpers.Transmitter)

	return v4aHandlers.NewConnectHandler(
		logger,
		sender,
		errorHandler,
		transmitter,
	)
}

func V4aBindHandler(ctn di.Container) (interface{}, error) {
	logger := ctn.Get("v4a_logger").(v4aLogger.Logger)
	addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
	sender := ctn.Get("v4a_sender").(v4aProtocol.Sender)
	errorHandler := ctn.Get("v4a_error_handler").(v4aHandlers.ErrorHandler)
	bindManager := ctn.Get("bind_manager").(managers.BindManager)
	transmitter := ctn.Get("v4a_transmitter").(v4aHelpers.Transmitter)

	return v4aHandlers.NewBindHandler(
		logger,
		addressUtils,
		sender,
		errorHandler,
		bindManager,
		transmitter,
	)
}

func V4aHandler(ctn di.Container) (interface{}, error) {
	parser := ctn.Get("v4a_parser").(v4aProtocol.Parser)
	logger := ctn.Get("v4a_logger").(v4aLogger.Logger)
	connectHandler := ctn.Get("v4a_connect_handler").(v4aHandlers.ConnectHandler)
	bindHandler := ctn.Get("v4a_bind_handler").(v4aHandlers.BindHandler)
	sender := ctn.Get("v4a_sender").(v4aProtocol.Sender)
	errorHandler := ctn.Get("v4a_error_handler").(v4aHandlers.ErrorHandler)
	validator := ctn.Get("v4a_validator").(v4aHelpers.Validator)
	cleaner := ctn.Get("v4a_cleaner").(v4aHelpers.Cleaner)
	_replicator := ctn.Get("v4a_config_replicator")

	if _replicator == nil {
		return nil, nil
	}

	replicator := _replicator.(v4aConfig.ConfigReplicator)

	return v4aHandlers.NewHandler(
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
}

func V4aErrorHandler(ctn di.Container) (interface{}, error) {
	sender := ctn.Get("v4a_sender").(v4aProtocol.Sender)
	logger := ctn.Get("v4a_logger").(v4aLogger.Logger)
	errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

	return v4aHandlers.NewErrorHandler(
		logger,
		sender,
		errorUtils,
	)
}

func V4aCleaner(ctn di.Container) (interface{}, error) {
	manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

	return v4aHelpers.NewCleaner(manager)
}

func V4aLimiter(ctn di.Container) (interface{}, error) {
	manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

	return v4aHelpers.NewLimiter(manager)
}

func V4aTransmitter(ctn di.Container) (interface{}, error) {
	connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
	bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
	bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

	return v4aHelpers.NewTransmitter(connect, bind, bindRate)
}

func V4aValidator(ctn di.Container) (interface{}, error) {
	whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)
	blacklist := ctn.Get("blacklist_manager").(managers.BlacklistManager)
	sender := ctn.Get("v4a_sender").(v4aProtocol.Sender)
	logger := ctn.Get("v4a_logger").(v4aLogger.Logger)
	limiter := ctn.Get("v4a_limiter").(v4aHelpers.Limiter)

	return v4aHelpers.NewValidator(whitelist, blacklist, sender, logger, limiter)
}

func AuthenticationHandler(ctn di.Container) (interface{}, error) {
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
}

func V5ConnectHandler(ctn di.Container) (interface{}, error) {
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
}

func V5BindHandler(ctn di.Container) (interface{}, error) {
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
}

func V5UdpAssociationHandler(ctn di.Container) (interface{}, error) {
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
}

func V5Handler(ctn di.Container) (interface{}, error) {
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
}

func PasswordAuthenticator(ctn di.Container) (interface{}, error) {
	errorHandler := ctn.Get("v5_error_handler").(v5Handlers.ErrorHandler)
	receiver := ctn.Get("auth_password_receiver").(password.Receiver)
	sender := ctn.Get("auth_password_sender").(password.Sender)

	return authenticator.NewPasswordAuthenticator(
		errorHandler,
		sender,
		receiver,
	)
}

func NoAuthAuthenticator(ctn di.Container) (interface{}, error) {
	return authenticator.NewNoAuthAuthenticator()
}

func V5Cleaner(ctn di.Container) (interface{}, error) {
	manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

	return v5Helpers.NewCleaner(manager)
}

func V5Limiter(ctn di.Container) (interface{}, error) {
	manager := ctn.Get("connections_manager").(*managers.ConnectionsManager)

	return v5Helpers.NewLimiter(manager)
}

func V5Transmitter(ctn di.Container) (interface{}, error) {
	connect := ctn.Get("transfer_connect_handler").(transfer.ConnectHandler)
	bind := ctn.Get("transfer_bind_handler").(transfer.BindHandler)
	bindRate := ctn.Get("bind_rate_manager").(managers.BindRateManager)

	return v5Helpers.NewTransmitter(connect, bind, bindRate)
}

func V5Validator(ctn di.Container) (interface{}, error) {
	whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)
	blacklist := ctn.Get("blacklist_manager").(managers.BlacklistManager)
	sender := ctn.Get("v5_sender").(v5Protocol.Sender)
	logger := ctn.Get("v5_logger").(v5Logger.Logger)
	limiter := ctn.Get("v5_limiter").(v5Helpers.Limiter)

	return v5Helpers.NewValidator(whitelist, blacklist, sender, logger, limiter)
}

func V5ErrorHandler(ctn di.Container) (interface{}, error) {
	sender := ctn.Get("v5_sender").(v5Protocol.Sender)
	logger := ctn.Get("v5_logger").(v5Logger.Logger)
	errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

	return v5Handlers.NewErrorHandler(
		logger,
		sender,
		errorUtils,
	)
}
