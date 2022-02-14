package main

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/sarulabs/di"
	"os"
	"socks/config/tcp"
	"socks/config/tree"
	"socks/config/udp"
	v43 "socks/config/v4"
	v4a3 "socks/config/v4a"
	v53 "socks/config/v5"
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
			cfg := ctn.Get("config_tree").(tree.Config)
			handler := ctn.Get("tcp_config_handler").(tcp.Handler)

			return handler.Handle(cfg.Tcp), nil
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
			cfg := ctn.Get("config_tree").(tree.Config).Log

			return tcp.NewLoggerConfig(cfg)
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
			cfg := ctn.Get("config_tree").(tree.Config)
			handler := ctn.Get("udp_config_handler").(udp.Handler)

			return handler.Handle(cfg.Udp), nil
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
			cfg := ctn.Get("config_tree").(tree.Config).Log

			return udp.NewLoggerConfig(cfg)
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
			cfg := ctn.Get("config_tree").(tree.Config)
			handler := ctn.Get("v4_config_handler").(v43.Handler)

			return handler.Handle(cfg.SocksV4)
		},
	}

	handlerDef := di.Def{
		Name:  "v4_config_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v43.NewHandler()
		},
	}

	loggerConfigDef := di.Def{
		Name:  "v4_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config).Log

			return v43.NewLoggerConfig(cfg)
		},
	}

	replicatorDef := di.Def{
		Name:  "v4_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4_base_config").(*v43.Config)

			return v43.NewConfigReplicator(cfg)
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
			cfg := ctn.Get("config_tree").(tree.Config)
			handler := ctn.Get("v4a_config_handler").(v4a3.Handler)

			return handler.Handle(cfg.SocksV4a)
		},
	}

	handlerDef := di.Def{
		Name:  "v4a_config_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v4a3.NewHandler()
		},
	}

	loggerConfigDef := di.Def{
		Name:  "v4a_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config).Log

			return v4a3.NewLoggerConfig(cfg)
		},
	}

	replicatorDef := di.Def{
		Name:  "v4a_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v4a_base_config").(*v4a3.Config)

			return v4a3.NewConfigReplicator(cfg)
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
			cfg := ctn.Get("config_tree").(tree.Config)
			handler := ctn.Get("v5_config_handler").(v53.Handler)

			return handler.Handle(cfg.SocksV5)
		},
	}

	handlerDef := di.Def{
		Name:  "v5_config_handler",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return v53.NewHandler()
		},
	}

	loggerConfigDef := di.Def{
		Name:  "v5_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_tree").(tree.Config).Log

			return v53.NewLoggerConfig(cfg)
		},
	}

	replicatorDef := di.Def{
		Name:  "v5_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("v5_base_config").(*v53.Config)

			return v53.NewConfigReplicator(cfg)
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
