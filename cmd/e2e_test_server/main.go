package main

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/sarulabs/di"
	"os"
	"socks/cmd/e2e_test_server/protocol"
)

func main() {
	builder, err := di.NewBuilder()

	if err != nil {
		panic(err)
	}

	register(*builder)
}

func register(builder di.Builder) {
	configPathDef := di.Def{
		Name:  "config_path",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path, ok := os.LookupEnv("config_path")

			if !ok {
				return nil, errors.New("Config path is not specified. ")
			}

			return path, nil
		},
	}

	validatorDef := di.Def{
		Name:  "validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return *validator.New(), nil
		},
	}
	configDef := di.Def{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			validate := ctn.Get("validator").(validator.Validate)
			configPath := ctn.Get("config_path").(string)

			return NewConfig(validate, configPath)
		},
	}

	zeroLoggerDef := di.Def{
		Name:  "zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config").(Config)

			consoleLogger := zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: "2006-01-02 15:04:05",
			}

			file, err := os.OpenFile(cfg.Log.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

			if err != nil {
				return nil, err
			}

			return zerolog.New(zerolog.MultiLevelWriter(consoleLogger, file)).
				With().
				Timestamp().
				Logger().
				Level(0), nil
		},
	}

	loggerDef := di.Def{
		Name:  "logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			logger := ctn.Get("zero_logger").(zerolog.Logger)

			return NewLogger(logger)
		},
	}

	connectionHandlerDef := di.Def{
		Name:  "connection_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config").(Config)
			logger := ctn.Get("logger").(Logger)
			sender := ctn.Get("picture_sender").(PictureSender)
			proto := ctn.Get("protocol").(protocol.Protocol)

			return NewConnectionHandler(cfg, logger, sender, proto)
		},
	}

	packetHandlerDef := di.Def{
		Name:  "packet_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config").(Config)
			logger := ctn.Get("logger").(Logger)

			return NewPacketHandler(cfg, logger)
		},
	}

	serverDef := di.Def{
		Name:  "server",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config").(Config)
			logger := ctn.Get("logger").(Logger)
			connectionHandler := ctn.Get("connection_handler").(ConnectionHandler)
			packetHandler := ctn.Get("packet_handler").(PacketHandler)

			return NewServer(cfg, connectionHandler, packetHandler, logger)
		},
	}

	builderDef := di.Def{
		Name:  "builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return protocol.NewBuilder()
		},
	}

	parserDef := di.Def{
		Name:  "parser",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return protocol.NewParser()
		},
	}

	protocolDef := di.Def{
		Name:  "protocol",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			parser := ctn.Get("parser").(protocol.Parser)
			builder := ctn.Get("builder").(protocol.Builder)

			return protocol.NewProtocol(parser, builder)
		},
	}

	pictureSenderDef := di.Def{
		Name:  "picture_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config").(Config)
			logger := ctn.Get("logger").(Logger)

			return NewPictureSender(logger, cfg)
		},
	}

	err := builder.Add(
		configPathDef,
		validatorDef,
		configDef,
		zeroLoggerDef,
		loggerDef,
		connectionHandlerDef,
		packetHandlerDef,
		serverDef,
		builderDef,
		parserDef,
		protocolDef,
		pictureSenderDef,
	)

	if err != nil {
		panic(err)
	}

	start(builder.Build())
}

func start(ctn di.Container) {
	server := ctn.Get("server").(Server)

	server.Start()
}
