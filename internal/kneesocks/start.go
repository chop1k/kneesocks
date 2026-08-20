package kneesocks

import (
	tcp2 "socks/internal/kneesocks/config/tcp"
	"socks/internal/kneesocks/config/tree"
	udp2 "socks/internal/kneesocks/config/udp"
	v9 "socks/internal/kneesocks/config/v4"
	v4a4 "socks/internal/kneesocks/config/v4a"
	v10 "socks/internal/kneesocks/config/v5"
	"socks/internal/kneesocks/handlers"
	v4 "socks/internal/kneesocks/handlers/v4"
	"socks/internal/kneesocks/handlers/v4/helpers"
	"socks/internal/kneesocks/handlers/v4a"
	helpers2 "socks/internal/kneesocks/handlers/v4a/helpers"
	v5Handlers "socks/internal/kneesocks/handlers/v5"
	"socks/internal/kneesocks/handlers/v5/authenticator"
	helpers3 "socks/internal/kneesocks/handlers/v5/helpers"
	"socks/internal/kneesocks/logger/tcp"
	"socks/internal/kneesocks/logger/udp"
	v7 "socks/internal/kneesocks/logger/v4"
	v4a3 "socks/internal/kneesocks/logger/v4a"
	v8 "socks/internal/kneesocks/logger/v5"
	"socks/internal/kneesocks/managers"
	"socks/internal/kneesocks/server"
	"socks/internal/kneesocks/transfer"
	"socks/pkg/protocol"
	"socks/pkg/protocol/auth/password"
	v6 "socks/pkg/protocol/v4"
	v4a2 "socks/pkg/protocol/v4a"
	v5 "socks/pkg/protocol/v5"
	"socks/pkg/utils"

	"github.com/go-playground/validator/v10"
)

