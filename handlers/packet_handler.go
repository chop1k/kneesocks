package handlers

import (
	"fmt"
	"net"
	udp2 "socks/config/udp"
	"socks/logger/udp"
	"socks/managers"
	v5 "socks/protocol/v5"
	"socks/utils"
	"time"
)

type PacketHandler struct {
	parser     v5.Parser
	builder    v5.Builder
	utils      utils.AddressUtils
	clients    managers.UdpClientManager
	hosts      managers.UdpHostManager
	logger     udp.Logger
	replicator udp2.ConfigReplicator
}

func NewPacketHandler(
	parser v5.Parser,
	builder v5.Builder,
	utils utils.AddressUtils,
	clients managers.UdpClientManager,
	hosts managers.UdpHostManager,
	logger udp.Logger,
	replicator udp2.ConfigReplicator,
) PacketHandler {
	return PacketHandler{
		parser:     parser,
		builder:    builder,
		utils:      utils,
		clients:    clients,
		hosts:      hosts,
		logger:     logger,
		replicator: replicator,
	}
}

func (b PacketHandler) HandlePacket(payload []byte, address net.Addr, conn net.PacketConn) {
	chunk, err := b.parser.ParseUdpRequest(payload)

	if err != nil {
		b.checkBound(payload, address.String(), conn)

		return
	}

	if chunk.Fragment != 0 {
		b.logger.Packet.InvalidFragment(address.String())

		return
	}

	addr, _, parseErr := b.utils.ParseAddress(address.String())

	if parseErr != nil {
		b.logger.Errors.ParseAddressError(address.String(), address.String(), parseErr)

		return
	}

	if b.clients.IsBound(addr) {
		b.sendPacket(chunk, address, conn)
	}
}

func (b PacketHandler) checkBound(payload []byte, address string, conn net.PacketConn) {
	client, err := b.hosts.Get(address)

	if err != nil {
		b.logger.Errors.ManagerError(address, err)

		return
	}

	addr, resolveErr := net.ResolveUDPAddr("udp", client)

	if resolveErr != nil {
		b.logger.Errors.ResolveError(address, client, resolveErr)

		return
	}

	_, err = conn.WriteTo(payload, addr)

	if err != nil {
		b.logger.Errors.WriteError(client, err)

		return
	}
}

func (b PacketHandler) sendPacket(chunk v5.UdpRequest, client net.Addr, conn net.PacketConn) {
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
		b.logger.Errors.ResolveError(client.String(), address, err)

		return
	}

	packet, dialErr := net.DialUDP("udp", nil, addr)

	if dialErr != nil {
		b.logger.Errors.DialError(client.String(), addr.String(), dialErr)

		return
	}

	_, _, err = packet.WriteMsgUDP(chunk.Data, nil, nil)

	if err != nil {
		b.logger.Errors.WriteMsgError(client.String(), addr.String(), err)

		_ = packet.Close()

		return
	}

	b.listen(client, packet, conn)
}

func (b PacketHandler) listen(client net.Addr, packet net.PacketConn, server net.PacketConn) {
	for {
		size := b.replicator.CopyBuffer().PacketSize

		deadline := b.replicator.CopyDeadline().Read

		deadlineErr := packet.SetReadDeadline(time.Now().Add(deadline))

		if deadlineErr != nil {
			b.logger.Errors.DeadlineError(client.String(), deadlineErr)

			break
		}

		buffer := make([]byte, size)

		i, address, err := packet.ReadFrom(buffer)

		if err != nil {
			break
		}

		b.sendToClient(buffer[:i], client, address, server)
	}

	_ = packet.Close()
}

func (b PacketHandler) sendToClient(data []byte, client net.Addr, host net.Addr, server net.PacketConn) {
	hostAddr, hostPort, parseErr := b.utils.ParseAddress(host.String())

	if parseErr != nil {
		b.logger.Errors.ParseAddressError(client.String(), host.String(), parseErr)

		return
	}

	addrType, determineErr := b.utils.DetermineAddressType(hostAddr)

	if determineErr != nil {
		b.logger.Errors.DeterminationError(client.String(), hostAddr, determineErr)

		return
	}

	response, err := b.builder.BuildUdpRequest(v5.UdpRequest{
		Fragment:    0,
		AddressType: addrType,
		Address:     hostAddr,
		Port:        uint16(hostPort),
		Data:        data,
	})

	if err != nil {
		b.logger.Errors.BuildError(client.String(), err)

		return
	}

	_, err = server.WriteTo(response, client)

	if err != nil {
		b.logger.Errors.WriteError(client.String(), err)

		return
	}
}
