package main

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/sarulabs/di"
	"io"
	"os"
	"socks/config"
	"socks/config/tree"
	"socks/logger"
	"socks/protocol/auth/password"
	v4 "socks/protocol/v4"
	"socks/protocol/v4a"
	v5 "socks/protocol/v5"
	"socks/server"
	"socks/utils"
)

func main() {
	builder, err := di.NewBuilder()

	if err != nil {
		panic(err)
	}

	registerConfig(*builder)
}

func registerConfig(builder di.Builder) {
	configPathDef := di.Def{
		Name:  "config_path",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path, ok := os.LookupEnv("socks_config")

			if !ok {
				return nil, errors.New("Config path is not specified. ")
			}

			return path, nil
		},
	}

	configTreeDef := di.Def{
		Name:  "config_tree",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			validate := ctn.Get("validator").(validator.Validate)
			configPath := ctn.Get("config_path").(string)

			return tree.NewConfig(validate, configPath)
		},
	}

	configDef := di.Def{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseConfig(cfg), nil
		},
	}

	tcpConfigDef := di.Def{
		Name:  "tcp_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseTcpConfig(cfg), nil
		},
	}

	udpConfigDef := di.Def{
		Name:  "udp_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseUdpConfig(cfg), nil
		},
	}

	v4ConfigDef := di.Def{
		Name:  "v4_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV4Config(cfg), nil
		},
	}

	v4aConfigDef := di.Def{
		Name:  "v4a_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV4aConfig(cfg), nil
		},
	}

	v5ConfigDef := di.Def{
		Name:  "v5_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV5Config(cfg), nil
		},
	}

	serverLoggerConfigDef := di.Def{
		Name:  "server_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseServerLoggerConfig(cfg)
		},
	}

	tcpLoggerConfigDef := di.Def{
		Name:  "tcp_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseTcpLoggerConfig(cfg)
		},
	}

	udpLoggerConfigDef := di.Def{
		Name:  "udp_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseUdpLoggerConfig(cfg)
		},
	}

	v4LoggerConfigDef := di.Def{
		Name:  "v4_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV4LoggerConfig(cfg)
		},
	}

	v4aLoggerConfigDef := di.Def{
		Name:  "v4a_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV4aLoggerConfig(cfg)
		},
	}

	v5LoggerConfigDef := di.Def{
		Name:  "v5_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV5LoggerConfig(cfg)
		},
	}

	whitelistConfigDef := di.Def{
		Name:  "whitelist_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseWhitelistConfig(cfg), nil
		},
	}

	blacklistConfigDef := di.Def{
		Name:  "blacklist_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseBlacklistConfig(cfg), nil
		},
	}

	validatorDef := di.Def{
		Name:  "validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return *validator.New(), nil
		},
	}

	usersConfigDef := di.Def{
		Name:  "users_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseUsersConfig(cfg)
		},
	}

	err := builder.Add(
		configPathDef,
		validatorDef,
		configTreeDef,
		configDef,
		serverLoggerConfigDef,
		tcpConfigDef,
		udpConfigDef,
		v4ConfigDef,
		v4aConfigDef,
		v5ConfigDef,
		tcpLoggerConfigDef,
		udpLoggerConfigDef,
		v4LoggerConfigDef,
		v4aLoggerConfigDef,
		v5LoggerConfigDef,
		whitelistConfigDef,
		blacklistConfigDef,
		usersConfigDef,
	)

	if err != nil {
		panic(err)
	}

	registerZeroLog(builder)
}

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
			cfg := ctn.Get("tcp_logger_config").(config.TcpLoggerConfig)

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
				if err == config.TcpLoggerDisabledError {
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
				if err == config.TcpLoggerDisabledError {
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
			cfg := ctn.Get("udp_logger_config").(config.UdpLoggerConfig)

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
				if err == config.UdpLoggerDisabledError {
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
				if err == config.UdpLoggerDisabledError {
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
			cfg := ctn.Get("v4_logger_config").(config.SocksV4LoggerConfig)

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
				if err == config.SocksV4LoggerDisabledError {
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
				if err == config.SocksV4LoggerDisabledError {
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
			cfg := ctn.Get("v4a_logger_config").(config.SocksV4aLoggerConfig)

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
				if err == config.SocksV4aLoggerDisabledError {
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
				if err == config.SocksV4aLoggerDisabledError {
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
			cfg := ctn.Get("v5_logger_config").(config.SocksV5LoggerConfig)

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
				if err == config.SocksV5LoggerDisabledError {
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
				if err == config.SocksV5LoggerDisabledError {
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

	registerLogger(builder)
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

	tcpLoggerDef := di.Def{
		Name:  "tcp_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("tcp_zero_logger").(zerolog.Logger)

			return logger.NewBaseTcpLogger(zero)
		},
	}

	udpLoggerDef := di.Def{
		Name:  "udp_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("udp_zero_logger").(zerolog.Logger)

			return logger.NewBaseUdpLogger(zero)
		},
	}

	v4LoggerDef := di.Def{
		Name:  "v4_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

			return logger.NewBaseSocksV4Logger(zero)
		},
	}

	v4aLoggerDef := di.Def{
		Name:  "v4a_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

			return logger.NewBaseSocksV4aLogger(zero)
		},
	}

	v5LoggerDef := di.Def{
		Name:  "v5_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

			return logger.NewBaseSocksV5Logger(zero)
		},
	}

	err := builder.Add(
		serverLoggerDef,
		tcpLoggerDef,
		udpLoggerDef,
		v4LoggerDef,
		v4aLoggerDef,
		v5LoggerDef,
	)

	if err != nil {
		panic(err)
	}

	registerProtocol(builder)
}

func registerProtocol(builder di.Builder) {

	passwordParserDef := di.Def{
		Name:  "auth_password_parser",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return password.NewBaseParser(), nil
		},
	}

	passwordBuilderDef := di.Def{
		Name:  "auth_password_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return password.NewBaseBuilder(), nil
		},
	}

	passwordDef := di.Def{
		Name:  "auth_password",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("auth_password_parser").(password.Parser)
			builder := ctn.Get("auth_password_builder").(password.Builder)

			return password.NewPassword(parser, builder), nil
		},
	}

	v4ParserDef := di.Def{
		Name:  "v4_parser",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4.NewBaseParser(), nil
		},
	}

	v4BuilderDef := di.Def{
		Name:  "v4_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4.NewBaseBuilder(), nil
		},
	}

	v4Def := di.Def{
		Name:  "v4",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			builder := ctn.Get("v4_builder").(v4.Builder)

			return v4.NewProtocol(builder), nil
		},
	}

	v4aParserDef := di.Def{
		Name:  "v4a_parser",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4a.NewBaseParser(), nil
		},
	}

	v4aBuilderDef := di.Def{
		Name:  "v4a_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4a.NewBaseBuilder(), nil
		},
	}

	v4aDef := di.Def{
		Name:  "v4a",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			builder := ctn.Get("v4a_builder").(v4a.Builder)

			return v4a.NewProtocol(builder), nil
		},
	}

	v5ParserDef := di.Def{
		Name:  "v5_parser",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)

			return v5.NewBaseParser(addressUtils), nil
		},
	}

	v5BuilderDef := di.Def{
		Name:  "v5_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v5.NewBaseBuilder(), nil
		},
	}

	v5Def := di.Def{
		Name:  "v5",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			builder := ctn.Get("v5_builder").(v5.Builder)
			parser := ctn.Get("v5_parser").(v5.Parser)

			return v5.NewProtocol(builder, parser), nil
		},
	}

	err := builder.Add(
		passwordParserDef,
		passwordBuilderDef,
		passwordDef,
		v4ParserDef,
		v4BuilderDef,
		v4Def,
		v4aParserDef,
		v4aBuilderDef,
		v4aDef,
		v5ParserDef,
		v5BuilderDef,
		v5Def,
	)

	if err != nil {
		panic(err)
	}

	registerServer(builder)
}

