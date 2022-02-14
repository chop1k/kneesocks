package dependency

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/sarulabs/di"
	"os"
	"socks/internal/kneesocks/config/tcp"
	"socks/internal/kneesocks/config/tree"
	"socks/internal/kneesocks/config/udp"
	"socks/internal/kneesocks/config/v4"
	"socks/internal/kneesocks/config/v4a"
	"socks/internal/kneesocks/config/v5"
)

func registerConfig(builder di.Builder) {
	registerTree(builder)
	registerTcpConfig(builder)
	registerUdpConfig(builder)
	registerV4Config(builder)
	registerV4aConfig(builder)
	registerV5Config(builder)
}

func registerTree(builder di.Builder) {
	validatorDef := di.Def{
		Name:  "validator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return *validator.New(), nil
		},
	}

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

	configTreeDef := di.Def{
		Name:  "config_tree",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			validate := ctn.Get("validator").(validator.Validate)
			configPath := ctn.Get("config_path").(string)

			return tree.NewConfig(validate, configPath)
		},
	}

	err := builder.Add(
		validatorDef,
		configPathDef,
		configTreeDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerTcpConfig(builder di.Builder) {
	bindConfigDef := di.Def{
		Name:  "tcp_base_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config_tree").(tree.Config)
			handler := ctn.Get("tcp_config_handler").(tcp.Handler)

			return handler.Handle(config.Tcp), nil
		},
	}

	handlerDef := di.Def{
		Name:  "tcp_config_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return tcp.NewHandler()
		},
	}

	loggerConfigDef := di.Def{
		Name:  "tcp_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config_tree").(tree.Config).Log

			return tcp.NewLoggerConfig(config)
		},
	}

	replicatorDef := di.Def{
		Name:  "tcp_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			base := ctn.Get("tcp_base_config").(tcp.Config)

			return tcp.NewConfigReplicator(base.Deadline)
		},
	}

	err := builder.Add(
		bindConfigDef,
		handlerDef,
		loggerConfigDef,
		replicatorDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerUdpConfig(builder di.Builder) {
	bindConfigDef := di.Def{
		Name:  "udp_base_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config_tree").(tree.Config)
			handler := ctn.Get("udp_config_handler").(udp.Handler)

			return handler.Handle(config.Udp), nil
		},
	}

	handlerDef := di.Def{
		Name:  "udp_config_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return udp.NewHandler()
		},
	}

	loggerConfigDef := di.Def{
		Name:  "udp_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config_tree").(tree.Config).Log

			return udp.NewLoggerConfig(config)
		},
	}

	replicatorDef := di.Def{
		Name:  "udp_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			base := ctn.Get("udp_base_config").(udp.Config)

			return udp.NewConfigReplicator(base.Buffer, base.Deadline)
		},
	}

	err := builder.Add(
		bindConfigDef,
		handlerDef,
		loggerConfigDef,
		replicatorDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV4Config(builder di.Builder) {
	configDef := di.Def{
		Name:  "v4_base_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config_tree").(tree.Config)
			handler := ctn.Get("v4_config_handler").(v4.Handler)

			return handler.Handle(config.SocksV4)
		},
	}

	handlerDef := di.Def{
		Name:  "v4_config_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4.NewHandler()
		},
	}

	loggerConfigDef := di.Def{
		Name:  "v4_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config_tree").(tree.Config).Log

			return v4.NewLoggerConfig(config)
		},
	}

	replicatorDef := di.Def{
		Name:  "v4_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("v4_base_config").(*v4.Config)

			return v4.NewConfigReplicator(config)
		},
	}

	err := builder.Add(
		configDef,
		handlerDef,
		loggerConfigDef,
		replicatorDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV4aConfig(builder di.Builder) {
	configDef := di.Def{
		Name:  "v4a_base_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config_tree").(tree.Config)
			handler := ctn.Get("v4a_config_handler").(v4a.Handler)

			return handler.Handle(config.SocksV4a)
		},
	}

	handlerDef := di.Def{
		Name:  "v4a_config_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4a.NewHandler()
		},
	}

	loggerConfigDef := di.Def{
		Name:  "v4a_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config_tree").(tree.Config).Log

			return v4a.NewLoggerConfig(config)
		},
	}

	replicatorDef := di.Def{
		Name:  "v4a_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("v4a_base_config").(*v4a.Config)

			return v4a.NewConfigReplicator(config)
		},
	}

	err := builder.Add(
		configDef,
		handlerDef,
		loggerConfigDef,
		replicatorDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV5Config(builder di.Builder) {
	configDef := di.Def{
		Name:  "v5_base_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config_tree").(tree.Config)
			handler := ctn.Get("v5_config_handler").(v5.Handler)

			return handler.Handle(config.SocksV5)
		},
	}

	handlerDef := di.Def{
		Name:  "v5_config_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v5.NewHandler()
		},
	}

	loggerConfigDef := di.Def{
		Name:  "v5_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("config_tree").(tree.Config).Log

			return v5.NewLoggerConfig(config)
		},
	}

	replicatorDef := di.Def{
		Name:  "v5_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get("v5_base_config").(*v5.Config)

			return v5.NewConfigReplicator(config)
		},
	}

	err := builder.Add(
		configDef,
		handlerDef,
		loggerConfigDef,
		replicatorDef,
	)

	if err != nil {
		panic(err)
	}
}
