package server

import (
	"fmt"
	"net"
	"socks/config/udp"
	"socks/handlers"
	udp2 "socks/logger/udp"
)

type UdpServer struct {
	config        udp.BindConfig
	logger        udp2.Logger
	packetHandler handlers.PacketHandler
	buffer        udp.BufferConfig
}

func NewUdpServer(
	config udp.BindConfig,
	logger udp2.Logger,
	packetHandler handlers.PacketHandler,
	buffer udp.BufferConfig,
) (UdpServer, error) {
	return UdpServer{
		config:        config,
		logger:        logger,
		packetHandler: packetHandler,
		buffer:        buffer,
	}, nil
}

func (s UdpServer) getAddress() string {
	address, err := s.config.GetAddress()

	if err != nil {
		panic(err)
	}

	port, configErr := s.config.GetPort()

	if configErr != nil {
		panic(configErr)
	}

	return fmt.Sprintf("%s:%d", address, port)
}

func (s UdpServer) Listen() {
	address := s.getAddress()

	packet, err := net.ListenPacket("udp", address)

	if err != nil {
		s.logger.Errors.ListenError(address, err)

		panic(err)
	}

	s.logger.Listen.Listen(address)

	for {
		size, configErr := s.buffer.GetPacketSize()

		if configErr != nil {
			panic(configErr)
		}

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
