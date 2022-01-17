package server

import (
	"net"
	"socks/config"
	"socks/logger"
	v5 "socks/protocol/v5"
	"socks/utils"
)

type V5UdpAssociationHandler interface {
	HandleV5UdpAssociation(client net.Conn)
}

type BaseV5UdpAssociationHandler struct {
	protocol              v5.Protocol
	config                config.SocksV5Config
	utils                 utils.AddressUtils
	udpAssociationManager UdpAssociationManager
	logger                logger.SocksV5Logger
	tcpConfig             config.TcpConfig
	udpConfig             config.UdpConfig
}

func NewBaseV5UdpAssociationHandler(
	protocol v5.Protocol,
	config config.SocksV5Config,
	utils utils.AddressUtils,
	udpAssociationManager UdpAssociationManager,
	logger logger.SocksV5Logger,
	tcpConfig config.TcpConfig,
	udpConfig config.UdpConfig,
) (BaseV5UdpAssociationHandler, error) {
	return BaseV5UdpAssociationHandler{
		protocol:              protocol,
		config:                config,
		utils:                 utils,
		udpAssociationManager: udpAssociationManager,
		logger:                logger,
		tcpConfig:             tcpConfig,
		udpConfig:             udpConfig,
	}, nil
}

func (b BaseV5UdpAssociationHandler) sendFailAndClose(client net.Conn) {
	_ = b.protocol.ResponseWithFail(1, "0.0.0.0", uint16(b.tcpConfig.GetBindPort()), client)
	_ = client.Close()
}

func (b BaseV5UdpAssociationHandler) HandleV5UdpAssociation(client net.Conn) {
	address, _, err := b.utils.ParseAddress(client.RemoteAddr().String())

	if err != nil {
		b.sendFailAndClose(client)

		b.logger.UdpAssociationFailed(client.RemoteAddr().String(), address)

		return
	}

	b.udpAssociationManager.Set(address, client.RemoteAddr())

	b.udpSendResponse(address, client)
}

func (b BaseV5UdpAssociationHandler) udpSendResponse(address string, client net.Conn) {
	err := b.protocol.ResponseWithSuccess(1, "0.0.0.0", uint16(b.udpConfig.GetBindPort()), client)

	if err != nil {
		b.sendFailAndClose(client)

		b.logger.UdpAssociationFailed(client.RemoteAddr().String(), address)

		return
	}

	b.logger.UdpAssociationSuccessful(client.RemoteAddr().String(), address)

	b.udpWaitForClose(address, client)
}

func (b BaseV5UdpAssociationHandler) udpWaitForClose(address string, client net.Conn) {
	for {
		buffer := make([]byte, 512)

		_, err := client.Read(buffer)

		_ = buffer

		if err != nil {
			_ = client.Close()

			break
		}
	}

	b.udpAssociationManager.Remove(address)
}
