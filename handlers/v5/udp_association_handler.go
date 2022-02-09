package v5

import (
	"net"
	"socks/config/v5"
	v52 "socks/logger/v5"
	"socks/managers"
	v53 "socks/protocol/v5"
	"socks/utils"
)

type UdpAssociationHandler interface {
	HandleUdpAssociation(name string, client net.Conn)
}

type BaseUdpAssociationHandler struct {
	config        v5.Config
	utils         utils.AddressUtils
	clientManager managers.UdpClientManager
	logger        v52.Logger
	sender        v53.Sender
	errorHandler  ErrorHandler
}

func NewBaseUdpAssociationHandler(
	config v5.Config,
	utils utils.AddressUtils,
	clientManager managers.UdpClientManager,
	logger v52.Logger,
	sender v53.Sender,
	errorHandler ErrorHandler,
) (BaseUdpAssociationHandler, error) {
	return BaseUdpAssociationHandler{
		config:        config,
		utils:         utils,
		clientManager: clientManager,
		logger:        logger,
		sender:        sender,
		errorHandler:  errorHandler,
	}, nil
}

func (b BaseUdpAssociationHandler) HandleUdpAssociation(_ string, client net.Conn) {
	address, _, err := b.utils.ParseAddress(client.RemoteAddr().String())

	if err != nil {
		b.errorHandler.HandleUdpAddressParsingError(err, client)

		return
	}

	err = b.clientManager.Add(address)

	if err != nil {
	}

	b.udpSendResponse(address, client)
}

func (b BaseUdpAssociationHandler) udpSendResponse(address string, client net.Conn) {
	err := b.sender.SendSuccessWithUdpPort(client)

	if err != nil {
		b.errorHandler.HandleUdpAssociationError(err, address, client)

		return
	}

	b.logger.Association.Successful(client.RemoteAddr().String(), address)

	b.udpWaitForClose(address, client)
}

func (b BaseUdpAssociationHandler) udpWaitForClose(address string, client net.Conn) {
	buffer := make([]byte, 512)

	_, _ = client.Read(buffer)

	_ = client.Close()

	b.clientManager.Remove(address)
}
