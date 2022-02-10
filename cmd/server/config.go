package main

import (
	"errors"
	"github.com/Jeffail/gabs"
	"github.com/go-playground/validator/v10"
	"github.com/sarulabs/di"
	"os"
	"socks/config"
	"socks/config/tcp"
	"socks/config/udp"
	v43 "socks/config/v4"
	v4a3 "socks/config/v4a"
	v53 "socks/config/v5"
)

func registerConfig(builder di.Builder) {
	configDef := di.Def{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return config.NewBaseConfig(cfg), nil
		},
	}

	serverLoggerConfigDef := di.Def{
		Name:  "server_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return config.NewBaseServerLoggerConfig(cfg)
		},
	}

	err := builder.Add(
		configDef,
		serverLoggerConfigDef,
	)

	if err != nil {
		panic(err)
	}

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
		Name:  "config_container",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			configPath := ctn.Get("config_path").(string)

			container, err := gabs.ParseJSONFile(configPath)

			return *container, err
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
	loggerConfigDef := di.Def{
		Name:  "tcp_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return tcp.NewBaseLoggerConfig(cfg)
		},
	}

	deadlineDef := di.Def{
		Name:  "tcp_deadline_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return tcp.NewBaseDeadlineConfig(cfg)
		},
	}

	configDef := di.Def{
		Name:  "tcp_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return tcp.NewBaseBindConfig(cfg)
		},
	}

	err := builder.Add(
		loggerConfigDef,
		deadlineDef,
		configDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerUdpConfig(builder di.Builder) {
	bufferConfigDef := di.Def{
		Name:  "udp_buffer_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return udp.NewBaseBufferConfig(cfg)
		},
	}

	loggerConfigDef := di.Def{
		Name:  "udp_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return udp.NewBaseLoggerConfig(cfg)
		},
	}

	configDef := di.Def{
		Name:  "udp_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return udp.NewBaseBindConfig(cfg)
		},
	}

	err := builder.Add(
		bufferConfigDef,
		loggerConfigDef,
		configDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV4Config(builder di.Builder) {
	loggerConfigDef := di.Def{
		Name:  "v4_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v43.NewBaseLoggerConfig(cfg)
		},
	}

	deadlineDef := di.Def{
		Name:  "v4_deadline_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v43.NewBaseDeadlineConfig(cfg)
		},
	}

	configDef := di.Def{
		Name:  "v4_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v43.NewBaseConfig(cfg)
		},
	}

	restrictionsConfigDef := di.Def{
		Name:  "v4_restrictions_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v43.NewBaseRestrictionsConfig(cfg)
		},
	}

	err := builder.Add(
		loggerConfigDef,
		deadlineDef,
		configDef,
		restrictionsConfigDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV4aConfig(builder di.Builder) {
	loggerConfigDef := di.Def{
		Name:  "v4a_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v4a3.NewBaseLoggerConfig(cfg)
		},
	}

	deadlineDef := di.Def{
		Name:  "v4a_deadline_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v4a3.NewBaseDeadlineConfig(cfg)
		},
	}

	configDef := di.Def{
		Name:  "v4a_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v4a3.NewBaseConfig(cfg)
		},
	}

	restrictionsConfigDef := di.Def{
		Name:  "v4a_restrictions_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v4a3.NewBaseRestrictionsConfig(cfg)
		},
	}

	err := builder.Add(
		loggerConfigDef,
		deadlineDef,
		configDef,
		restrictionsConfigDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerV5Config(builder di.Builder) {
	loggerConfigDef := di.Def{
		Name:  "v5_logger_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v53.NewBaseLoggerConfig(cfg)
		},
	}

	configDef := di.Def{
		Name:  "v5_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v53.NewBaseConfig(cfg)
		},
	}

	deadlineDef := di.Def{
		Name:  "v5_deadline_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v53.NewBaseDeadlineConfig(cfg)
		},
	}

	usersConfigDef := di.Def{
		Name:  "users_config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get("config_container").(gabs.Container)

			return v53.NewBaseUsersConfig(cfg)
		},
	}

	err := builder.Add(
		loggerConfigDef,
		configDef,
		deadlineDef,
		usersConfigDef,
	)

	if err != nil {
		panic(err)
	}
}
