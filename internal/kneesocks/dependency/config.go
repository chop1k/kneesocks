package dependency

import (
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
		Build: build.TcpBaseConfig,
	}

	handlerDef := di.Def{
		Name:  "tcp_config_handler",
		Scope: di.App,
		Build: build.TcpConfigHandler,
	}

	loggerConfigDef := di.Def{
		Name:  "tcp_logger_config",
		Scope: di.App,
		Build: build.TcpLoggerConfig,
	}

	replicatorDef := di.Def{
		Name:  "tcp_config_replicator",
		Scope: di.App,
		Build: build.TcpConfigReplicator,
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
		Build: build.UdpBaseConfig,
	}

	handlerDef := di.Def{
		Name:  "udp_config_handler",
		Scope: di.App,
		Build: build.UdpConfigHandler,
	}

	loggerConfigDef := di.Def{
		Name:  "udp_logger_config",
		Scope: di.App,
		Build: build.UdpLoggerConfig,
	}

	replicatorDef := di.Def{
		Name:  "udp_config_replicator",
		Scope: di.App,
		Build: build.UdpConfigReplicator,
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
		Build: build.V4BaseConfig,
	}

	handlerDef := di.Def{
		Name:  "v4_config_handler",
		Scope: di.App,
		Build: build.V4ConfigHandler,
	}

	loggerConfigDef := di.Def{
		Name:  "v4_logger_config",
		Scope: di.App,
		Build: build.V4LoggerConfig,
	}

	replicatorDef := di.Def{
		Name:  "v4_config_replicator",
		Scope: di.App,
		Build: build.V4ConfigReplicator,
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
		Build: build.V4aBaseConfig,
	}

	handlerDef := di.Def{
		Name:  "v4a_config_handler",
		Scope: di.App,
		Build: build.V4aConfigHandler,
	}

	loggerConfigDef := di.Def{
		Name:  "v4a_logger_config",
		Scope: di.App,
		Build: build.V4aLoggerConfig,
	}

	replicatorDef := di.Def{
		Name:  "v4a_config_replicator",
		Scope: di.App,
		Build: build.V4aConfigReplicator,
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
		Build: build.V5BaseConfig,
	}

	handlerDef := di.Def{
		Name:  "v5_config_handler",
		Scope: di.App,
		Build: build.V5ConfigHandler,
	}

	loggerConfigDef := di.Def{
		Name:  "v5_logger_config",
		Scope: di.App,
		Build: build.V5LoggerConfig,
	}

	replicatorDef := di.Def{
		Name:  "v5_config_replicator",
		Scope: di.App,
		Build: build.V5ConfigReplicator,
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
