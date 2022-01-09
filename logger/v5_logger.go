package logger

import (
	"socks/config"
	"socks/logger/output"
	v5 "socks/protocol/v5"
	"strconv"
)

type SocksV5Logger interface {
	ConnectRequest(client string, chunk v5.RequestChunk)
	ConnectFailed(client string, chunk v5.RequestChunk)
	ConnectSuccessful(client string, chunk v5.RequestChunk)
	BindRequest(client string, chunk v5.RequestChunk)
	BindFailed(client string, chunk v5.RequestChunk)
	BindSuccessful(client string, chunk v5.RequestChunk)
	Bound(client string, host string, chunk v5.RequestChunk)
	UdpAssociationRequest(client string, chunk v5.RequestChunk)
	UdpAssociationSuccessful(client string, chunk v5.RequestChunk)
	UdpAssociationFailed(client string, chunk v5.RequestChunk)
	AuthenticationSuccessful(client string)
	AuthenticationFailed(client string)
	TransferFinished(client string, host string)
}

type BaseSocksV5Logger struct {
	outputs []Output
	config  config.SocksV5LoggerConfig
	enabled bool
}

func NewBaseSocksV5Logger(config config.SocksV5LoggerConfig, replacer string, enabled bool) BaseSocksV5Logger {
	if enabled {
		var outputs []Output

		if config.IsConsoleOutputEnabled() {
			outputs = append(outputs, output.NewConsoleOutput(replacer))
		}

		if config.IsFileOutputEnabled() {
			outputs = append(outputs, output.NewFileOutput(config.GetFilePathFormat(), replacer))
		}

		return BaseSocksV5Logger{
			outputs: outputs,
			config:  config,
			enabled: true,
		}
	} else {
		return BaseSocksV5Logger{
			enabled: false,
		}
	}
}

func (b BaseSocksV5Logger) ConnectRequest(client string, chunk v5.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV5Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectRequestFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) ConnectFailed(client string, chunk v5.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV5Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectFailedFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) ConnectSuccessful(client string, chunk v5.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV5Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetConnectSuccessfulFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) BindRequest(client string, chunk v5.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV5Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetBindRequestFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) BindFailed(client string, chunk v5.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV5Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetBindFailedFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) BindSuccessful(client string, chunk v5.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV5Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetBindSuccessfulFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) Bound(client string, host string, chunk v5.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV5Parameters(client, chunk)

	parameters["host"] = host

	for _, output := range b.outputs {
		output.Log(b.config.GetBoundFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) UdpAssociationRequest(client string, chunk v5.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV5Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetUdpAssociationRequestFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) UdpAssociationSuccessful(client string, chunk v5.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV5Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetUdpAssociationSuccessfulFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) UdpAssociationFailed(client string, chunk v5.RequestChunk) {
	if !b.enabled {
		return
	}

	parameters := getBasicSocksV5Parameters(client, chunk)

	for _, output := range b.outputs {
		output.Log(b.config.GetUdpAssociationFailedFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) AuthenticationSuccessful(client string) {
	if !b.enabled {
		return
	}

	parameters := getBasicParameters()

	parameters["client"] = client

	for _, output := range b.outputs {
		output.Log(b.config.GetAuthenticationSuccessfulFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) AuthenticationFailed(client string) {
	if !b.enabled {
		return
	}

	parameters := getBasicParameters()

	parameters["client"] = client

	for _, output := range b.outputs {
		output.Log(b.config.GetAuthenticationFailedFormat(), parameters)
	}
}

func (b BaseSocksV5Logger) TransferFinished(client string, host string) {
}

func getBasicSocksV5Parameters(client string, chunk v5.RequestChunk) map[string]string {
	parameters := getBasicParameters()

	parameters["client"] = client

	parameters["chunk.CommandCode"] = string(chunk.CommandCode)
	parameters["chunk.Address"] = chunk.Address
	parameters["chunk.Port"] = strconv.Itoa(int(chunk.Port))
	parameters["chunk.SocksVersion"] = string(chunk.SocksVersion)

	return parameters
}
