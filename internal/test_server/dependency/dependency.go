package dependency

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"github.com/sarulabs/di"
	"os"
	_config "socks/internal/test_server/config"
	_handlers "socks/internal/test_server/handlers"
	_logger "socks/internal/test_server/logger"
	_picture "socks/internal/test_server/picture"
	_server "socks/internal/test_server/server"
)

func Register(builder di.Builder) {
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

			return _config.NewConfig(validate, configPath)
		},
	}

	zeroLoggerDef := di.Def{
		Name:  "zero_logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config").(_config.Config)

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

			return _logger.NewLogger(logger)
		},
	}

	connectionHandlerDef := di.Def{
		Name:  "connection_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config").(_config.Config)
			logger := ctn.Get("logger").(_logger.Logger)
			sender := ctn.Get("picture_sender").(_picture.Sender)

			return _handlers.NewConnectHandler(config, logger, sender)
		},
	}

	packetHandlerDef := di.Def{
		Name:  "packet_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config").(_config.Config)
			logger := ctn.Get("logger").(_logger.Logger)

			return _handlers.NewPacketHandler(config, logger)
		},
	}

	bindHandlerDef := di.Def{
		Name:  "bind_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config").(_config.Config)
			logger := ctn.Get("logger").(_logger.Logger)
			picture := ctn.Get("picture_sender").(_picture.Sender)

			return _handlers.NewBindHandler(config, logger, picture)
		},
	}

	serverDef := di.Def{
		Name:  "server",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config").(_config.Config)
			logger := ctn.Get("logger").(_logger.Logger)
			connectionHandler := ctn.Get("connection_handler").(_handlers.ConnectHandler)
			packetHandler := ctn.Get("packet_handler").(_handlers.PacketHandler)
			bindHandler := ctn.Get("bind_handler").(_handlers.BindHandler)

			return _server.NewServer(config, connectionHandler, packetHandler, logger, bindHandler)
		},
	}

	pictureSenderDef := di.Def{
		Name:  "picture_sender",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config").(_config.Config)
			logger := ctn.Get("logger").(_logger.Logger)

			return _picture.NewSender(logger, config)
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
		bindHandlerDef,
		serverDef,
		pictureSenderDef,
	)

	if err != nil {
		panic(err)
	}
}
