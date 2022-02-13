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
	HandleUdpAssociation(config v5.Config, name string, client net.Conn)
}

type BaseUdpAssociationHandler struct {
	utils         utils.AddressUtils
	clientManager managers.UdpClientManager
	logger        v52.Logger
	sender        v53.Sender
	errorHandler  ErrorHandler
}

func NewBaseUdpAssociationHandler(
	utils utils.AddressUtils,
	clientManager managers.UdpClientManager,
	logger v52.Logger,
	sender v53.Sender,
	errorHandler ErrorHandler,
) (BaseUdpAssociationHandler, error) {
	return BaseUdpAssociationHandler{
		utils:         utils,
		clientManager: clientManager,
		logger:        logger,
		sender:        sender,
		errorHandler:  errorHandler,
	}, nil
}

func (b BaseUdpAssociationHandler) HandleUdpAssociation(config v5.Config, _ string, client net.Conn) {
	address, _, err := b.utils.ParseAddress(client.RemoteAddr().String())

	if err != nil {
		b.errorHandler.HandleUdpAddressParsingError(config, err, client)

		return
	}

	err = b.clientManager.Add(address)

	if err != nil {
	}

	b.udpSendResponse(config, address, client)
}

func (b BaseUdpAssociationHandler) udpSendResponse(config v5.Config, address string, client net.Conn) {
	err := b.sender.SendSuccessWithUdpPort(config, client)

	if err != nil {
		b.errorHandler.HandleUdpAssociationError(config, err, address, client)

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
