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
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseConfig(cfgTree), nil
		},
	}

	tcpConfigDef := di.Def{
		Name:  "tcp_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseTcpConfig(cfgTree), nil
		},
	}

	udpConfigDef := di.Def{
		Name:  "udp_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseUdpConfig(cfgTree), nil
		},
	}

	v4ConfigDef := di.Def{
		Name:  "v4_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV4Config(cfgTree), nil
		},
	}

	v4aConfigDef := di.Def{
		Name:  "v4a_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV4aConfig(cfgTree), nil
		},
	}

	v5ConfigDef := di.Def{
		Name:  "v5_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV5Config(cfgTree), nil
		},
	}

	tcpLoggerConfigDef := di.Def{
		Name:  "tcp_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseTcpLoggerConfig(cfgTree)
		},
	}

	udpLoggerConfigDef := di.Def{
		Name:  "udp_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseUdpLoggerConfig(cfgTree)
		},
	}

	v4LoggerConfigDef := di.Def{
		Name:  "v4_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV4LoggerConfig(cfgTree)
		},
	}

	v4aLoggerConfigDef := di.Def{
		Name:  "v4a_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV4aLoggerConfig(cfgTree)
		},
	}

	v5LoggerConfigDef := di.Def{
		Name:  "v5_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseSocksV5LoggerConfig(cfgTree)
		},
	}

	unixLoggerConfigDef := di.Def{
		Name:  "unix_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			panic("should not be used. ")
		},
	}

	errorsLoggerConfigDef := di.Def{
		Name:  "errors_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			panic("should not be used. ")
		},
	}

	whitelistConfigDef := di.Def{
		Name:  "whitelist_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseWhitelistConfig(cfgTree), nil
		},
	}

	blacklistConfigDef := di.Def{
		Name:  "blacklist_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfgTree := ctn.Get("config_tree").(tree.Config)

			return config.NewBaseBlacklistConfig(cfgTree), nil
		},
	}

	unixConfigDef := di.Def{
		Name:  "unix_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			panic("should not be used. ")
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
		unixLoggerConfigDef,
		errorsLoggerConfigDef,
		whitelistConfigDef,
		blacklistConfigDef,
		unixConfigDef,
	)

	if err != nil {
		panic(err)
	}

	registerZeroLog(builder)
}

