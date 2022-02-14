package dependency

import (
	"github.com/rs/zerolog"
	"github.com/sarulabs/di"
	tcpConfig "socks/internal/kneesocks/config/tcp"
	udpConfig "socks/internal/kneesocks/config/udp"
	v4Config "socks/internal/kneesocks/config/v4"
	v4aConfig "socks/internal/kneesocks/config/v4a"
	v5Config "socks/internal/kneesocks/config/v5"
	"socks/internal/kneesocks/logger/tcp"
	"socks/internal/kneesocks/logger/udp"
	"socks/internal/kneesocks/logger/v4"
	"socks/internal/kneesocks/logger/v4a"
	"socks/internal/kneesocks/logger/v5"
)

func registerZeroLog(builder di.Builder) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	tcpDef := di.Def{
		Name:  "tcp_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("tcp_logger_config").(tcpConfig.LoggerConfig)

			return tcp.BuildZerolog(cfg)
		},
	}

	udpDef := di.Def{
		Name:  "udp_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("udp_logger_config").(udpConfig.LoggerConfig)

			return udp.BuildZerolog(cfg)
		},
	}

	v4Def := di.Def{
		Name:  "v4_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_logger_config").(v4Config.LoggerConfig)

			return v4.BuildZerolog(cfg)
		},
	}

	v4aDef := di.Def{
		Name:  "v4a_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_logger_config").(v4aConfig.LoggerConfig)

			return v4a.BuildZerolog(cfg)
		},
	}

	v5Def := di.Def{
		Name:  "v5_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_logger_config").(v5Config.LoggerConfig)

			return v5.BuildZerolog(cfg)
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("tcp_zero_logger").(zerolog.Logger)

			return tcp.NewConnectionLogger(zero)
		},
	}

	errorsLoggerDef := di.Def{
		Name:  "tcp_errors_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("tcp_zero_logger").(zerolog.Logger)

			return tcp.NewErrorsLogger(zero)
		},
	}

	listenLoggerDef := di.Def{
		Name:  "tcp_listen_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("tcp_zero_logger").(zerolog.Logger)

			return tcp.NewListenLogger(zero)
		},
	}

	loggerDef := di.Def{
		Name:  "tcp_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connection := ctn.Get("tcp_connection_logger").(tcp.ConnectionLogger)
			errors := ctn.Get("tcp_errors_logger").(tcp.ErrorsLogger)
			listen := ctn.Get("tcp_listen_logger").(tcp.ListenLogger)

			return tcp.NewLogger(connection, errors, listen)
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("udp_zero_logger").(zerolog.Logger)

			return udp.NewErrorsLogger(zero)
		},
	}

	listenLoggerDef := di.Def{
		Name:  "udp_listen_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("udp_zero_logger").(zerolog.Logger)

			return udp.NewListenLogger(zero)
		},
	}

	packetLoggerDef := di.Def{
		Name:  "udp_packet_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("udp_zero_logger").(zerolog.Logger)

			return udp.NewPacketLogger(zero)
		},
	}

	loggerDef := di.Def{
		Name:  "udp_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			errors := ctn.Get("udp_errors_logger").(udp.ErrorsLogger)
			listen := ctn.Get("udp_listen_logger").(udp.ListenLogger)
			packet := ctn.Get("udp_packet_logger").(udp.PacketLogger)

			return udp.NewLogger(errors, listen, packet)
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

			return v4.NewBindLogger(zero)
		},
	}

	connectDef := di.Def{
		Name:  "v4_connect_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

			return v4.NewConnectLogger(zero)
		},
	}

	errorsDef := di.Def{
		Name:  "v4_errors_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

			return v4.NewErrorsLogger(zero)
		},
	}

	restrictionsDef := di.Def{
		Name:  "v4_restrictions_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

			return v4.NewRestrictionsLogger(zero)
		},
	}

	transferDef := di.Def{
		Name:  "v4_transfer_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

			return v4.NewTransferLogger(zero)
		},
	}

	loggerDef := di.Def{
		Name:  "v4_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			bind := ctn.Get("v4_bind_logger").(v4.BindLogger)
			connect := ctn.Get("v4_connect_logger").(v4.ConnectLogger)
			errors := ctn.Get("v4_errors_logger").(v4.ErrorsLogger)
			restrictions := ctn.Get("v4_restrictions_logger").(v4.RestrictionsLogger)
			transfer := ctn.Get("v4_transfer_logger").(v4.TransferLogger)

			return v4.NewLogger(bind, connect, errors, restrictions, transfer)
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

			return v4a.NewBindLogger(zero)
		},
	}

	connectDef := di.Def{
		Name:  "v4a_connect_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

			return v4a.NewConnectLogger(zero)
		},
	}

	errorsDef := di.Def{
		Name:  "v4a_errors_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

			return v4a.NewErrorsLogger(zero)
		},
	}

	restrictionsDef := di.Def{
		Name:  "v4a_restrictions_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

			return v4a.NewRestrictionsLogger(zero)
		},
	}

	transferDef := di.Def{
		Name:  "v4a_transfer_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

			return v4a.NewTransferLogger(zero)
		},
	}

	loggerDef := di.Def{
		Name:  "v4a_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			bind := ctn.Get("v4a_bind_logger").(v4a.BindLogger)
			connect := ctn.Get("v4a_connect_logger").(v4a.ConnectLogger)
			errors := ctn.Get("v4a_errors_logger").(v4a.ErrorsLogger)
			restrictions := ctn.Get("v4a_restrictions_logger").(v4a.RestrictionsLogger)
			transfer := ctn.Get("v4a_transfer_logger").(v4a.TransferLogger)

			return v4a.NewLogger(bind, connect, errors, restrictions, transfer)
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewAssociationLogger(zero)
		},
	}

	authDef := di.Def{
		Name:  "v5_auth_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewAuthLogger(zero)
		},
	}

	bindDef := di.Def{
		Name:  "v5_bind_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewBindLogger(zero)
		},
	}

	connectDef := di.Def{
		Name:  "v5_connect_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewConnectLogger(zero)
		},
	}

	errorsDef := di.Def{
		Name:  "v5_errors_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewErrorsLogger(zero)
		},
	}

	restrictionsDef := di.Def{
		Name:  "v5_restrictions_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewRestrictionsLogger(zero)
		},
	}

	transferDef := di.Def{
		Name:  "v5_transfer_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewTransferLogger(zero)
		},
	}
	loggerDef := di.Def{
		Name:  "v5_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			association := ctn.Get("v5_association_logger").(v5.AssociationLogger)
			auth := ctn.Get("v5_auth_logger").(v5.AuthLogger)
			bind := ctn.Get("v5_bind_logger").(v5.BindLogger)
			connect := ctn.Get("v5_connect_logger").(v5.ConnectLogger)
			errors := ctn.Get("v5_errors_logger").(v5.ErrorsLogger)
			restrictions := ctn.Get("v5_restrictions_logger").(v5.RestrictionsLogger)
			transfer := ctn.Get("v5_transfer_logger").(v5.TransferLogger)

			return v5.NewLogger(association, auth, bind, connect, errors, restrictions, transfer)
		},
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
