package handlers

import (
	"fmt"
	"net"
	"socks/managers"
	v5 "socks/protocol/v5"
	"socks/utils"
)

type PacketHandler interface {
	HandlePacket(payload []byte, address net.Addr, conn net.PacketConn)
}

type BasePacketHandler struct {
	parser                v5.Parser
	builder               v5.Builder
	udpAssociationManager managers.UdpAssociationManager
	utils                 utils.AddressUtils
}

func NewBasePacketHandler(
	parser v5.Parser,
	builder v5.Builder,
	udpAssociationManager managers.UdpAssociationManager,
	utils utils.AddressUtils,
) BasePacketHandler {
	return BasePacketHandler{
		parser:                parser,
		builder:               builder,
		udpAssociationManager: udpAssociationManager,
		utils:                 utils,
	}
}

func (b BasePacketHandler) HandlePacket(payload []byte, address net.Addr, conn net.PacketConn) {
	chunk, err := b.parser.ParseUdpRequest(payload)

	if err != nil {
		b.checkBound(payload, address, conn)

		return
	}

	addr, _, parseErr := b.utils.ParseAddress(address.String())

	if parseErr != nil {
		return
	}

	_, ok := b.udpAssociationManager.Get(addr)

	if ok {
		b.sendPacket(chunk, address, conn)
	}
}

func (b BasePacketHandler) checkBound(payload []byte, address net.Addr, conn net.PacketConn) {
	client, ok := b.udpAssociationManager.GetHost(address.String())

	if !ok {
		b.checkDomain(payload, address, conn)

		return
	}

	_, _ = conn.WriteTo(payload, client)
}

func (b BasePacketHandler) checkDomain(payload []byte, addr net.Addr, conn net.PacketConn) {
	address, port, err := b.utils.ParseAddress(addr.String())

	if err != nil {
		return
	}

	addresses, lookupErr := net.LookupAddr(address)

	if lookupErr != nil {
		return
	}

	for _, domain := range addresses {
		address := fmt.Sprintf("%s:%d", domain, port)

		client, ok := b.udpAssociationManager.GetHost(address)

		if !ok {
			continue
		}

		_, _ = conn.WriteTo(payload, client)

		return
	}
}

func (b BasePacketHandler) sendPacket(chunk v5.UdpRequest, clientAddr net.Addr, conn net.PacketConn) {
	var address string

	if chunk.AddressType == 1 || chunk.AddressType == 3 {
		address = fmt.Sprintf("%s:%d", chunk.Address, chunk.Port)
	} else if chunk.AddressType == 4 {
		address = fmt.Sprintf("[%s]:%d", chunk.Address, chunk.Port)
	} else {
		return
	}

	addr, err := net.ResolveUDPAddr("udp", address)

	if err != nil {
		return
	}

	ip, _, parseErr := b.utils.ParseAddress(clientAddr.String())

	if parseErr != nil {
		return
	}

	err = b.udpAssociationManager.AddHost(address, ip, clientAddr)

	if err != nil {
		return
	}

	_, _ = conn.WriteTo(chunk.Data, addr)
}
