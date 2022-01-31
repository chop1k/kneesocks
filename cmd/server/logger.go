package main

import (
	"github.com/rs/zerolog"
	"github.com/sarulabs/di"
	"io"
	"os"
	"socks/config"
	"socks/config/tcp"
	"socks/config/udp"
	v43 "socks/config/v4"
	v4a3 "socks/config/v4a"
	v53 "socks/config/v5"
	"socks/logger"
	tcp2 "socks/logger/tcp"
	udp2 "socks/logger/udp"
	v4 "socks/logger/v4"
	"socks/logger/v4a"
	v5 "socks/logger/v5"
)

func registerZeroLog(builder di.Builder) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	buildLogger := func(level int, loggers []io.Writer) (zerolog.Logger, error) {
		return zerolog.New(zerolog.MultiLevelWriter(loggers...)).
			With().
			Timestamp().
			Logger().
			Level(zerolog.Level(level)), nil
	}

	serverZeroLoggerDef := di.Def{
		Name:  "server_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("server_logger_config").(config.ServerLoggerConfig)

			level, err := cfg.GetLevel()

			var loggers []io.Writer

			if err != nil {
				return buildLogger(126, loggers)
			}

			if output, err := cfg.GetConsoleOutput(); err == nil {
				loggers = append(loggers, zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: output.TimeFormat,
				})
			} else {
				if err == config.ServerLoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			if output, err := cfg.GetFileOutput(); err == nil {
				file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

				if err != nil {
					return nil, err
				}

				loggers = append(loggers, file)
			} else {
				if err == config.ServerLoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			return buildLogger(level, loggers)
		},
	}

	tcpZeroLoggerDef := di.Def{
		Name:  "tcp_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("tcp_logger_config").(tcp.LoggerConfig)

			level, err := cfg.GetLevel()

			var loggers []io.Writer

			if err != nil {
				return buildLogger(126, loggers)
			}

			if output, err := cfg.GetConsoleOutput(); err == nil {
				loggers = append(loggers, zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: output.TimeFormat,
				})
			} else {
				if err == tcp.LoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			if output, err := cfg.GetFileOutput(); err == nil {
				file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

				if err != nil {
					return nil, err
				}

				loggers = append(loggers, file)
			} else {
				if err == tcp.LoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			return buildLogger(level, loggers)
		},
	}

	udpZeroLoggerDef := di.Def{
		Name:  "udp_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("udp_logger_config").(udp.LoggerConfig)

			level, err := cfg.GetLevel()

			var loggers []io.Writer

			if err != nil {
				return buildLogger(126, loggers)
			}

			if output, err := cfg.GetConsoleOutput(); err == nil {
				loggers = append(loggers, zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: output.TimeFormat,
				})
			} else {
				if err == udp.LoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			if output, err := cfg.GetFileOutput(); err == nil {
				file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

				if err != nil {
					return nil, err
				}

				loggers = append(loggers, file)
			} else {
				if err == udp.LoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			return buildLogger(level, loggers)
		},
	}

	v4ZeroLoggerDef := di.Def{
		Name:  "v4_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_logger_config").(v43.LoggerConfig)

			level, err := cfg.GetLevel()

			var loggers []io.Writer

			if err != nil {
				return buildLogger(126, loggers)
			}

			if output, err := cfg.GetConsoleOutput(); err == nil {
				loggers = append(loggers, zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: output.TimeFormat,
				})
			} else {
				if err == v43.LoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			if output, err := cfg.GetFileOutput(); err == nil {
				file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

				if err != nil {
					return nil, err
				}

				loggers = append(loggers, file)
			} else {
				if err == v43.LoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			return buildLogger(level, loggers)
		},
	}

	v4aZeroLoggerDef := di.Def{
		Name:  "v4a_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_logger_config").(v4a3.LoggerConfig)

			level, err := cfg.GetLevel()

			var loggers []io.Writer

			if err != nil {
				return buildLogger(126, loggers)
			}

			if output, err := cfg.GetConsoleOutput(); err == nil {
				loggers = append(loggers, zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: output.TimeFormat,
				})
			} else {
				if err == v4a3.LoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			if output, err := cfg.GetFileOutput(); err == nil {
				file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

				if err != nil {
					return nil, err
				}

				loggers = append(loggers, file)
			} else {
				if err == v4a3.LoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			return buildLogger(level, loggers)
		},
	}

	v5ZeroLoggerDef := di.Def{
		Name:  "v5_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_logger_config").(v53.LoggerConfig)

			level, err := cfg.GetLevel()

			var loggers []io.Writer

			if err != nil {
				return buildLogger(126, loggers)
			}

			if output, err := cfg.GetConsoleOutput(); err == nil {
				loggers = append(loggers, zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: output.TimeFormat,
				})
			} else {
				if err == v53.LoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			if output, err := cfg.GetFileOutput(); err == nil {
				file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

				if err != nil {
					return nil, err
				}

				loggers = append(loggers, file)
			} else {
				if err == v53.LoggerDisabledError {
					return buildLogger(126, loggers)
				}
			}

			return buildLogger(level, loggers)
		},
	}

	err := builder.Add(
		serverZeroLoggerDef,
		tcpZeroLoggerDef,
		udpZeroLoggerDef,
		v4ZeroLoggerDef,
		v4aZeroLoggerDef,
		v5ZeroLoggerDef,
	)

	if err != nil {
		panic(err)
	}

}

