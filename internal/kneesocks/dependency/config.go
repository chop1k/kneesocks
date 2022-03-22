package dependency

import (
	"socks/internal/kneesocks/config/tcp"
	"socks/internal/kneesocks/config/tree"
	"socks/internal/kneesocks/config/udp"
	v4 "socks/internal/kneesocks/config/v4"
	v4a "socks/internal/kneesocks/config/v4a"
	v5 "socks/internal/kneesocks/config/v5"
	"socks/internal/kneesocks/dependency/build"

	"github.com/sarulabs/di"
)

func registerConfig(builder di.Builder) {
	registerTree(builder)
	registerPath(builder)
	registerFiles(builder)
	registerBuilder(builder)
	registerTcpConfig(builder)
	registerUdpConfig(builder)
	registerV4Config(builder)
	registerV4aConfig(builder)
	registerV5Config(builder)
}

func registerPath(builder di.Builder) {
	tcpConfigPathDef := di.Def{
		Name:  "tcp_config_path",
		Scope: di.App,
		Build: build.TcpConfigPath,
	}

	udpConfigPathDef := di.Def{
		Name:  "udp_config_path",
		Scope: di.App,
		Build: build.UdpConfigPath,
	}

	httpConfigPathDef := di.Def{
		Name:  "http_config_path",
		Scope: di.App,
		Build: build.HttpConfigPath,
	}

	v4ConfigPathDef := di.Def{
		Name:  "v4_config_path",
		Scope: di.App,
		Build: build.V4ConfigPath,
	}

	v4aConfigPathDef := di.Def{
		Name:  "v4a_config_path",
		Scope: di.App,
		Build: build.V4aConfigPath,
	}

	v5ConfigPathDef := di.Def{
		Name:  "v5_config_path",
		Scope: di.App,
		Build: build.V5ConfigPath,
	}

	logConfigPathDef := di.Def{
		Name:  "log_config_path",
		Scope: di.App,
		Build: build.LogConfigPath,
	}

	err := builder.Add(
		tcpConfigPathDef,
		udpConfigPathDef,
		httpConfigPathDef,
		v4ConfigPathDef,
		v4aConfigPathDef,
		v5ConfigPathDef,
		logConfigPathDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerFiles(builder di.Builder) {
	tcpConfigFileDef := di.Def{
		Name:  "tcp_config_file",
		Scope: di.App,
		Build: build.TcpConfigFile,
	}

	udpConfigFileDef := di.Def{
		Name:  "udp_config_file",
		Scope: di.App,
		Build: build.UdpConfigFile,
	}

	httpConfigFileDef := di.Def{
		Name:  "http_config_file",
		Scope: di.App,
		Build: build.HttpConfigFile,
	}

	v4ConfigFileDef := di.Def{
		Name:  "v4_config_file",
		Scope: di.App,
		Build: build.V4ConfigFile,
	}

	v4aConfigFileDef := di.Def{
		Name:  "v4a_config_file",
		Scope: di.App,
		Build: build.V4aConfigFile,
	}

	v5ConfigFileDef := di.Def{
		Name:  "v5_config_file",
		Scope: di.App,
		Build: build.V5ConfigFile,
	}

	logConfigFileDef := di.Def{
		Name:  "log_config_file",
		Scope: di.App,
		Build: build.LogConfigFile,
	}

	err := builder.Add(
		tcpConfigFileDef,
		udpConfigFileDef,
		httpConfigFileDef,
		v4ConfigFileDef,
		v4aConfigFileDef,
		v5ConfigFileDef,
		logConfigFileDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerTree(builder di.Builder) {
	validatorDef := di.Def{
		Name:  "validator",
		Scope: di.App,
		Build: build.Validator,
	}

	tcpDef := di.Def{
		Name:  "tcp_tree",
		Scope: di.App,
		Build: build.TcpTree,
	}

	udpDef := di.Def{
		Name:  "udp_tree",
		Scope: di.App,
		Build: build.UdpTree,
	}

	httpDef := di.Def{
		Name:  "http_tree",
		Scope: di.App,
		Build: build.HttpTree,
	}

	v4Def := di.Def{
		Name:  "v4_tree",
		Scope: di.App,
		Build: build.V4Tree,
	}

	v4aDef := di.Def{
		Name:  "v4a_tree",
		Scope: di.App,
		Build: build.V4aTree,
	}

	v5Def := di.Def{
		Name:  "v5_tree",
		Scope: di.App,
		Build: build.V5Tree,
	}

	logDef := di.Def{
		Name:  "log_tree",
		Scope: di.App,
		Build: build.LogTree,
	}

	err := builder.Add(
		validatorDef,
		tcpDef,
		udpDef,
		httpDef,
		v4Def,
		v4aDef,
		v5Def,
		logDef,
	)

	if err != nil {
		panic(err)
	}
}

func registerBuilder(builder di.Builder) {
	tcpDef := di.Def{
		Name:  "tcp_tree_builder",
		Scope: di.App,
		Build: build.TcpTreeBuilder,
	}

	udpDef := di.Def{
		Name:  "udp_tree_builder",
		Scope: di.App,
		Build: build.UdpTreeBuilder,
	}

	httpDef := di.Def{
		Name:  "http_tree_builder",
		Scope: di.App,
		Build: build.HttpTreeBuilder,
	}

	v4Def := di.Def{
		Name:  "v4_tree_builder",
		Scope: di.App,
		Build: build.V4TreeBuilder,
	}

	v4aDef := di.Def{
		Name:  "v4a_tree_builder",
		Scope: di.App,
		Build: build.V4aTreeBuilder,
	}

	v5Def := di.Def{
		Name:  "v5_tree_builder",
		Scope: di.App,
		Build: build.V5TreeBuilder,
	}

	logDef := di.Def{
		Name:  "log_tree_builder",
		Scope: di.App,
		Build: build.LogTreeBuilder,
	}

	err := builder.Add(
		tcpDef,
		udpDef,
		httpDef,
		v4Def,
		v4aDef,
		v5Def,
		logDef,
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
			config := ctn.Get("tcp_tree").(tree.TcpConfig)
			handler := ctn.Get("tcp_config_handler").(tcp.Handler)

			return handler.Handle(config), nil
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
			_config := ctn.Get("log_tree")

			if _config == nil {
				return tcp.NewLoggerConfig(nil)
			}

			config := _config.(*tree.TcpLoggerConfig)

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
			_config := ctn.Get("udp_tree")
			handler := ctn.Get("udp_config_handler").(udp.Handler)

			if _config == nil {
				return nil, nil
			}

			config := _config.(tree.UdpConfig)

			return handler.Handle(config), nil
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
			_config := ctn.Get("log_tree")

			if _config == nil {
				return udp.NewLoggerConfig(nil)
			}

			__config := _config.(*tree.LogConfig)

			if __config == nil {
				return udp.NewLoggerConfig(nil)
			}

			config := __config.Udp

			return udp.NewLoggerConfig(config)
		},
	}

	replicatorDef := di.Def{
		Name:  "udp_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			_base := ctn.Get("udp_base_config")

			if _base == nil {
				return nil, nil
			}

			base := _base.(udp.Config)

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
			_config := ctn.Get("v4_tree")
			handler := ctn.Get("v4_config_handler").(v4.Handler)

			if _config == nil {
				return nil, nil
			}

			config := _config.(tree.SocksV4Config)

			return handler.Handle(config), nil
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
			_config := ctn.Get("log_tree")

			if _config == nil {
				return v4.NewLoggerConfig(nil)
			}

			__config := _config.(*tree.LogConfig)

			if __config == nil {
				return udp.NewLoggerConfig(nil)
			}

			config := __config.SocksV4

			return v4.NewLoggerConfig(config)
		},
	}

	replicatorDef := di.Def{
		Name:  "v4_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			_config := ctn.Get("v4_base_config")

			if _config == nil {
				return nil, nil
			}

			config := _config.(v4.Config)

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
			_config := ctn.Get("v4a_tree")
			handler := ctn.Get("v4a_config_handler").(v4a.Handler)

			if _config == nil {
				return nil, nil
			}

			config := _config.(tree.SocksV4aConfig)

			return handler.Handle(config), nil
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
			_config := ctn.Get("log_tree")

			if _config == nil {
				return v4a.NewLoggerConfig(nil)
			}

			__config := _config.(*tree.LogConfig)

			if __config == nil {
				return udp.NewLoggerConfig(nil)
			}

			config := __config.SocksV4a

			return v4a.NewLoggerConfig(config)
		},
	}

	replicatorDef := di.Def{
		Name:  "v4a_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			_config := ctn.Get("v4a_base_config")

			if _config == nil {
				return nil, nil
			}

			config := _config.(v4a.Config)

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
			_config := ctn.Get("v5_tree")
			handler := ctn.Get("v5_config_handler").(v5.Handler)

			if _config == nil {
				return nil, nil
			}

			config := _config.(tree.SocksV5Config)

			return handler.Handle(config), nil
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
			_config := ctn.Get("log_tree")

			if _config == nil {
				return v5.NewLoggerConfig(nil)
			}

			__config := _config.(*tree.LogConfig)

			if __config == nil {
				return udp.NewLoggerConfig(nil)
			}

			config := __config.SocksV5

			return v5.NewLoggerConfig(config)
		},
	}

	replicatorDef := di.Def{
		Name:  "v5_config_replicator",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			_config := ctn.Get("v5_base_config")

			if _config == nil {
				return nil, nil
			}

			config := _config.(v5.Config)

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