func registerZeroLog(builder di.Builder) {
	buildLogger := func(level int, loggers ...io.Writer) (zerolog.Logger, error) {
		return zerolog.New(zerolog.MultiLevelWriter(loggers...)).
			With().
			Timestamp().
			Logger().
			Level(zerolog.Level(level)), nil
	}

	tcpZeroLoggerDef := di.Def{
		Name:  "tcp_zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("tcp_logger_config").(config.TcpLoggerConfig)

			level, err := cfg.GetLevel()

			var loggers []io.Writer

			if err != nil {
				return buildLogger(126, loggers...)
			}

			if output, err := cfg.GetConsoleOutput(); err == nil {
				loggers = append(loggers, zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: output.TimeFormat,
				})
			} else {
				if err == config.TcpLoggerDisabledError {
					return buildLogger(126, loggers...)
				}
			}

			if output, err := cfg.GetFileOutput(); err == nil {
				file, err := os.Open(output.Path)

				if err != nil {
					return nil, err
				}

				loggers = append(loggers, zerolog.New(file).Level(zerolog.Level(level)))
			} else {
				if err == config.TcpLoggerDisabledError {
					return buildLogger(126, loggers...)
				}
			}

			return buildLogger(level, loggers...)
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
				return buildLogger(126, loggers...)
			}

			if output, err := cfg.GetConsoleOutput(); err == nil {
				loggers = append(loggers, zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: output.TimeFormat,
				})
			} else {
				if err == config.SocksV4LoggerDisabledError {
					return buildLogger(126, loggers...)
				}
			}

			if output, err := cfg.GetFileOutput(); err == nil {
				file, err := os.Open(output.Path)

				if err != nil {
					return nil, err
				}

				loggers = append(loggers, zerolog.New(file).Level(zerolog.Level(level)))
			} else {
				if err == config.SocksV4LoggerDisabledError {
					return buildLogger(126, loggers...)
				}
			}

			return buildLogger(level, loggers...)
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
				return buildLogger(126, loggers...)
			}

			if output, err := cfg.GetConsoleOutput(); err == nil {
				loggers = append(loggers, zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: output.TimeFormat,
				})
			} else {
				if err == config.SocksV4aLoggerDisabledError {
					return buildLogger(126, loggers...)
				}
			}

			if output, err := cfg.GetFileOutput(); err == nil {
				file, err := os.Open(output.Path)

				if err != nil {
					return nil, err
				}

				loggers = append(loggers, zerolog.New(file).Level(zerolog.Level(level)))
			} else {
				if err == config.SocksV4aLoggerDisabledError {
					return buildLogger(126, loggers...)
				}
			}

			return buildLogger(level, loggers...)
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
				return buildLogger(126, loggers...)
			}

			if output, err := cfg.GetConsoleOutput(); err == nil {
				loggers = append(loggers, zerolog.ConsoleWriter{
					Out:        os.Stdout,
					TimeFormat: output.TimeFormat,
				})
			} else {
				if err == config.SocksV5LoggerDisabledError {
					return buildLogger(126, loggers...)
				}
			}

			if output, err := cfg.GetFileOutput(); err == nil {
				file, err := os.Open(output.Path)

				if err != nil {
					return nil, err
				}

				loggers = append(loggers, zerolog.New(file).Level(zerolog.Level(level)))
			} else {
				if err == config.SocksV5LoggerDisabledError {
					return buildLogger(126, loggers...)
				}
			}

			return buildLogger(level, loggers...)
		},
	}

	err := builder.Add(
		tcpZeroLoggerDef,
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
	//errorsLoggerDef := di.Def{}

	tcpLoggerDef := di.Def{
		Name:  "tcp_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			zero := ctn.Get("tcp_zero_logger").(zerolog.Logger)

			return logger.NewBaseTcpLogger(zero)
		},
	}

	//udpLoggerDef := di.Def{}

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
		//errorsLoggerDef,
		tcpLoggerDef,
		//udpLoggerDef,
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

			return v5.NewProtocol(builder), nil
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
			cfg := ctn.Get("config_tree").(tree.Config)
			passwd := ctn.Get("auth_password").(password.Password)
			protocol := ctn.Get("v5").(v5.Protocol)

			return server.NewBaseAuthenticationHandler(cfg, passwd, protocol), nil
		},
	}

	connectionHandlerDef := di.Def{
		Name:  "connection_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			authenticationHandler := ctn.Get("authentication_handler").(server.AuthenticationHandler)
			streamHandler := ctn.Get("stream_handler").(server.StreamHandler)
			v4Handler := ctn.Get("v4_handler").(server.V4Handler)
			v4aHandler := ctn.Get("v4a_handler").(server.V4aHandler)
			v5Handler := ctn.Get("v5_handler").(server.V5Handler)
			bindManager := ctn.Get("bind_manager").(server.BindManager)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			tcpLogger := ctn.Get("tcp_logger").(logger.TcpLogger)
			tcpConfig := ctn.Get("tcp_config").(config.TcpConfig)

			return server.NewBaseConnectionHandler(
				authenticationHandler,
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

	v4HandlerDef := di.Def{
		Name:  "v4_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			streamHandler := ctn.Get("stream_handler").(server.StreamHandler)
			bindManager := ctn.Get("bind_manager").(server.BindManager)
			protocol := ctn.Get("v4").(v4.Protocol)
			parser := ctn.Get("v4_parser").(v4.Parser)
			cfg := ctn.Get("v4_config").(config.SocksV4Config)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			v4Logger := ctn.Get("v4_logger").(logger.SocksV4Logger)
			tcpConfig := ctn.Get("tcp_config").(config.TcpConfig)

			return server.NewBaseV4Handler(
				protocol,
				parser,
				bindManager,
				cfg,
				tcpConfig,
				streamHandler,
				addressUtils,
				v4Logger,
			)
		},
	}

	v4aHandlerDef := di.Def{
		Name:  "v4a_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			streamHandler := ctn.Get("stream_handler").(server.StreamHandler)
			bindManager := ctn.Get("bind_manager").(server.BindManager)
			protocol := ctn.Get("v4a").(v4a.Protocol)
			parser := ctn.Get("v4a_parser").(v4a.Parser)
			cfg := ctn.Get("v4a_config").(config.SocksV4aConfig)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			v4aLogger := ctn.Get("v4a_logger").(logger.SocksV4aLogger)
			tcpConfig := ctn.Get("tcp_config").(config.TcpConfig)

			return server.NewBaseV4aHandler(
				protocol,
				parser,
				bindManager,
				cfg,
				streamHandler,
				addressUtils,
				v4aLogger,
				tcpConfig,
			)
		},
	}

	v5HandlerDef := di.Def{
		Name:  "v5_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			streamHandler := ctn.Get("stream_handler").(server.StreamHandler)
			bindManager := ctn.Get("bind_manager").(server.BindManager)
			protocol := ctn.Get("v5").(v5.Protocol)
			parser := ctn.Get("v5_parser").(v5.Parser)
			cfg := ctn.Get("v5_config").(config.SocksV5Config)
			addressUtils := ctn.Get("address_utils").(utils.AddressUtils)
			udpAssociationManager := ctn.Get("udp_association_manager").(server.UdpAssociationManager)
			authenticationHandler := ctn.Get("authentication_handler").(server.AuthenticationHandler)
			v5Logger := ctn.Get("v5_logger").(logger.SocksV5Logger)
			tcpConfig := ctn.Get("tcp_config").(config.TcpConfig)
			udpConfig := ctn.Get("udp_config").(config.UdpConfig)

			return server.NewBaseV5Handler(
				protocol,
				parser,
				bindManager,
				cfg,
				streamHandler,
				addressUtils,
				udpAssociationManager,
				authenticationHandler,
				v5Logger,
				tcpConfig,
				udpConfig,
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
			cfg := ctn.Get("config_tree").(tree.Config)
			connectionHandler := ctn.Get("connection_handler").(server.ConnectionHandler)
			packetHandler := ctn.Get("packet_handler").(server.PacketHandler)
			tcpLogger := ctn.Get("tcp_logger").(logger.TcpLogger)

			return server.NewServer(cfg, connectionHandler, packetHandler, tcpLogger), nil
		},
	}

	err := builder.Add(
		authenticationHandlerDef,
		connectionHandlerDef,
		streamHandlerDef,
		bindManagerDef,
		v4HandlerDef,
		v4aHandlerDef,
		v5HandlerDef,
		packetHandlerDef,
		udpAssociationManagerDef,
		serverDef,
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
			return utils.NewUtils(), nil
		},
	}

	err := builder.Add(
		addressUtilsDef,
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
