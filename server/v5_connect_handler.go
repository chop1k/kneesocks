package server

import (
	"net"
	"os"
	"socks/config"
	"socks/logger"
	v5 "socks/protocol/v5"
	"socks/utils"
	"syscall"
	"time"
)

type V5ConnectHandler interface {
	HandleV5Connect(address string, client net.Conn)
}

type BaseV5ConnectHandler struct {
	protocol      v5.Protocol
	config        config.SocksV5Config
	streamHandler StreamHandler
	logger        logger.SocksV5Logger
	tcpConfig     config.TcpConfig
	utils         utils.AddressUtils
}

func NewBaseV5ConnectHandler(
	protocol v5.Protocol,
	config config.SocksV5Config,
	streamHandler StreamHandler,
	logger logger.SocksV5Logger,
	tcpConfig config.TcpConfig,
	addressUtils utils.AddressUtils,
) (BaseV5ConnectHandler, error) {
	return BaseV5ConnectHandler{
		protocol:      protocol,
		config:        config,
		streamHandler: streamHandler,
		logger:        logger,
		tcpConfig:     tcpConfig,
		utils:         addressUtils,
	}, nil
}

func (b BaseV5ConnectHandler) sendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5ConnectHandler) sendConnectionRefusedAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithConnectionRefused(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5ConnectHandler) sendNetworkUnreachableAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithNetworkUnreachable(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5ConnectHandler) sendHostUnreachableAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithHostUnreachable(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5ConnectHandler) HandleV5Connect(address string, client net.Conn) {
	deadline := time.Second * time.Duration(b.config.GetConnectDeadline())

	host, err := net.DialTimeout("tcp", address, deadline)

	if err != nil {
		if b.checkConnectionRefusedError(err) {
			b.sendConnectionRefusedAndClose(client)

			b.logger.ConnectRefused(client.RemoteAddr().String(), address)

			return
		}

		if b.checkNetworkUnreachableError(err) {
			b.sendNetworkUnreachableAndClose(client)

			b.logger.ConnectNetworkUnreachable(client.RemoteAddr().String(), address)

			return
		}

		if b.checkHostUnreachableError(err) {
			b.sendHostUnreachableAndClose(client)

			b.logger.ConnectHostUnreachable(client.RemoteAddr().String(), address)

			return
		}

		b.sendFailAndClose(client)

		//b.errors.UnknownConnectError(client.RemoteAddr().String(), err)
		b.logger.ConnectFailed(client.RemoteAddr().String(), address)

		return
	}

	b.connectSendResponse(host, client)
}

func (b BaseV5ConnectHandler) connectSendResponse(host, client net.Conn) {
	addr, port, parseErr := b.utils.ParseAddress(host.RemoteAddr().String())

	if parseErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		//b.errors.ParseAddressError(client.RemoteAddr().String(), host.RemoteAddr().String())

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(addr)

	if determineErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		//b.errors.DetermineAddressError(client.RemoteAddr().String(), addr)

		return
	}

	responseErr := b.protocol.ResponseWithSuccess(addrType, addr, uint16(port), client)

	if responseErr != nil {
		b.sendFailAndClose(client)

		_ = host.Close()

		// check and log

		return
	}

	b.logger.ConnectSuccessful(client.RemoteAddr().String(), host.RemoteAddr().String())

	go b.streamHandler.ClientToHost(client, host)
	b.streamHandler.HostToClient(client, host)
}

func (b BaseV5ConnectHandler) errorToErrno(err error) int {
	opErr, ok := err.(*net.OpError)

	if !ok {
		return -1
	}

	sysErr, ko := opErr.Err.(*os.SyscallError)

	if !ko {
		return -1
	}

	errno, oo := sysErr.Err.(syscall.Errno)

	if !oo {
		return -1
	}

	return int(errno)
}

func (b BaseV5ConnectHandler) checkConnectionRefusedError(err error) bool {
	return b.errorToErrno(err) == 111
}

func (b BaseV5ConnectHandler) checkNetworkUnreachableError(err error) bool {
	return b.errorToErrno(err) == 101
}

func (b BaseV5ConnectHandler) checkHostUnreachableError(err error) bool {
	errno := b.errorToErrno(err)

	return errno == 113 || errno == 112
}