func registerLogger(builder di.Builder) {
	serverLoggerDef := di.Def{
		Name:  "server_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("server_zero_logger").(zerolog.Logger)

			return logger.NewBaseServerLogger(zero)
		},
	}

	err := builder.Add(
		serverLoggerDef,
	)

	if err != nil {
		panic(err)
	}

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

			return tcp2.NewBaseConnectionLogger(zero)
		},
	}

	errorsLoggerDef := di.Def{
		Name:  "tcp_errors_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("tcp_zero_logger").(zerolog.Logger)

			return tcp2.NewBaseErrorsLogger(zero)
		},
	}

	listenLoggerDef := di.Def{
		Name:  "tcp_listen_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("tcp_zero_logger").(zerolog.Logger)

			return tcp2.NewBaseListenLogger(zero)
		},
	}

	tcpLoggerDef := di.Def{
		Name:  "tcp_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connection := ctn.Get("tcp_connection_logger").(tcp2.ConnectionLogger)
			errors := ctn.Get("tcp_errors_logger").(tcp2.ErrorsLogger)
			listen := ctn.Get("tcp_listen_logger").(tcp2.ListenLogger)

			return tcp2.NewLogger(connection, errors, listen)
		},
	}

	err := builder.Add(
		connectionLoggerDef,
		errorsLoggerDef,
		listenLoggerDef,
		tcpLoggerDef,
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

			return udp2.NewBaseErrorsLogger(zero)
		},
	}

	listenLoggerDef := di.Def{
		Name:  "udp_listen_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("udp_zero_logger").(zerolog.Logger)

			return udp2.NewBaseListenLogger(zero)
		},
	}

	packetLoggerDef := di.Def{
		Name:  "udp_packet_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("udp_zero_logger").(zerolog.Logger)

			return udp2.NewBasePacketLogger(zero)
		},
	}

	udpLoggerDef := di.Def{
		Name:  "udp_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			errors := ctn.Get("udp_errors_logger").(udp2.ErrorsLogger)
			listen := ctn.Get("udp_listen_logger").(udp2.ListenLogger)
			packet := ctn.Get("udp_packet_logger").(udp2.PacketLogger)

			return udp2.NewLogger(errors, listen, packet)
		},
	}

	err := builder.Add(
		errorsLoggerDef,
		listenLoggerDef,
		packetLoggerDef,
		udpLoggerDef,
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

			return v4.NewBaseBindLogger(zero)
		},
	}

	connectDef := di.Def{
		Name:  "v4_connect_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

			return v4.NewBaseConnectLogger(zero)
		},
	}

	errorsDef := di.Def{
		Name:  "v4_errors_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

			return v4.NewBaseErrorsLogger(zero)
		},
	}

	restrictionsDef := di.Def{
		Name:  "v4_restrictions_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

			return v4.NewBaseRestrictionsLogger(zero)
		},
	}

	transferDef := di.Def{
		Name:  "v4_transfer_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

			return v4.NewBaseTransferLogger(zero)
		},
	}

	v4LoggerDef := di.Def{
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
		v4LoggerDef,
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

			return v4a.NewBaseBindLogger(zero)
		},
	}

	connectDef := di.Def{
		Name:  "v4a_connect_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

			return v4a.NewBaseConnectLogger(zero)
		},
	}

	errorsDef := di.Def{
		Name:  "v4a_errors_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

			return v4a.NewBaseErrorsLogger(zero)
		},
	}

	restrictionsDef := di.Def{
		Name:  "v4a_restrictions_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

			return v4a.NewBaseRestrictionsLogger(zero)
		},
	}

	transferDef := di.Def{
		Name:  "v4a_transfer_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

			return v4a.NewBaseTransferLogger(zero)
		},
	}

	v4aLoggerDef := di.Def{
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
		v4aLoggerDef,
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

			return v5.NewBaseAssociationLogger(zero)
		},
	}

	authDef := di.Def{
		Name:  "v5_auth_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewBaseAuthLogger(zero)
		},
	}

	bindDef := di.Def{
		Name:  "v5_bind_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewBaseBindLogger(zero)
		},
	}

	connectDef := di.Def{
		Name:  "v5_connect_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewBaseConnectLogger(zero)
		},
	}

	errorsDef := di.Def{
		Name:  "v5_errors_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewBaseErrorsLogger(zero)
		},
	}

	restrictionsDef := di.Def{
		Name:  "v5_restrictions_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewBaseRestrictionsLogger(zero)
		},
	}

	transferDef := di.Def{
		Name:  "v5_transfer_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return v5.NewBaseTransferLogger(zero)
		},
	}
	v5LoggerDef := di.Def{
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
		v5LoggerDef,
	)

	if err != nil {
		panic(err)
	}
}
