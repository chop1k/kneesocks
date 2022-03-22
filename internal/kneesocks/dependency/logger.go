package dependency

import (
	"socks/internal/kneesocks/dependency/build"

	"github.com/rs/zerolog"
	"github.com/sarulabs/di"
)

func registerZeroLog(builder di.Builder) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	tcpDef := di.Def{
		Name:  "tcp_zero_logger",
		Scope: di.App,
		Build: build.TcpZeroLogger,
	}

	udpDef := di.Def{
		Name:  "udp_zero_logger",
		Scope: di.App,
		Build: build.UdpZeroLogger,
	}

	v4Def := di.Def{
		Name:  "v4_zero_logger",
		Scope: di.App,
		Build: build.V4ZeroLogger,
	}

	v4aDef := di.Def{
		Name:  "v4a_zero_logger",
		Scope: di.App,
		Build: build.V4aZeroLogger,
	}

	v5Def := di.Def{
		Name:  "v5_zero_logger",
		Scope: di.App,
		Build: build.V5ZeroLogger,
	}

	err := builder.Add(
		tcpDef,
		udpDef,
		v4Def,
		v4aDef,
		v5Def,
	)

	if err != nil {
		panic(err)
	}

}

func registerLogger(builder di.Builder) {
	registerTcpLogger(builder)
	registerUdpLogger(builder)
	registerV4Logger(builder)
	registerV4aLogger(builder)
	registerV5Logger(builder)
}

func registerTcpLogger(builder di.Builder) {
	connectionLoggerDef := di.Def{
		Name:  "tcp_connection_logger",
		Scope: di.App,
		Build: build.TcpConnectionLogger,
	}

	errorsLoggerDef := di.Def{
		Name:  "tcp_errors_logger",
		Scope: di.App,
		Build: build.TcpErrorsLogger,
	}

	listenLoggerDef := di.Def{
		Name:  "tcp_listen_logger",
		Scope: di.App,
		Build: build.TcpListenLogger,
	}

	loggerDef := di.Def{
		Name:  "tcp_logger",
		Scope: di.App,
		Build: build.TcpLogger,
	}

	err := builder.Add(
		connectionLoggerDef,
		errorsLoggerDef,
		listenLoggerDef,
		loggerDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerUdpLogger(builder di.Builder) {
	errorsLoggerDef := di.Def{
		Name:  "udp_errors_logger",
		Scope: di.App,
		Build: build.UdpErrorsLogger,
	}

	listenLoggerDef := di.Def{
		Name:  "udp_listen_logger",
		Scope: di.App,
		Build: build.UdpListenLogger,
	}

	packetLoggerDef := di.Def{
		Name:  "udp_packet_logger",
		Scope: di.App,
		Build: build.UdpPacketLogger,
	}

	loggerDef := di.Def{
		Name:  "udp_logger",
		Scope: di.App,
		Build: build.UdpLogger,
	}

	err := builder.Add(
		errorsLoggerDef,
		listenLoggerDef,
		packetLoggerDef,
		loggerDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV4Logger(builder di.Builder) {
	bindDef := di.Def{
		Name:  "v4_bind_logger",
		Scope: di.App,
		Build: build.V4BindLogger,
	}

	connectDef := di.Def{
		Name:  "v4_connect_logger",
		Scope: di.App,
		Build: build.V4ConnectLogger,
	}

	errorsDef := di.Def{
		Name:  "v4_errors_logger",
		Scope: di.App,
		Build: build.V4ErrorsLogger,
	}

	restrictionsDef := di.Def{
		Name:  "v4_restrictions_logger",
		Scope: di.App,
		Build: build.V4RestrictionsLogger,
	}

	transferDef := di.Def{
		Name:  "v4_transfer_logger",
		Scope: di.App,
		Build: build.V4TransferLogger,
	}

	loggerDef := di.Def{
		Name:  "v4_logger",
		Scope: di.App,
		Build: build.V4Logger,
	}

	err := builder.Add(
		bindDef,
		connectDef,
		errorsDef,
		restrictionsDef,
		transferDef,
		loggerDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV4aLogger(builder di.Builder) {
	bindDef := di.Def{
		Name:  "v4a_bind_logger",
		Scope: di.App,
		Build: build.V4aBindLogger,
	}

	connectDef := di.Def{
		Name:  "v4a_connect_logger",
		Scope: di.App,
		Build: build.V4aConnectLogger,
	}

	errorsDef := di.Def{
		Name:  "v4a_errors_logger",
		Scope: di.App,
		Build: build.V4aErrorsLogger,
	}

	restrictionsDef := di.Def{
		Name:  "v4a_restrictions_logger",
		Scope: di.App,
		Build: build.V4aRestrictionsLogger,
	}

	transferDef := di.Def{
		Name:  "v4a_transfer_logger",
		Scope: di.App,
		Build: build.V4aTransferLogger,
	}

	loggerDef := di.Def{
		Name:  "v4a_logger",
		Scope: di.App,
		Build: build.V4aLogger,
	}

	err := builder.Add(
		bindDef,
		connectDef,
		errorsDef,
		restrictionsDef,
		transferDef,
		loggerDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV5Logger(builder di.Builder) {
	associationDef := di.Def{
		Name:  "v5_association_logger",
		Scope: di.App,
		Build: build.V5AssociationLogger,
	}

	authDef := di.Def{
		Name:  "v5_auth_logger",
		Scope: di.App,
		Build: build.V5AuthLogger,
	}

	bindDef := di.Def{
		Name:  "v5_bind_logger",
		Scope: di.App,
		Build: build.V5BindLogger,
	}

	connectDef := di.Def{
		Name:  "v5_connect_logger",
		Scope: di.App,
		Build: build.V5ConnectLogger,
	}

	errorsDef := di.Def{
		Name:  "v5_errors_logger",
		Scope: di.App,
		Build: build.V5ErrorsLogger,
	}

	restrictionsDef := di.Def{
		Name:  "v5_restrictions_logger",
		Scope: di.App,
		Build: build.V5RestrictionsLogger,
	}

	transferDef := di.Def{
		Name:  "v5_transfer_logger",
		Scope: di.App,
		Build: build.V5TransferLogger,
	}
	loggerDef := di.Def{
		Name:  "v5_logger",
		Scope: di.App,
		Build: build.V5Logger,
	}

	err := builder.Add(
		associationDef,
		authDef,
		bindDef,
		connectDef,
		errorsDef,
		restrictionsDef,
		transferDef,
		loggerDef,
	)

	if err != nil {
		panic(err)
	}
}