func registerServer(builder di.Builder) {

	authenticationHandlerDef := di.Def{
		Name:  "authentication_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			passwd := ctn.Get("auth_password").(password.Password)
			protocol := ctn.Get("v5").(v5.Protocol)
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
			users := ctn.Get("users_config").(config.UsersConfig)
			errorHandler := ctn.Get("v5_error_handler").(server.V5ErrorHandler)

			return server.NewBaseAuthenticationHandler(
				passwd,
				protocol,
				cfg,
				users,
				errorHandler,
			), nil
		},
	}

	connectionHandlerDef := di.Def{
		Name:  "connection_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			streamHandler := ctn.Get("stream_handler").(server.StreamHandler)
			v4Handler := ctn.Get("v4_handler").(server.V4Handler)
			v4aHandler := ctn.Get("v4a_handler").(server.V4aHandler)
			v5Handler := ctn.Get("v5_handler").(server.V5Handler)
			bindManager := ctn.Get("bind_manager").(server.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			tcpLogger := ctn.Get("tcp_logger").(logger.TcpLogger)
			tcpConfig := ctn.Get("tcp_config").(config.TcpConfig)

			return server.NewBaseConnectionHandler(
				streamHandler,
				v4Handler,
				v4aHandler,
				v5Handler,
				bindManager,
				addressUtils,
				tcpLogger,
				tcpConfig,
			)
		},
	}

	packetHandlerDef := di.Def{
		Name:  "packet_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("v5_parser").(v5.Parser)
			builder := ctn.Get("v5_builder").(v5.Builder)
			udpAssociationManager := ctn.Get("udp_association_manager").(server.UdpAssociationManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)

			return server.NewBasePacketHandler(
				parser,
				builder,
				udpAssociationManager,
				addressUtils,
			), nil
		},
	}

	streamHandlerDef := di.Def{
		Name:  "stream_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return server.NewBaseStreamHandler(), nil
		},
	}

	bindManagerDef := di.Def{
		Name:  "bind_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return server.NewBindManager(), nil
		},
	}

	v4ConnectHandlerDef := di.Def{
		Name:  "v4_connect_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_config").(config.SocksV4Config)
			v4Logger := ctn.Get("v4_logger").(logger.SocksV4Logger)
			streamHandler := ctn.Get("stream_handler").(server.StreamHandler)
			sender := ctn.Get("v4_sender").(server.V4Sender)
			errorHandler := ctn.Get("v4_error_handler").(server.V4ErrorHandler)
			whitelist := ctn.Get("whitelist_manager").(server.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(server.BlacklistManager)

			return server.NewBaseV4ConnectHandler(
				cfg,
				streamHandler,
				v4Logger,
				sender,
				errorHandler,
				whitelist,
				blacklist,
			)
		},
	}

	v4BindHandlerDef := di.Def{
		Name:  "v4_bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_config").(config.SocksV4Config)
			v4Logger := ctn.Get("v4_logger").(logger.SocksV4Logger)
			streamHandler := ctn.Get("stream_handler").(server.StreamHandler)
			bindManager := ctn.Get("bind_manager").(server.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v4_sender").(server.V4Sender)
			whitelist := ctn.Get("whitelist_manager").(server.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(server.BlacklistManager)
			errorHandler := ctn.Get("v4_error_handler").(server.V4ErrorHandler)

			return server.NewBaseV4BindHandler(
				cfg,
				v4Logger,
				streamHandler,
				bindManager,
				addressUtils,
				sender,
				whitelist,
				blacklist,
				errorHandler,
			)
		},
	}

	v4HandlerDef := di.Def{
		Name:  "v4_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("v4_parser").(v4.Parser)
			cfg := ctn.Get("v4_config").(config.SocksV4Config)
			v4Logger := ctn.Get("v4_logger").(logger.SocksV4Logger)
			connectHandler := ctn.Get("v4_connect_handler").(server.V4ConnectHandler)
			bindHandler := ctn.Get("v4_bind_handler").(server.V4BindHandler)
			sender := ctn.Get("v4_sender").(server.V4Sender)
			errorHandler := ctn.Get("v4_error_handler").(server.V4ErrorHandler)

			return server.NewBaseV4Handler(
				parser,
				cfg,
				v4Logger,
				connectHandler,
				bindHandler,
				sender,
				errorHandler,
			)
		},
	}

	v4SenderDef := di.Def{
		Name:  "v4_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			protocol := ctn.Get("v4").(v4.Protocol)
			tcpConfig := ctn.Get("tcp_config").(config.TcpConfig)

			return server.NewBaseV4Sender(
				protocol,
				tcpConfig,
			)
		},
	}

	v4ErrorHandlerDef := di.Def{
		Name:  "v4_error_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			sender := ctn.Get("v4_sender").(server.V4Sender)
			v4Logger := ctn.Get("v4_logger").(logger.SocksV4Logger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return server.NewBaseV4ErrorHandler(
				v4Logger,
				sender,
				errorUtils,
			)
		},
	}

	v4aConnectHandlerDef := di.Def{
		Name:  "v4a_connect_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_config").(config.SocksV4aConfig)
			v4aLogger := ctn.Get("v4a_logger").(logger.SocksV4aLogger)
			streamHandler := ctn.Get("stream_handler").(server.StreamHandler)
			sender := ctn.Get("v4a_sender").(server.V4aSender)
			errorHandler := ctn.Get("v4a_error_handler").(server.V4aErrorHandler)
			whitelist := ctn.Get("whitelist_manager").(server.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(server.BlacklistManager)

			return server.NewBaseV4aConnectHandler(
				cfg,
				streamHandler,
				v4aLogger,
				sender,
				errorHandler,
				whitelist,
				blacklist,
			)
		},
	}

	v4aBindHandlerDef := di.Def{
		Name:  "v4a_bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_config").(config.SocksV4aConfig)
			v4aLogger := ctn.Get("v4a_logger").(logger.SocksV4aLogger)
			streamHandler := ctn.Get("stream_handler").(server.StreamHandler)
			bindManager := ctn.Get("bind_manager").(server.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v4a_sender").(server.V4aSender)
			whitelist := ctn.Get("whitelist_manager").(server.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(server.BlacklistManager)
			errorHandler := ctn.Get("v4a_error_handler").(server.V4aErrorHandler)

			return server.NewBaseV4aBindHandler(
				cfg,
				v4aLogger,
				streamHandler,
				bindManager,
				addressUtils,
				sender,
				whitelist,
				blacklist,
				errorHandler,
			)
		},
	}

	v4aHandlerDef := di.Def{
		Name:  "v4a_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("v4a_parser").(v4a.Parser)
			cfg := ctn.Get("v4a_config").(config.SocksV4aConfig)
			v4aLogger := ctn.Get("v4a_logger").(logger.SocksV4aLogger)
			connectHandler := ctn.Get("v4a_connect_handler").(server.V4aConnectHandler)
			bindHandler := ctn.Get("v4a_bind_handler").(server.V4aBindHandler)
			sender := ctn.Get("v4a_sender").(server.V4aSender)
			errorHandler := ctn.Get("v4a_error_handler").(server.V4aErrorHandler)

			return server.NewBaseV4aHandler(
				parser,
				cfg,
				v4aLogger,
				connectHandler,
				bindHandler,
				sender,
				errorHandler,
			)
		},
	}

	v4aSenderDef := di.Def{
		Name:  "v4a_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			protocol := ctn.Get("v4a").(v4a.Protocol)
			tcpConfig := ctn.Get("tcp_config").(config.TcpConfig)

			return server.NewBaseV4aSender(
				protocol,
				tcpConfig,
			)
		},
	}

	v4aErrorHandlerDef := di.Def{
		Name:  "v4a_error_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			sender := ctn.Get("v4a_sender").(server.V4aSender)
			v4Logger := ctn.Get("v4a_logger").(logger.SocksV4aLogger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return server.NewBaseV4aErrorHandler(
				v4Logger,
				sender,
				errorUtils,
			)
		},
	}
	v5ConnectHandler := di.Def{
		Name:  "v5_connect_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
			v5Logger := ctn.Get("v5_logger").(logger.SocksV5Logger)
			streamHandler := ctn.Get("stream_handler").(server.StreamHandler)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v5_sender").(server.V5Sender)
			errorHandler := ctn.Get("v5_error_handler").(server.V5ErrorHandler)
			whitelist := ctn.Get("whitelist_manager").(server.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(server.BlacklistManager)

			return server.NewBaseV5ConnectHandler(
				cfg,
				streamHandler,
				v5Logger,
				addressUtils,
				sender,
				errorHandler,
				whitelist,
				blacklist,
			)
		},
	}

	v5BindHandler := di.Def{
		Name:  "v5_bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
			v5Logger := ctn.Get("v5_logger").(logger.SocksV5Logger)
			streamHandler := ctn.Get("stream_handler").(server.StreamHandler)
			bindManager := ctn.Get("bind_manager").(server.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v5_sender").(server.V5Sender)
			whitelist := ctn.Get("whitelist_manager").(server.WhitelistManager)
			blacklist := ctn.Get("blacklist_manager").(server.BlacklistManager)
			errorHandler := ctn.Get("v5_error_handler").(server.V5ErrorHandler)

			return server.NewBaseV5BindHandler(
				bindManager,
				cfg,
				streamHandler,
				addressUtils,
				v5Logger,
				sender,
				whitelist,
				blacklist,
				errorHandler,
			)
		},
	}

	v5UdpAssociationHandler := di.Def{
		Name:  "v5_udp_association_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
			v5Logger := ctn.Get("v5_logger").(logger.SocksV5Logger)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			udpAssociationManager := ctn.Get("udp_association_manager").(server.UdpAssociationManager)
			sender := ctn.Get("v5_sender").(server.V5Sender)
			errorHandler := ctn.Get("v5_error_handler").(server.V5ErrorHandler)

			return server.NewBaseV5UdpAssociationHandler(
				cfg,
				addressUtils,
				udpAssociationManager,
				v5Logger,
				sender,
				errorHandler,
			)
		},
	}

	v5HandlerDef := di.Def{
		Name:  "v5_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			protocol := ctn.Get("v5").(v5.Protocol)
			parser := ctn.Get("v5_parser").(v5.Parser)
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
			authenticationHandler := ctn.Get("authentication_handler").(server.V5AuthenticationHandler)
			v5Logger := ctn.Get("v5_logger").(logger.SocksV5Logger)
			connectHandler := ctn.Get("v5_connect_handler").(server.V5ConnectHandler)
			bindHandler := ctn.Get("v5_bind_handler").(server.V5BindHandler)
			associationHandler := ctn.Get("v5_udp_association_handler").(server.V5UdpAssociationHandler)
			sender := ctn.Get("v5_sender").(server.V5Sender)
			errorHandler := ctn.Get("v5_error_handler").(server.V5ErrorHandler)

			return server.NewBaseV5Handler(
				protocol,
				parser,
				cfg,
				authenticationHandler,
				v5Logger,
				connectHandler,
				bindHandler,
				associationHandler,
				sender,
				errorHandler,
			)
		},
	}

	v5SenderDef := di.Def{
		Name:  "v5_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			protocol := ctn.Get("v5").(v5.Protocol)
			tcpConfig := ctn.Get("tcp_config").(config.TcpConfig)
			udpConfig := ctn.Get("udp_config").(config.UdpConfig)

			return server.NewBaseV5Sender(
				protocol,
				tcpConfig,
				udpConfig,
			)
		},
	}

	v5ErrorHandlerDef := di.Def{
		Name:  "v5_error_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			sender := ctn.Get("v5_sender").(server.V5Sender)
			v5Logger := ctn.Get("v5_logger").(logger.SocksV5Logger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return server.NewBaseV5ErrorHandler(
				v5Logger,
				sender,
				errorUtils,
			)
		},
	}

	udpAssociationManagerDef := di.Def{
		Name:  "udp_association_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return server.NewUdpAssociationManager(), nil
		},
	}

	serverDef := di.Def{
		Name:  "server",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connectionHandler := ctn.Get("connection_handler").(server.ConnectionHandler)
			packetHandler := ctn.Get("packet_handler").(server.PacketHandler)
			tcpConfig := ctn.Get("tcp_config").(config.TcpConfig)
			tcpLogger := ctn.Get("tcp_logger").(logger.TcpLogger)
			udpConfig := ctn.Get("udp_config").(config.UdpConfig)
			udpLogger := ctn.Get("udp_logger").(logger.UdpLogger)

			return server.NewServer(
				connectionHandler,
				packetHandler,
				tcpLogger,
				tcpConfig,
				udpLogger,
				udpConfig,
			)
		},
	}

	whitelistManagerDef := di.Def{
		Name:  "whitelist_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("whitelist_config").(config.WhitelistConfig)
			serverLogger := ctn.Get("server_logger").(logger.ServerLogger)

			return server.NewBaseWhitelistManager(cfg, serverLogger)
		},
	}

	blacklistManagerDef := di.Def{
		Name:  "blacklist_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("blacklist_config").(config.BlacklistConfig)
			serverLogger := ctn.Get("server_logger").(logger.ServerLogger)

			return server.NewBaseBlacklistManager(cfg, serverLogger)
		},
	}

	err := builder.Add(
		authenticationHandlerDef,
		connectionHandlerDef,
		streamHandlerDef,
		bindManagerDef,
		v4BindHandlerDef,
		v4ConnectHandlerDef,
		v4HandlerDef,
		v4SenderDef,
		v4ErrorHandlerDef,
		v4aConnectHandlerDef,
		v4aBindHandlerDef,
		v4aHandlerDef,
		v4aSenderDef,
		v4aErrorHandlerDef,
		v5ConnectHandler,
		v5BindHandler,
		v5UdpAssociationHandler,
		v5HandlerDef,
		v5SenderDef,
		v5ErrorHandlerDef,
		packetHandlerDef,
		udpAssociationManagerDef,
		serverDef,
		whitelistManagerDef,
		blacklistManagerDef,
	)

	if err != nil {
		panic(err)
	}

	registerUtils(builder)
}

func registerUtils(builder di.Builder) {
	addressUtilsDef := di.Def{
		Name:  "address_utils",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return utils.NewUtils()
		},
	}

	errorUtils := di.Def{
		Name:  "error_utils",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return utils.NewErrorUtils()
		},
	}

	err := builder.Add(
		addressUtilsDef,
		errorUtils,
	)

	if err != nil {
		panic(err)
	}

	start(builder)
}

func start(builder di.Builder) {
	ctn := builder.Build()

	serv := ctn.Get("server").(server.Server)

	serv.Start()
}
