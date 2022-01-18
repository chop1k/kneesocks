package server

import (
	"fmt"
	"net"
	"socks/config"
	"socks/logger"
)

type Server struct {
	connectionHandler ConnectionHandler
	packetHandler     PacketHandler
	tcpLogger         logger.TcpLogger
	tcpConfig         config.TcpConfig
	udpLogger         logger.UdpLogger
	udpConfig         config.UdpConfig
}

func NewServer(
	connectionHandler ConnectionHandler,
	packetHandler PacketHandler,
	tcpLogger logger.TcpLogger,
	tcpConfig config.TcpConfig,
	udpLogger logger.UdpLogger,
	udpConfig config.UdpConfig,
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
		Port: s.tcpConfig.GetBindPort(),
		Zone: s.tcpConfig.GetBindZone(),
	}

	listener, err := net.ListenTCP("tcp", address)

	if err != nil {
		s.tcpLogger.ListenError(address.String(), err)

		panic(err)
	}

	s.tcpLogger.Listen(address.String())

	for {
		conn, err := listener.AcceptTCP()

		if err != nil {
			continue
		}

		s.tcpLogger.ConnectionAccepted(conn.RemoteAddr().String())

		go s.connectionHandler.HandleConnection(conn)
	}
}

func (s Server) listenUdp() {
	address := fmt.Sprintf("%s:%d", s.udpConfig.GetBindIp(), s.udpConfig.GetBindPort())

	packet, err := net.ListenPacket("udp", address)

	if err != nil {
		s.udpLogger.ListenError(address, err)

		panic(err)
	}

	s.udpLogger.Listen(address)

	for {
		payload := make([]byte, s.udpConfig.GetBufferSize())

		i, address, err := packet.ReadFrom(payload)

		if err != nil {
			s.udpLogger.AcceptError(err)

			continue
		}

		s.udpLogger.PacketAccepted(address.String())

		go s.packetHandler.HandlePacket(payload[:i], address, packet)
	}
}

func (s Server) Start() {
	go s.listenUdp()

	s.listenTcp()
}
