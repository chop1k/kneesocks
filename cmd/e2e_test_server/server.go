package main

import (
	"net"
)

type Server struct {
	config            Config
	connectionHandler ConnectionHandler
	packetHandler     PacketHandler
	logger            Logger
}

func NewServer(
	config Config,
	connectionHandler ConnectionHandler,
	packetHandler PacketHandler,
	logger Logger,
) (Server, error) {
	return Server{
		config:            config,
		connectionHandler: connectionHandler,
		packetHandler:     packetHandler,
		logger:            logger,
	}, nil
}

func (s Server) listenTcp(address *net.TCPAddr) {
	listener, err := net.ListenTCP("tcp", address)

	if err != nil {
		panic(err)
	}

	s.logger.ListenTcp(address.String())

	for {
		conn, err := listener.Accept()

		if err != nil {
			s.logger.AcceptError(conn.RemoteAddr().String(), err)

			continue
		}

		s.logger.Connection(conn.RemoteAddr().String(), address.String())

		go s.connectionHandler.HandleConnection(conn)
	}
}

func (s Server) listenUdp(address *net.UDPAddr) {
	listener, err := net.ListenUDP("udp", address)

	if err != nil {
		panic(err)
	}

	s.logger.ListenUdp(address.String())

	for {
		buf := make([]byte, 1024)

		_, addr, err := listener.ReadFrom(buf)

		if err != nil {
			s.logger.AcceptPacketError(address.String(), err)

			continue
		}

		s.logger.PacketAccepted(addr.String(), address.String())

		go s.packetHandler.HandlePacket(buf[0], addr, listener)
	}
}

func (s Server) Start() {
	udpAddressV4 := &net.UDPAddr{
		IP:   net.ParseIP(s.config.Udp.BindIPv4),
		Port: int(s.config.Tcp.BindPort),
		Zone: s.config.Tcp.BindZone,
	}

	go s.listenUdp(udpAddressV4)

	udpAddressV6 := &net.UDPAddr{
		IP:   net.ParseIP(s.config.Udp.BindIPv6),
		Port: int(s.config.Tcp.BindPort),
		Zone: s.config.Tcp.BindZone,
	}

	go s.listenUdp(udpAddressV6)

	tcpAddressV4 := &net.TCPAddr{
		IP:   net.ParseIP(s.config.Tcp.BindIPv4),
		Port: int(s.config.Tcp.BindPort),
		Zone: s.config.Tcp.BindZone,
	}

	go s.listenTcp(tcpAddressV4)

	tcpAddressV6 := &net.TCPAddr{
		IP:   net.ParseIP(s.config.Tcp.BindIPv6),
		Port: int(s.config.Tcp.BindPort),
		Zone: s.config.Tcp.BindZone,
	}

	s.listenTcp(tcpAddressV6)
}
