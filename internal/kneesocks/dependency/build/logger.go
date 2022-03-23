package build

import (
	"io"
	"os"
	"path"
	tcpConfig "socks/internal/kneesocks/config/tcp"
	udpConfig "socks/internal/kneesocks/config/udp"
	v4Config "socks/internal/kneesocks/config/v4"
	v4aConfig "socks/internal/kneesocks/config/v4a"
	v5Config "socks/internal/kneesocks/config/v5"
	"socks/internal/kneesocks/logger/tcp"
	"socks/internal/kneesocks/logger/udp"
	v4 "socks/internal/kneesocks/logger/v4"
	"socks/internal/kneesocks/logger/v4a"
	v5 "socks/internal/kneesocks/logger/v5"
	"socks/pkg/utils"

	"github.com/rs/zerolog"
	"github.com/sarulabs/di"
)

func TcpZeroLogger(ctn di.Container) (interface{}, error) {
	config := ctn.Get("tcp_logger_config").(tcpConfig.LoggerConfig)

	level, err := config.GetLevel()

	var loggers []io.Writer

	if err != nil {
		return utils.BuildDefaultZerolog(126, loggers)
	}

	if output, err := config.GetConsoleOutput(); err == nil {
		loggers = append(loggers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: output.TimeFormat,
		})
	} else {
		if err == tcpConfig.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	if output, err := config.GetFileOutput(); err == nil {
		dir := path.Dir(output.Path)

		dirErr := os.MkdirAll(dir, 0700)

		if dirErr != nil {
			return zerolog.Logger{}, dirErr
		}

		file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

		if err != nil {
			return zerolog.Logger{}, err
		}

		loggers = append(loggers, file)
	} else {
		if err == tcpConfig.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	return utils.BuildDefaultZerolog(level, loggers)
}

func UdpZeroLogger(ctn di.Container) (interface{}, error) {
	config := ctn.Get("udp_logger_config").(udpConfig.LoggerConfig)

	level, err := config.GetLevel()

	var loggers []io.Writer

	if err != nil {
		return utils.BuildDefaultZerolog(126, loggers)
	}

	if output, err := config.GetConsoleOutput(); err == nil {
		loggers = append(loggers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: output.TimeFormat,
		})
	} else {
		if err == udpConfig.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	if output, err := config.GetFileOutput(); err == nil {
		dir := path.Dir(output.Path)

		dirErr := os.MkdirAll(dir, 0700)

		if dirErr != nil {
			return zerolog.Logger{}, dirErr
		}

		file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

		if err != nil {
			return zerolog.Logger{}, err
		}

		loggers = append(loggers, file)
	} else {
		if err == udpConfig.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	return utils.BuildDefaultZerolog(level, loggers)
}

func V4ZeroLogger(ctn di.Container) (interface{}, error) {
	config := ctn.Get("v4_logger_config").(v4Config.LoggerConfig)

	level, err := config.GetLevel()

	var loggers []io.Writer

	if err != nil {
		return utils.BuildDefaultZerolog(126, loggers)
	}

	if output, err := config.GetConsoleOutput(); err == nil {
		loggers = append(loggers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: output.TimeFormat,
		})
	} else {
		if err == v4Config.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	if output, err := config.GetFileOutput(); err == nil {
		dir := path.Dir(output.Path)

		dirErr := os.MkdirAll(dir, 0700)

		if dirErr != nil {
			return zerolog.Logger{}, dirErr
		}

		file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

		if err != nil {
			return zerolog.Logger{}, err
		}

		loggers = append(loggers, file)
	} else {
		if err == v4Config.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	return utils.BuildDefaultZerolog(level, loggers)
}

func V4aZeroLogger(ctn di.Container) (interface{}, error) {
	config := ctn.Get("v4a_logger_config").(v4aConfig.LoggerConfig)

	level, err := config.GetLevel()

	var loggers []io.Writer

	if err != nil {
		return utils.BuildDefaultZerolog(126, loggers)
	}

	if output, err := config.GetConsoleOutput(); err == nil {
		loggers = append(loggers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: output.TimeFormat,
		})
	} else {
		if err == v4aConfig.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	if output, err := config.GetFileOutput(); err == nil {
		dir := path.Dir(output.Path)

		dirErr := os.MkdirAll(dir, 0700)

		if dirErr != nil {
			return zerolog.Logger{}, dirErr
		}

		file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

		if err != nil {
			return zerolog.Logger{}, err
		}

		loggers = append(loggers, file)
	} else {
		if err == v4aConfig.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	return utils.BuildDefaultZerolog(level, loggers)
}

func V5ZeroLogger(ctn di.Container) (interface{}, error) {
	config := ctn.Get("v5_logger_config").(v5Config.LoggerConfig)

	level, err := config.GetLevel()

	var loggers []io.Writer

	if err != nil {
		return utils.BuildDefaultZerolog(126, loggers)
	}

	if output, err := config.GetConsoleOutput(); err == nil {
		loggers = append(loggers, zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: output.TimeFormat,
		})
	} else {
		if err == v5Config.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	if output, err := config.GetFileOutput(); err == nil {
		dir := path.Dir(output.Path)

		dirErr := os.MkdirAll(dir, 0700)

		if dirErr != nil {
			return zerolog.Logger{}, dirErr
		}

		file, err := os.OpenFile(output.Path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)

		if err != nil {
			return zerolog.Logger{}, err
		}

		loggers = append(loggers, file)
	} else {
		if err == v5Config.LoggerDisabledError {
			return utils.BuildDefaultZerolog(126, loggers)
		}
	}

	return utils.BuildDefaultZerolog(level, loggers)
}

func TcpConnectionLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("tcp_zero_logger").(zerolog.Logger)

	return tcp.NewConnectionLogger(zero)
}

func TcpErrorsLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("tcp_zero_logger").(zerolog.Logger)

	return tcp.NewErrorsLogger(zero)
}

func TcpListenLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("tcp_zero_logger").(zerolog.Logger)

	return tcp.NewListenLogger(zero)
}

func TcpLogger(ctn di.Container) (interface{}, error) {
	connection := ctn.Get("tcp_connection_logger").(tcp.ConnectionLogger)
	errors := ctn.Get("tcp_errors_logger").(tcp.ErrorsLogger)
	listen := ctn.Get("tcp_listen_logger").(tcp.ListenLogger)

	return tcp.NewLogger(connection, errors, listen)
}

func UdpErrorsLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("udp_zero_logger").(zerolog.Logger)

	return udp.NewErrorsLogger(zero)
}

func UdpListenLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("udp_zero_logger").(zerolog.Logger)

	return udp.NewListenLogger(zero)
}

func UdpPacketLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("udp_zero_logger").(zerolog.Logger)

	return udp.NewPacketLogger(zero)
}

func UdpLogger(ctn di.Container) (interface{}, error) {
	errors := ctn.Get("udp_errors_logger").(udp.ErrorsLogger)
	listen := ctn.Get("udp_listen_logger").(udp.ListenLogger)
	packet := ctn.Get("udp_packet_logger").(udp.PacketLogger)

	return udp.NewLogger(errors, listen, packet)
}

func V4BindLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

	return v4.NewBindLogger(zero)
}

func V4ConnectLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

	return v4.NewConnectLogger(zero)
}

func V4ErrorsLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

	return v4.NewErrorsLogger(zero)
}

func V4RestrictionsLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

	return v4.NewRestrictionsLogger(zero)
}

func V4TransferLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v4_zero_logger").(zerolog.Logger)

	return v4.NewTransferLogger(zero)
}

func V4Logger(ctn di.Container) (interface{}, error) {
	bind := ctn.Get("v4_bind_logger").(v4.BindLogger)
	connect := ctn.Get("v4_connect_logger").(v4.ConnectLogger)
	errors := ctn.Get("v4_errors_logger").(v4.ErrorsLogger)
	restrictions := ctn.Get("v4_restrictions_logger").(v4.RestrictionsLogger)
	transfer := ctn.Get("v4_transfer_logger").(v4.TransferLogger)

	return v4.NewLogger(bind, connect, errors, restrictions, transfer)
}

func V4aBindLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

	return v4a.NewBindLogger(zero)
}

func V4aConnectLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

	return v4a.NewConnectLogger(zero)
}

func V4aErrorsLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

	return v4a.NewErrorsLogger(zero)
}

func V4aRestrictionsLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

	return v4a.NewRestrictionsLogger(zero)
}

func V4aTransferLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v4a_zero_logger").(zerolog.Logger)

	return v4a.NewTransferLogger(zero)
}

func V4aLogger(ctn di.Container) (interface{}, error) {
	bind := ctn.Get("v4a_bind_logger").(v4a.BindLogger)
	connect := ctn.Get("v4a_connect_logger").(v4a.ConnectLogger)
	errors := ctn.Get("v4a_errors_logger").(v4a.ErrorsLogger)
	restrictions := ctn.Get("v4a_restrictions_logger").(v4a.RestrictionsLogger)
	transfer := ctn.Get("v4a_transfer_logger").(v4a.TransferLogger)

	return v4a.NewLogger(bind, connect, errors, restrictions, transfer)
}

func V5AssociationLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

	return v5.NewAssociationLogger(zero)
}

func V5AuthLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

	return v5.NewAuthLogger(zero)
}

func V5BindLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

	return v5.NewBindLogger(zero)
}

func V5ConnectLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

	return v5.NewConnectLogger(zero)
}

func V5ErrorsLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

	return v5.NewErrorsLogger(zero)
}

func V5RestrictionsLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

	return v5.NewRestrictionsLogger(zero)
}

func V5TransferLogger(ctn di.Container) (interface{}, error) {
	zero := ctn.Get("v5_zero_logger").(zerolog.Logger)

	return v5.NewTransferLogger(zero)
}

func V5Logger(ctn di.Container) (interface{}, error) {
	association := ctn.Get("v5_association_logger").(v5.AssociationLogger)
	auth := ctn.Get("v5_auth_logger").(v5.AuthLogger)
	bind := ctn.Get("v5_bind_logger").(v5.BindLogger)
	connect := ctn.Get("v5_connect_logger").(v5.ConnectLogger)
	errors := ctn.Get("v5_errors_logger").(v5.ErrorsLogger)
	restrictions := ctn.Get("v5_restrictions_logger").(v5.RestrictionsLogger)
	transfer := ctn.Get("v5_transfer_logger").(v5.TransferLogger)

	return v5.NewLogger(association, auth, bind, connect, errors, restrictions, transfer)
}
