package logger

import (
	"socks/config"
	"socks/logger/output"
)

type TcpLogger interface {
	ConnectionAccepted(addr string)
	ConnectionDenied(addr string)
	ConnectionProtocolDetermined(addr string, protocol string)
	ConnectionBound(client string, host string)
	Listen(addr string)
}

type BaseTcpLogger struct {
	outputs []Output
	config  config.TcpLoggerConfig
	enabled bool
}

func NewBaseTcpLogger(config config.TcpLoggerConfig, replacer string, enabled bool) BaseTcpLogger {
	if enabled {
		var outputs []Output

		if config.IsConsoleOutputEnabled() {
			outputs = append(outputs, output.NewConsoleOutput(replacer))
		}

		if config.IsFileOutputEnabled() {
			outputs = append(outputs, output.NewFileOutput(config.GetFilePathFormat(), replacer))
		}

		return BaseTcpLogger{
			outputs: outputs,
			config:  config,
			enabled: true,
		}

	} else {
		return BaseTcpLogger{
			enabled: false,
		}
	}
}

func (b BaseTcpLogger) ConnectionAccepted(client string) {
	if !b.enabled {
		return
	}

	parameters := getBasicParameters()

	parameters["client"] = client

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectionAcceptedFormat(), parameters)
	}
}

func (b BaseTcpLogger) ConnectionDenied(client string) {
	if !b.enabled {
		return
	}

	parameters := getBasicParameters()

	parameters["client"] = client

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectionDeniedFormat(), parameters)
	}
}

func (b BaseTcpLogger) ConnectionProtocolDetermined(client string, protocol string) {
	if !b.enabled {
		return
	}

	parameters := getBasicParameters()

	parameters["client"] = client
	parameters["protocol"] = protocol

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectionProtocolDeterminedFormat(), parameters)
	}
}

func (b BaseTcpLogger) ConnectionBound(client string, host string) {
	if !b.enabled {
		return
	}

	parameters := getBasicParameters()

	parameters["client"] = client
	parameters["host"] = host

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectionBoundFormat(), parameters)
	}
}

func (b BaseTcpLogger) Listen(addr string) {
	if !b.enabled {
		return
	}

	parameters := getBasicParameters()

	parameters["addr"] = addr

	for _, output := range b.outputs {
		output.Log(b.config.GetListenFormat(), parameters)
	}
}
