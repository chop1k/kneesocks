package logger

import (
	"socks/config"
	v4 "socks/protocol/v4"
	"strconv"
)

type SocksV4Logger interface {
	ConnectRequest(client string, chunk v4.RequestChunk)
	ConnectFailed(client string, chunk v4.RequestChunk)
	ConnectSuccessful(client string, chunk v4.RequestChunk)
	BindRequest(client string, chunk v4.RequestChunk)
	BindFailed(client string, chunk v4.RequestChunk)
	BindSuccessful(client string, chunk v4.RequestChunk)
	Bound(client string, host string, chunk v4.RequestChunk)
	TransferFinished(client string, host string)
}

type BaseSocksV4Logger struct {
	outputs []Output
	config  config.SocksV4LoggerConfig
	enabled bool
}

func NewBaseSocksV4Logger(config config.SocksV4LoggerConfig, replacer string, enabled bool) BaseSocksV4Logger {
	if enabled {
		var outputs []Output

		if config.IsConsoleOutputEnabled() {
			outputs = append(outputs, NewConsoleOutput(replacer))
		}

		if config.IsFileOutputEnabled() {
			outputs = append(outputs, NewFileOutput(config.GetFilePathFormat(), replacer))
		}

		return BaseSocksV4Logger{
			outputs: outputs,
			config:  config,
			enabled: true,
		}
	} else {
		return BaseSocksV4Logger{
			enabled: false,
		}
	}
}

func (b BaseSocksV4Logger) ConnectRequest(client string, chunk v4.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectRequestFormat(), parameters)
	}
}

func (b BaseSocksV4Logger) ConnectFailed(client string, chunk v4.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectFailedFormat(), parameters)
	}
}

func (b BaseSocksV4Logger) ConnectSuccessful(client string, chunk v4.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectSuccessfulFormat(), parameters)
	}
}

func (b BaseSocksV4Logger) BindRequest(client string, chunk v4.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetBindRequestFormat(), parameters)
	}
}

func (b BaseSocksV4Logger) BindFailed(client string, chunk v4.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetBindFailedFormat(), parameters)
	}
}

func (b BaseSocksV4Logger) BindSuccessful(client string, chunk v4.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetBindSuccessfulFormat(), parameters)
	}
}

func (b BaseSocksV4Logger) Bound(client string, host string, chunk v4.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV4Parameters(client, chunk)

	parameters["host"] = host

	for _, output := range b.outputs {
		output.Log(b.config.GetBoundFormat(), parameters)
	}
}

func (b BaseSocksV4Logger) TransferFinished(client string, host string) {

}

func getBasicSocksV4Parameters(client string, chunk v4.RequestChunk) map[string]string {
	parameters := getBasicParameters()

	parameters["client"] = client

	parameters["chunk.CommandCode"] = string(chunk.CommandCode)
	parameters["chunk.DestinationIp"] = chunk.DestinationIp.String()
	parameters["chunk.DestinationPort"] = strconv.Itoa(int(chunk.DestinationPort))
	parameters["chunk.SocksVersion"] = string(chunk.SocksVersion)
	parameters["chunk.UserId"] = chunk.UserId

	return parameters
}
