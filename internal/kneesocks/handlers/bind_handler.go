package handlers

import (
	"fmt"
	"net"
	tcp2 "socks/internal/kneesocks/config/tcp"
	"socks/internal/kneesocks/logger/tcp"
	"socks/internal/kneesocks/managers"
	"socks/internal/kneesocks/transfer"
	"socks/pkg/utils"
)

type BindHandler struct {
	utils       utils.AddressUtils
	logger      tcp.Logger
	bind        transfer.BindHandler
	bindManager managers.BindManager
}

func NewBindHandler(
	utils utils.AddressUtils,
	logger tcp.Logger,
	bind transfer.BindHandler,
	bindManager managers.BindManager,
) (BindHandler, error) {
	return BindHandler{
		utils:       utils,
		logger:      logger,
		bind:        bind,
		bindManager: bindManager,
	}, nil
}

func (b BindHandler) Handle(config tcp2.DeadlineConfig, request []byte, host net.Conn) {
	addr := host.RemoteAddr().String()

	if b.bindManager.IsBound(addr) {
		b.exchange(config, request, addr, host)
	} else {
		b.checkDomain(config, request, addr, host)
	}
}

func (b BindHandler) checkDomain(config tcp2.DeadlineConfig, request []byte, address string, host net.Conn) {
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

		if b.bindManager.IsBound(address) {
			b.exchange(config, request, address, host)

			return
		}
	}

	_ = host.Close()

	b.logger.Connection.Denied(address)
}

func (b BindHandler) exchange(config tcp2.DeadlineConfig, request []byte, address string, host net.Conn) {
	err := b.bindManager.SendHost(address, host)

	if err != nil {
		_ = host.Close()

		b.logger.Errors.SendHostError(address, err)

		return
	}

	client, receiveErr := b.bindManager.ReceiveClient(address, config.Exchange)

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
