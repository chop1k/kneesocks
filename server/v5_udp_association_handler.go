package server

import (
	"net"
	"socks/config"
	"socks/logger"
	"socks/utils"
)

type V5UdpAssociationHandler interface {
	HandleV5UdpAssociation(client net.Conn)
}

type BaseV5UdpAssociationHandler struct {
	config                config.SocksV5Config
	utils                 utils.AddressUtils
	udpAssociationManager UdpAssociationManager
	logger                logger.SocksV5Logger
	sender                V5Sender
	errorHandler          V5ErrorHandler
}

func NewBaseV5UdpAssociationHandler(
	config config.SocksV5Config,
	utils utils.AddressUtils,
	udpAssociationManager UdpAssociationManager,
	logger logger.SocksV5Logger,
	sender V5Sender,
	errorHandler V5ErrorHandler,
) (BaseV5UdpAssociationHandler, error) {
	return BaseV5UdpAssociationHandler{
		config:                config,
		utils:                 utils,
		udpAssociationManager: udpAssociationManager,
		logger:                logger,
		sender:                sender,
		errorHandler:          errorHandler,
	}, nil
}

func (b BaseV5UdpAssociationHandler) HandleV5UdpAssociation(client net.Conn) {
	address, _, err := b.utils.ParseAddress(client.RemoteAddr().String())

	if err != nil {
		b.sender.SendFailAndClose(client)

		b.logger.UdpAssociationFailed(client.RemoteAddr().String(), address)

		return
	}

	b.udpAssociationManager.Set(address, client.RemoteAddr())

	b.udpSendResponse(address, client)
}

func (b BaseV5UdpAssociationHandler) udpSendResponse(address string, client net.Conn) {
	err := b.sender.SendSuccessWithUdpPort(client)

	if err != nil {
		b.errorHandler.HandleV5UdpAssociationError(err, address, client)

		return
	}

	b.logger.UdpAssociationSuccessful(client.RemoteAddr().String(), address)

	b.udpWaitForClose(address, client)
}

func (b BaseV5UdpAssociationHandler) udpWaitForClose(address string, client net.Conn) {
	buffer := make([]byte, 512)

	_, _ = client.Read(buffer)

	_ = client.Close()

	b.udpAssociationManager.Remove(address)
}
