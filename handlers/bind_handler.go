package handlers

import (
	"fmt"
	"net"
	"socks/handlers/helpers"
	"socks/logger/tcp"
	"socks/transfer"
	"socks/utils"
)

type BindHandler interface {
	Handle(request []byte, host net.Conn)
}

type BaseBindHandler struct {
	utils  utils.AddressUtils
	binder helpers.Binder
	logger tcp.Logger
	bind   transfer.BindHandler
}

func NewBaseBindHandler(
	utils utils.AddressUtils,
	binder helpers.Binder,
	logger tcp.Logger,
	bind transfer.BindHandler,
) (BaseBindHandler, error) {
	return BaseBindHandler{
		utils:  utils,
		binder: binder,
		logger: logger,
		bind:   bind,
	}, nil
}

func (b BaseBindHandler) Handle(request []byte, host net.Conn) {
	addr := host.RemoteAddr().String()

	if b.binder.IsBound(addr) {
		b.exchange(request, addr, host)
	} else {
		b.checkDomain(request, addr, host)
	}
}

func (b BaseBindHandler) checkDomain(request []byte, address string, host net.Conn) {
	hostAddr, hostPort, parseErr := b.utils.ParseAddress(address)

	if parseErr != nil {
		_ = host.Close()

		b.logger.Errors.AddressParseError(address, parseErr)

		return
	}

	addresses, err := net.LookupAddr(hostAddr)

	if err != nil {
		_ = host.Close()

		b.logger.Errors.LookupError(hostAddr, err)

		return
	}

	for _, address := range addresses {
		address = fmt.Sprintf("%s:%d", address, hostPort)

		if b.binder.IsBound(address) {
			b.exchange(request, address, host)

			return
		}
	}

	_ = host.Close()

	b.logger.Connection.Denied(address)
}

func (b BaseBindHandler) exchange(request []byte, address string, host net.Conn) {
	err := b.binder.Send(address, host)

	if err != nil {
		_ = host.Close()

		b.logger.Errors.SendHostError(address, err)

		return
	}

	client, receiveErr := b.binder.Receive(address)

	if receiveErr != nil {
		_ = host.Close()

		b.logger.Errors.ReceiveClientError(address, receiveErr)

		return
	}

	b.logger.Connection.Bound(client.RemoteAddr().String(), host.RemoteAddr().String())

	_, err = client.Write(request)

	if err != nil {
		_ = host.Close()
		_ = client.Close()

		b.logger.Errors.WriteRequestError(client.RemoteAddr().String(), host.RemoteAddr().String(), err)

		return
	}

	b.bind.HandleHost(client, host)
}
