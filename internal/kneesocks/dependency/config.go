package dependency

import (
	"os"
	"socks/internal/kneesocks/config/tcp"
	"socks/internal/kneesocks/config/tree"
	"socks/internal/kneesocks/config/udp"
	v4 "socks/internal/kneesocks/config/v4"
	"socks/internal/kneesocks/config/v4a"
	v5 "socks/internal/kneesocks/config/v5"

	"github.com/go-playground/validator/v10"
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
		Build: func(ctn di.Container) (interface{}, error) {
			path, ok := os.LookupEnv("tcp_config_path")

			if !ok {
				return "/etc/kneesocks/proxy/tcp.json", nil
			}

			return path, nil
		},
	}

	udpConfigPathDef := di.Def{
		Name:  "udp_config_path",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path, ok := os.LookupEnv("udp_config_path")

			if !ok {
				return "/etc/kneesocks/proxy/udp.json", nil
			}

			return path, nil
		},
	}

	httpConfigPathDef := di.Def{
		Name:  "http_config_path",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path, ok := os.LookupEnv("http_config_path")

			if !ok {
				return "/etc/kneesocks/proxy/http.json", nil
			}

			return path, nil
		},
	}

	v4ConfigPathDef := di.Def{
		Name:  "v4_config_path",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path, ok := os.LookupEnv("v4_config_path")

			if !ok {
				return "/etc/kneesocks/proxy/v4.json", nil
			}

			return path, nil
		},
	}

	v4aConfigPathDef := di.Def{
		Name:  "v4a_config_path",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path, ok := os.LookupEnv("v4a_config_path")

			if !ok {
				return "/etc/kneesocks/proxy/v4a.json", nil
			}

			return path, nil
		},
	}

	v5ConfigPathDef := di.Def{
		Name:  "v5_config_path",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path, ok := os.LookupEnv("v5_config_path")

			if !ok {
				return "/etc/kneesocks/proxy/v5.json", nil
			}

			return path, nil
		},
	}

	logConfigPathDef := di.Def{
		Name:  "log_config_path",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path, ok := os.LookupEnv("log_config_path")

			if !ok {
				return "/etc/kneesocks/proxy/log.json", nil
			}

			return path, nil
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			path := ctn.Get("tcp_config_path").(string)

			file, err := os.Open(path)

			if err != nil {
				return nil, err
			}

			return file, nil
		},
	}

	udpConfigFileDef := di.Def{
		Name:  "udp_config_file",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path := ctn.Get("udp_config_path").(string)

			file, err := os.Open(path)

			if err != nil {
				return nil, err
			}

			return file, nil
		},
	}

	httpConfigFileDef := di.Def{
		Name:  "http_config_file",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path := ctn.Get("http_config_path").(string)

			file, err := os.Open(path)

			if err != nil {
				return nil, err
			}

			return file, nil
		},
	}

	v4ConfigFileDef := di.Def{
		Name:  "v4_config_file",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path := ctn.Get("v4_config_path").(string)

			file, err := os.Open(path)

			if err != nil {
				return nil, err
			}

			return file, nil
		},
	}

	v4aConfigFileDef := di.Def{
		Name:  "v4a_config_file",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path := ctn.Get("v4a_config_path").(string)

			file, err := os.Open(path)

			if err != nil {
				return nil, err
			}

			return file, nil
		},
	}

	v5ConfigFileDef := di.Def{
		Name:  "v5_config_file",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path := ctn.Get("v5_config_path").(string)

			file, err := os.Open(path)

			if err != nil {
				return nil, err
			}

			return file, nil
		},
	}

	logConfigFileDef := di.Def{
		Name:  "log_config_file",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			path := ctn.Get("log_config_path").(string)

			file, err := os.Open(path)

			if err != nil {
				return nil, err
			}

			return file, nil
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			return *validator.New(), nil
		},
	}

	tcpDef := di.Def{
		Name:  "tcp_tree",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			file := ctn.Get("tcp_config_file").(*os.File)
			builder := ctn.Get("tcp_tree_builder").(tree.TcpBuilder)

			defer file.Close()

			return builder.Build(file)
		},
	}

	udpDef := di.Def{
		Name:  "udp_tree",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			file := ctn.Get("udp_config_file").(*os.File)
			builder := ctn.Get("udp_tree_builder").(tree.UdpBuilder)

			defer file.Close()

			return builder.Build(file)
		},
	}

	httpDef := di.Def{
		Name:  "http_tree",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return nil, nil
		},
	}

	v4Def := di.Def{
		Name:  "v4_tree",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			file := ctn.Get("v4_config_file").(*os.File)
			builder := ctn.Get("v4_tree_builder").(tree.SocksV4Builder)

			defer file.Close()

			return builder.Build(file)
		},
	}

	v4aDef := di.Def{
		Name:  "v4a_tree",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			file := ctn.Get("v4a_config_file").(*os.File)
			builder := ctn.Get("v4a_tree_builder").(tree.SocksV4aBuilder)

			defer file.Close()

			return builder.Build(file)
		},
	}

	v5Def := di.Def{
		Name:  "v5_tree",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			file := ctn.Get("v5_config_file").(*os.File)
			builder := ctn.Get("v5_tree_builder").(tree.SocksV5Builder)

			defer file.Close()

			return builder.Build(file)
		},
	}

	logDef := di.Def{
		Name:  "log_tree",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			file := ctn.Get("log_config_file").(*os.File)
			builder := ctn.Get("log_tree_builder").(tree.LogBuilder)

			defer file.Close()

			return builder.Build(file)
		},
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
		Build: func(ctn di.Container) (interface{}, error) {
			validator := ctn.Get("validator").(validator.Validate)

			return tree.NewTcpBuilder(validator)
		},
	}

	udpDef := di.Def{
		Name:  "udp_tree_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			validator := ctn.Get("validator").(validator.Validate)

			return tree.NewUdpBuilder(validator)
		},
	}

	httpDef := di.Def{
		Name:  "http_tree_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return nil, nil
		},
	}

	v4Def := di.Def{
		Name:  "v4_tree_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			validator := ctn.Get("validator").(validator.Validate)

			return tree.NewSocksV4Builder(validator)
		},
	}

	v4aDef := di.Def{
		Name:  "v4a_tree_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			validator := ctn.Get("validator").(validator.Validate)

			return tree.NewSocksV4aBuilder(validator)
		},
	}

	v5Def := di.Def{
		Name:  "v5_tree_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			validator := ctn.Get("validator").(validator.Validate)

			return tree.NewSocksV5Builder(validator)
		},
	}

	logDef := di.Def{
		Name:  "log_tree_builder",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			validator := ctn.Get("validator").(validator.Validate)

			return tree.NewLogBuilder(validator)
		},
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
			config := ctn.Get("log_tree").(tree.LogConfig).Tcp

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
			config := ctn.Get("udp_tree").(tree.UdpConfig)
			handler := ctn.Get("udp_config_handler").(udp.Handler)

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
			config := ctn.Get("log_tree").(tree.LogConfig).Udp

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
			config := ctn.Get("v4_tree").(*tree.SocksV4Config)
			handler := ctn.Get("v4_config_handler").(v4.Handler)

			return handler.Handle(config)
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
			config := ctn.Get("log_tree").(tree.LogConfig).SocksV4

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
			config := ctn.Get("v4a_tree").(*tree.SocksV4aConfig)
			handler := ctn.Get("v4a_config_handler").(v4a.Handler)

			return handler.Handle(config)
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
			config := ctn.Get("log_tree").(tree.LogConfig).SocksV4a

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
			config := ctn.Get("v5_tree").(*tree.SocksV5Config)
			handler := ctn.Get("v5_config_handler").(v5.Handler)

			return handler.Handle(config)
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
			config := ctn.Get("log_tree").(tree.LogConfig).SocksV5

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
