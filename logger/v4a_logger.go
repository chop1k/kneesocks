package logger

import (
	"socks/config"
	"socks/protocol/v4a"
	"strconv"
)

type SocksV4aLogger interface {
	ConnectRequest(client string, chunk v4a.RequestChunk)
	ConnectFailed(client string, chunk v4a.RequestChunk)
	ConnectSuccessful(client string, chunk v4a.RequestChunk)
	BindRequest(client string, chunk v4a.RequestChunk)
	BindFailed(client string, chunk v4a.RequestChunk)
	BindSuccessful(client string, chunk v4a.RequestChunk)
	Bound(client string, host string, chunk v4a.RequestChunk)
	TransferFinished(client string, host string)
}

type BaseSocksV4aLogger struct {
	outputs []Output
	config  config.SocksV4aLoggerConfig
	enabled bool
}

func NewBaseSocksV4aLogger(config config.SocksV4aLoggerConfig, replacer string, enabled bool) BaseSocksV4aLogger {
	if enabled {
		var outputs []Output

		if config.IsConsoleOutputEnabled() {
			outputs = append(outputs, NewConsoleOutput(replacer))
		}

		if config.IsFileOutputEnabled() {
			outputs = append(outputs, NewFileOutput(config.GetFilePathFormat(), replacer))
		}

		return BaseSocksV4aLogger{
			outputs: outputs,
			config:  config,
			enabled: true,
		}
	} else {
		return BaseSocksV4aLogger{
			enabled: false,
		}
	}
}

func (b BaseSocksV4aLogger) ConnectRequest(client string, chunk v4a.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4aParameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectRequestFormat(), parameters)
	}
}

func (b BaseSocksV4aLogger) ConnectFailed(client string, chunk v4a.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4aParameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectFailedFormat(), parameters)
	}
}

func (b BaseSocksV4aLogger) ConnectSuccessful(client string, chunk v4a.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4aParameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectSuccessfulFormat(), parameters)
	}
}

func (b BaseSocksV4aLogger) BindRequest(client string, chunk v4a.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4aParameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetBindRequestFormat(), parameters)
	}
}

func (b BaseSocksV4aLogger) BindFailed(client string, chunk v4a.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4aParameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetBindFailedFormat(), parameters)
	}
}

func (b BaseSocksV4aLogger) BindSuccessful(client string, chunk v4a.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4aParameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetBindSuccessfulFormat(), parameters)
	}
}

func (b BaseSocksV4aLogger) Bound(client string, host string, chunk v4a.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4aParameters(client, chunk)

	parameters["host"] = host

	for _, output := range b.outputs {
		output.Log(b.config.GetBoundFormat(), parameters)
	}
}

func (b BaseSocksV4aLogger) TransferFinished(client string, host string) {

}

func getBasicSocksV4aParameters(client string, chunk v4a.RequestChunk) map[string]string {
	parameters := getBasicParameters()

	parameters["client"] = client

	parameters["chunk.CommandCode"] = string(chunk.CommandCode)
	parameters["chunk.DestinationIp"] = chunk.DestinationIp.String()
	parameters["chunk.DestinationPort"] = strconv.Itoa(int(chunk.DestinationPort))
	parameters["chunk.SocksVersion"] = string(chunk.SocksVersion)
	parameters["chunk.Domain"] = chunk.Domain

	return parameters
}
