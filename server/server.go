package server

import (
	"fmt"
	"net"
	"socks/config/tcp"
	"socks/config/udp"
	"socks/handlers"
	tcp2 "socks/logger/tcp"
	udp2 "socks/logger/udp"
)

type Server struct {
	connectionHandler handlers.ConnectionHandler
	packetHandler     handlers.PacketHandler
	tcpLogger         tcp2.Logger
	tcpConfig         tcp.Config
	udpLogger         udp2.Logger
	udpConfig         udp.Config
}

func NewServer(
	connectionHandler handlers.ConnectionHandler,
	packetHandler handlers.PacketHandler,
	tcpLogger tcp2.Logger,
	tcpConfig tcp.Config,
	udpLogger udp2.Logger,
	udpConfig udp.Config,
) (Server, error) {
	return Server{
		connectionHandler: connectionHandler,
		packetHandler:     packetHandler,
		tcpLogger:         tcpLogger,
		tcpConfig:         tcpConfig,
		udpLogger:         udpLogger,
		udpConfig:         udpConfig,
	}, nil
}

func (s Server) listenTcp() {
	address := &net.TCPAddr{
		IP:   net.ParseIP(s.tcpConfig.GetBindIP()),
		Port: int(s.tcpConfig.GetBindPort()),
		Zone: s.tcpConfig.GetBindZone(),
	}

	listener, err := net.ListenTCP("tcp", address)

	if err != nil {
		s.tcpLogger.Errors.ListenError(address.String(), err)

		panic(err)
	}

	s.tcpLogger.Listen.Listen(address.String())

	for {
		conn, err := listener.AcceptTCP()

		if err != nil {
			continue
		}

		s.tcpLogger.Connection.Accepted(conn.RemoteAddr().String())

		go s.connectionHandler.HandleConnection(conn)
	}
}

func (s Server) listenUdp() {
	address := fmt.Sprintf("%s:%d", s.udpConfig.GetBindIp(), s.udpConfig.GetBindPort())

	packet, err := net.ListenPacket("udp", address)

	if err != nil {
		s.udpLogger.Errors.ListenError(address, err)

		panic(err)
	}

	s.udpLogger.Listen.Listen(address)

	for {
		payload := make([]byte, s.udpConfig.GetBufferSize())

		i, address, err := packet.ReadFrom(payload)

		if err != nil {
			s.udpLogger.Errors.AcceptError(err)

			continue
		}

		s.udpLogger.Packet.Accepted(address.String())

		go s.packetHandler.HandlePacket(payload[:i], address, packet)
	}
}

func (s Server) Start() {
	go s.listenUdp()

	s.listenTcp()
}
