package server

import (
	"fmt"
	"net"
	"socks/internal/kneesocks/config/udp"
	"socks/internal/kneesocks/handlers"
	udp2 "socks/internal/kneesocks/logger/udp"
)

type UdpServer struct {
	logger        udp2.Logger
	packetHandler handlers.PacketHandler
	bindConfig    udp.BindConfig
	replicator    udp.ConfigReplicator
}

func NewUdpServer(
	logger udp2.Logger,
	packetHandler handlers.PacketHandler,
	bindConfig udp.BindConfig,
	replicator udp.ConfigReplicator,
) (UdpServer, error) {
	return UdpServer{
		logger:        logger,
		packetHandler: packetHandler,
		bindConfig:    bindConfig,
		replicator:    replicator,
	}, nil
}

func (s UdpServer) Listen() {
	address := fmt.Sprintf("%s:%d", s.bindConfig.Address, s.bindConfig.Port)

	packet, err := net.ListenPacket("udp", address)

	if err != nil {
		s.logger.Errors.ListenError(address, err)

		panic(err)
	}

	s.logger.Listen.Listen(address)

	for {
		size := s.replicator.CopyBuffer().PacketSize

		payload := make([]byte, size)

		i, address, err := packet.ReadFrom(payload)

		if err != nil {
			s.logger.Errors.AcceptError(err)

			continue
		}

		s.logger.Packet.Accepted(address.String())

		go s.packetHandler.HandlePacket(payload[:i], address, packet)
	}
}