func Start() {
	_validator := validator.New()

	config, err := tree.NewConfig(*_validator, "/etc/kneesocks/config.json")

	if err != nil {
		panic(err)
	}

	udpConfig := udp2.NewLoggerConfig(config.Log)
	udpZeroLogger, err := udp.BuildZerolog(udpConfig)

	if err != nil {
		panic(err)
	}

	errorsLogger := udp.NewErrorsLogger(udpZeroLogger)
	listenLogger := udp.NewListenLogger(udpZeroLogger)
	packetLogger := udp.NewPacketLogger(udpZeroLogger)
	udpLogger := udp.NewLogger(errorsLogger, listenLogger, packetLogger)

	addressUtils := utils.NewUtils()

	udpHandler := udp2.NewHandler()

	parsedUdpConfig := udpHandler.Handle(config.Udp)

	udpReplicator := udp2.NewConfigReplicator(parsedUdpConfig.Buffer, parsedUdpConfig.Deadline)

	packetHandler := handlers.NewPacketHandler(v5.NewParser(addressUtils), v5.NewBuilder(), addressUtils, managers.NewUdpClientManager(), managers.NewUdpHostManager(), udpLogger, udpReplicator)

	udpServer := server.NewUdpServer(udpLogger, packetHandler, parsedUdpConfig.Bind, udpReplicator)

	// this shit will be cleaned up later, for now just make it work
	go udpServer.Listen()

	tcpHandler := tcp2.NewHandler()
	tcpConfig := tcpHandler.Handle(config.Tcp)

	v4LoggerConfig := v9.NewLoggerConfig(config.Log)
	v4ZeroLogger, err := v7.BuildZerolog(v4LoggerConfig)

	if err != nil {
		panic(err)
	}

	v4Logger := v7.NewLogger(v7.NewBindLogger(v4ZeroLogger), v7.NewConnectLogger(v4ZeroLogger), v7.NewErrorsLogger(v4ZeroLogger), v7.NewRestrictionsLogger(v4ZeroLogger), v7.NewTransferLogger(v4ZeroLogger))
	v4Sender := v6.NewSender(tcpConfig.Bind, v6.NewBuilder())
	v4ErrorHandler := v4.NewErrorHandler(v4Logger, v4Sender, utils.NewErrorUtils())
	v4TransferConnectHandler := transfer.NewConnectHandler(transfer.NewHandler())
	v4BindRate := managers.NewBindRateManager()
	v4TransferBindHandler := transfer.NewBindHandler(v4BindRate, transfer.NewHandler())
	v4Transmitter := helpers.NewTransmitter(v4TransferConnectHandler, v4TransferBindHandler, v4BindRate)
	v4ConnectHandler := v4.NewConnectHandler(v4Logger, v4Sender, v4ErrorHandler, v4Transmitter)
	v4BindManager := managers.NewBindManager()
	v4BindHandler := v4.NewBindHandler(v4Logger, addressUtils, v4Sender, v4ErrorHandler, v4BindManager, v4Transmitter)
	v4WhiteList := managers.NewWhitelistManager()
	v4BlackList := managers.NewBlacklistManager()
	v4ConnectionManager := managers.NewConnectionsManager()
	v4Limiter := helpers.NewLimiter(v4ConnectionManager)
	v4Validator := helpers.NewValidator(v4WhiteList, v4BlackList, v4Sender, v4Logger, v4Limiter)
	v4Cleaner := helpers.NewCleaner(v4ConnectionManager)
	v4ConfigHandler := v9.NewHandler()
	v4ConfigReplicator := v9.NewConfigReplicator(v4ConfigHandler.Handle(config.SocksV4))

	v4Handler := v4.NewHandler(v6.NewParser(), v4Logger, v4ConnectHandler, v4BindHandler, v4Sender, v4ErrorHandler, v4Validator, v4Cleaner, v4ConfigReplicator)

	v4aLoggerConfig := v4a4.NewLoggerConfig(config.Log)
	v4aZeroLogger, err := v4a3.BuildZerolog(v4aLoggerConfig)

	if err != nil {
		panic(err)
	}

	v4aParser := v4a2.NewParser()
	v4aLogger := v4a3.NewLogger(v4a3.NewBindLogger(v4aZeroLogger), v4a3.NewConnectLogger(v4aZeroLogger), v4a3.NewErrorsLogger(v4aZeroLogger), v4a3.NewRestrictionsLogger(v4aZeroLogger), v4a3.NewTransferLogger(v4aZeroLogger))
	v4aSender := v4a2.NewSender(tcpConfig.Bind, v4a2.NewBuilder())
	v4aTransferConnectHandler := transfer.NewConnectHandler(transfer.NewHandler())
	v4aBindRate := managers.NewBindRateManager()
	v4aTransferBindHandler := transfer.NewBindHandler(v4aBindRate, transfer.NewHandler())
	v4aTransmitter := helpers2.NewTransmitter(v4aTransferConnectHandler, v4aTransferBindHandler, v4aBindRate)
	v4aErrorHandler := v4a.NewErrorHandler(v4aLogger, v4aSender, utils.NewErrorUtils())
	v4aConnectHandler := v4a.NewConnectHandler(v4aLogger, v4aSender, v4aErrorHandler, v4aTransmitter)
	v4aBindHandler := v4a.NewBindHandler(v4aLogger, addressUtils, v4aSender, v4aErrorHandler, managers.NewBindManager(), v4aTransmitter)
	v4aWhiteList := managers.NewWhitelistManager()
	v4aBlackList := managers.NewBlacklistManager()
	v4aConnectionsManager := managers.NewConnectionsManager()
	v4aValidator := helpers2.NewValidator(v4aWhiteList, v4aBlackList, v4aSender, v4aLogger, helpers2.NewLimiter(v4aConnectionsManager))
	v4aCleaner := helpers2.NewCleaner(v4aConnectionsManager)
	v4aConfigHandler := v4a4.NewHandler()
	v4aReplicator := v4a4.NewConfigReplicator(v4aConfigHandler.Handle(config.SocksV4a))

	v4aHandler := v4a.NewHandler(v4aParser, v4aLogger, v4aConnectHandler, v4aBindHandler, v4aSender, v4aErrorHandler, v4aValidator, v4aCleaner, v4aReplicator)

	v5LoggerConfig := v10.NewLoggerConfig(config.Log)
	v5ZeroLogger, err := v8.BuildZerolog(v5LoggerConfig)

	if err != nil {
		panic(err)
	}

	v5Parser := v5.NewParser(addressUtils)
	v5Logger := v8.NewLogger(v8.NewAssociationLogger(v5ZeroLogger), v8.NewAuthLogger(v5ZeroLogger), v8.NewBindLogger(v5ZeroLogger), v8.NewConnectLogger(v5ZeroLogger), v8.NewErrorsLogger(v5ZeroLogger), v8.NewRestrictionsLogger(v5ZeroLogger), v8.NewTransferLogger(v5ZeroLogger))
	v5Sender := v5.NewSender(tcpConfig.Bind, udp2.NewHandler().Handle(config.Udp).Bind, v5.NewBuilder())
	v5ErrorHandler := v5Handlers.NewErrorHandler(v5Logger, v5Sender, utils.NewErrorUtils())
	v5Receiver := v5.NewReceiver(v5.NewParser(addressUtils), utils.NewBufferReader())
	v5AuthenticationHandler := v5Handlers.NewAuthenticationHandler(v5ErrorHandler, authenticator.NewPasswordAuthenticator(v5ErrorHandler, password.NewSender(password.NewBuilder()), password.NewReceiver(password.NewParser(), utils.NewBufferReader())), authenticator.NewNoAuthAuthenticator(), v5Sender)
	v5TransferConnectHandler := transfer.NewConnectHandler(transfer.NewHandler())
	v5BindManager := managers.NewBindRateManager()
	v5TransferBindHandler := transfer.NewBindHandler(v5BindManager, transfer.NewHandler())
	v5Transmitter := helpers3.NewTransmitter(v5TransferConnectHandler, v5TransferBindHandler, v5BindManager)
	v5ConnectHandler := v5Handlers.NewConnectHandler(v5Logger, addressUtils, v5Sender, v5ErrorHandler, v5Transmitter)
	v5BindHandler := v5Handlers.NewBindHandler(addressUtils, v5Logger, v5Sender, v5ErrorHandler, managers.NewBindManager(), v5Transmitter)
	v5UdpAssociationHandler := v5Handlers.NewUdpAssociationHandler(addressUtils, managers.NewUdpClientManager(), v5Logger, v5Sender, v5ErrorHandler)
	v5WhiteList := managers.NewWhitelistManager()
	v5BlackList := managers.NewBlacklistManager()
	v5ConnectionsManager := managers.NewConnectionsManager()
	v5Validator := helpers3.NewValidator(v5WhiteList, v5BlackList, v5Sender, v5Logger, helpers3.NewLimiter(v5ConnectionsManager))
	v5Cleaner := helpers3.NewCleaner(v5ConnectionsManager)
	v5Replicator := v10.NewConfigReplicator(v10.NewHandler().Handle(config.SocksV5))

	v5Handler := v5Handlers.NewHandler(v5Parser, v5AuthenticationHandler, v5Logger, v5ConnectHandler, v5BindHandler, v5UdpAssociationHandler, v5ErrorHandler, v5Sender, v5Receiver, v5Validator, v5Cleaner, v5Replicator)

	tcpLoggerConfig := tcp2.NewLoggerConfig(config.Log)
	tcpZeroLogger, err := tcp.BuildZerolog(tcpLoggerConfig)

	if err != nil {
		panic(err)
	}

	connectionLogger := tcp.NewConnectionLogger(tcpZeroLogger)
	tcpErrorsLogger := tcp.NewErrorsLogger(tcpZeroLogger)
	tcpListenLogger := tcp.NewListenLogger(tcpZeroLogger)

	tcpLogger := tcp.NewLogger(connectionLogger, tcpErrorsLogger, tcpListenLogger)

	tcpReplicator := tcp2.NewConfigReplicator(tcpConfig.Deadline)

	bindManager := managers.NewBindManager()
	bindTransfer := transfer.NewBindHandler(managers.NewBindRateManager(), transfer.NewHandler())

	bindHandler := handlers.NewBindHandler(addressUtils, tcpLogger, bindTransfer, bindManager)
	connectionHandler := handlers.NewConnectionHandler(v4Handler, v4aHandler, v5Handler, tcpLogger, protocol.NewReceiver(utils.NewBufferReader()), bindHandler, tcpReplicator)

	tcpServer := server.NewTcpServer(connectionHandler, tcpLogger, tcpConfig.Bind)

	tcpServer.Listen()
}
