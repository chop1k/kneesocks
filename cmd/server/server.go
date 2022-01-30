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
	"socks/handlers"
	v42 "socks/handlers/v4"
	v4a2 "socks/handlers/v4a"
	v52 "socks/handlers/v5"
	"socks/handlers/v5/authenticator"
	"socks/logger"
	"socks/managers"
	"socks/protocol/auth/password"
	v4 "socks/protocol/v4"
	"socks/protocol/v4a"
	v5 "socks/protocol/v5"
	"socks/server"
	"socks/transfer"
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
			return password.NewBaseBuilder()
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
			return v5.NewBaseBuilder()
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
	passwordAuthenticatorDef := di.Def{
		Name:  "password_authenticator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			passwd := ctn.Get("auth_password").(password.Password)
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)

			return authenticator.NewBasePasswordAuthenticator(
				passwd,
				cfg,
				errorHandler,
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

	authenticationHandlerDef := di.Def{
		Name:  "authentication_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			protocol := ctn.Get("v5").(v5.Protocol)
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
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
			tcpLogger := ctn.Get("tcp_logger").(logger.TcpLogger)
			tcpConfig := ctn.Get("tcp_config").(config.TcpConfig)

			return handlers.NewBaseConnectionHandler(
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

	streamHandlerDef := di.Def{
		Name:  "stream_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return transfer.NewBaseStreamHandler(), nil
		},
	}

	bindManagerDef := di.Def{
		Name:  "bind_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return managers.NewBindManager(), nil
		},
	}

	v4WhitelistDef := di.Def{
		Name:  "v4_whitelist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_config").(config.SocksV4Config)
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)

			return v42.NewBaseWhitelist(cfg, whitelist)
		},
	}

	v4BlacklistDef := di.Def{
		Name:  "v4_blacklist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_config").(config.SocksV4Config)
			whitelist := ctn.Get("blacklist_manager").(managers.BlacklistManager)

			return v42.NewBaseBlacklist(cfg, whitelist)
		},
	}

	v4ConnectHandlerDef := di.Def{
		Name:  "v4_connect_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_config").(config.SocksV4Config)
			v4Logger := ctn.Get("v4_logger").(logger.SocksV4Logger)
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			sender := ctn.Get("v4_sender").(v42.Sender)
			errorHandler := ctn.Get("v4_error_handler").(v42.ErrorHandler)
			whitelist := ctn.Get("v4_whitelist").(v42.Whitelist)
			blacklist := ctn.Get("v4_blacklist").(v42.Blacklist)

			return v42.NewBaseConnectHandler(
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
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v4_sender").(v42.Sender)
			whitelist := ctn.Get("v4_whitelist").(v42.Whitelist)
			blacklist := ctn.Get("v4_blacklist").(v42.Blacklist)
			errorHandler := ctn.Get("v4_error_handler").(v42.ErrorHandler)

			return v42.NewBaseBindHandler(
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
			connectHandler := ctn.Get("v4_connect_handler").(v42.ConnectHandler)
			bindHandler := ctn.Get("v4_bind_handler").(v42.BindHandler)
			sender := ctn.Get("v4_sender").(v42.Sender)
			errorHandler := ctn.Get("v4_error_handler").(v42.ErrorHandler)

			return v42.NewBaseHandler(
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

			return v42.NewBaseSender(
				protocol,
				tcpConfig,
			)
		},
	}

	v4ErrorHandlerDef := di.Def{
		Name:  "v4_error_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			sender := ctn.Get("v4_sender").(v42.Sender)
			v4Logger := ctn.Get("v4_logger").(logger.SocksV4Logger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return v42.NewBaseErrorHandler(
				v4Logger,
				sender,
				errorUtils,
			)
		},
	}

	v4aWhitelistDef := di.Def{
		Name:  "v4a_whitelist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_config").(config.SocksV4aConfig)
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)

			return v4a2.NewBaseWhitelist(cfg, whitelist)
		},
	}

	v4aBlacklistDef := di.Def{
		Name:  "v4a_blacklist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_config").(config.SocksV4aConfig)
			whitelist := ctn.Get("blacklist_manager").(managers.BlacklistManager)

			return v4a2.NewBaseBlacklist(cfg, whitelist)
		},
	}

	v4aConnectHandlerDef := di.Def{
		Name:  "v4a_connect_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_config").(config.SocksV4aConfig)
			v4aLogger := ctn.Get("v4a_logger").(logger.SocksV4aLogger)
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			sender := ctn.Get("v4a_sender").(v4a2.Sender)
			errorHandler := ctn.Get("v4a_error_handler").(v4a2.ErrorHandler)
			whitelist := ctn.Get("v4a_whitelist").(v4a2.Whitelist)
			blacklist := ctn.Get("v4a_blacklist").(v4a2.Blacklist)

			return v4a2.NewBaseConnectHandler(
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
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v4a_sender").(v4a2.Sender)
			whitelist := ctn.Get("v4a_whitelist").(v4a2.Whitelist)
			blacklist := ctn.Get("v4a_blacklist").(v4a2.Blacklist)
			errorHandler := ctn.Get("v4a_error_handler").(v4a2.ErrorHandler)

			return v4a2.NewBaseBindHandler(
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
			connectHandler := ctn.Get("v4a_connect_handler").(v4a2.ConnectHandler)
			bindHandler := ctn.Get("v4a_bind_handler").(v4a2.BindHandler)
			sender := ctn.Get("v4a_sender").(v4a2.Sender)
			errorHandler := ctn.Get("v4a_error_handler").(v4a2.ErrorHandler)

			return v4a2.NewBaseHandler(
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

			return v4a2.NewBaseSender(
				protocol,
				tcpConfig,
			)
		},
	}

	v4aErrorHandlerDef := di.Def{
		Name:  "v4a_error_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			sender := ctn.Get("v4a_sender").(v4a2.Sender)
			v4Logger := ctn.Get("v4a_logger").(logger.SocksV4aLogger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return v4a2.NewBaseErrorHandler(
				v4Logger,
				sender,
				errorUtils,
			)
		},
	}

	v5WhitelistDef := di.Def{
		Name:  "v5_whitelist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
			whitelist := ctn.Get("whitelist_manager").(managers.WhitelistManager)

			return v52.NewBaseWhitelist(cfg, whitelist)
		},
	}

	v5BlacklistDef := di.Def{
		Name:  "v5_blacklist",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
			whitelist := ctn.Get("blacklist_manager").(managers.BlacklistManager)

			return v52.NewBaseBlacklist(cfg, whitelist)
		},
	}

	v5ConnectHandler := di.Def{
		Name:  "v5_connect_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
			v5Logger := ctn.Get("v5_logger").(logger.SocksV5Logger)
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v5_sender").(v52.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)
			whitelist := ctn.Get("v5_whitelist").(v52.Whitelist)
			blacklist := ctn.Get("v5_blacklist").(v52.Blacklist)

			return v52.NewBaseConnectHandler(
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
			streamHandler := ctn.Get("stream_handler").(transfer.StreamHandler)
			bindManager := ctn.Get("bind_manager").(managers.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			sender := ctn.Get("v5_sender").(v52.Sender)
			whitelist := ctn.Get("v5_whitelist").(v52.Whitelist)
			blacklist := ctn.Get("v5_blacklist").(v52.Blacklist)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)

			return v52.NewBaseBindHandler(
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
			udpAssociationManager := ctn.Get("udp_association_manager").(managers.UdpAssociationManager)
			sender := ctn.Get("v5_sender").(v52.Sender)
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

	v5HandlerDef := di.Def{
		Name:  "v5_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			protocol := ctn.Get("v5").(v5.Protocol)
			parser := ctn.Get("v5_parser").(v5.Parser)
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
			authenticationHandler := ctn.Get("authentication_handler").(v52.AuthenticationHandler)
			v5Logger := ctn.Get("v5_logger").(logger.SocksV5Logger)
			connectHandler := ctn.Get("v5_connect_handler").(v52.ConnectHandler)
			bindHandler := ctn.Get("v5_bind_handler").(v52.BindHandler)
			associationHandler := ctn.Get("v5_udp_association_handler").(v52.UdpAssociationHandler)
			sender := ctn.Get("v5_sender").(v52.Sender)
			errorHandler := ctn.Get("v5_error_handler").(v52.ErrorHandler)

			return v52.NewBaseHandler(
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

			return v52.NewBaseSender(
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
			sender := ctn.Get("v5_sender").(v52.Sender)
			v5Logger := ctn.Get("v5_logger").(logger.SocksV5Logger)
			errorUtils := ctn.Get("error_utils").(utils.ErrorUtils)

			return v52.NewBaseErrorHandler(
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
			return managers.NewUdpAssociationManager(), nil
		},
	}

	serverDef := di.Def{
		Name:  "server",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			connectionHandler := ctn.Get("connection_handler").(handlers.ConnectionHandler)
			packetHandler := ctn.Get("packet_handler").(handlers.PacketHandler)
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

			return managers.NewBaseWhitelistManager(cfg, serverLogger)
		},
	}

	blacklistManagerDef := di.Def{
		Name:  "blacklist_manager",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("blacklist_config").(config.BlacklistConfig)
			serverLogger := ctn.Get("server_logger").(logger.ServerLogger)

			return managers.NewBaseBlacklistManager(cfg, serverLogger)
		},
	}

	err := builder.Add(
		passwordAuthenticatorDef,
		noAuthAuthenticatorDef,
		authenticationHandlerDef,
		connectionHandlerDef,
		streamHandlerDef,
		bindManagerDef,
		v4WhitelistDef,
		v4BlacklistDef,
		v4BindHandlerDef,
		v4ConnectHandlerDef,
		v4HandlerDef,
		v4SenderDef,
		v4ErrorHandlerDef,
		v4aWhitelistDef,
		v4aBlacklistDef,
		v4aConnectHandlerDef,
		v4aBindHandlerDef,
		v4aHandlerDef,
		v4aSenderDef,
		v4aErrorHandlerDef,
		v5WhitelistDef,
		v5BlacklistDef,
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
