package main

import (
	"net"
)

type Server struct {
	config         Config
	connectHandler ConnectHandler
	packetHandler  PacketHandler
	logger         Logger
	bindHandler    BindHandler
}

func NewServer(
	config Config,
	connectHandler ConnectHandler,
	packetHandler PacketHandler,
	logger Logger,
	bindHandler BindHandler,
) (Server, error) {
	return Server{
		config:         config,
		connectHandler: connectHandler,
		packetHandler:  packetHandler,
		logger:         logger,
		bindHandler:    bindHandler,
	}, nil
}

func (s Server) listenConnect(address *net.TCPAddr) {
	listener, err := net.ListenTCP("tcp", address)

	if err != nil {
		panic(err)
	}

	s.logger.ListenConnect(address.String())

	for {
		conn, err := listener.Accept()

		if err != nil {
			s.logger.AcceptError(conn.RemoteAddr().String(), err)

			continue
		}

		s.logger.Connection(conn.RemoteAddr().String(), address.String())

		go s.connectHandler.HandleConnect(conn)
	}
}

func (s Server) listenBind(address *net.TCPAddr) {
	listener, err := net.ListenTCP("tcp", address)

	if err != nil {
		panic(err)
	}

	s.logger.ListenBind(address.String())

	for {
		conn, err := listener.Accept()

		if err != nil {
			s.logger.AcceptError(conn.RemoteAddr().String(), err)

			continue
		}

		s.logger.Connection(conn.RemoteAddr().String(), address.String())

		go s.bindHandler.HandleBind(conn)
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
		Port: int(s.config.Udp.BindPort),
		Zone: s.config.Udp.BindZone,
	}

	go s.listenUdp(udpAddressV4)

	udpAddressV6 := &net.UDPAddr{
		IP:   net.ParseIP(s.config.Udp.BindIPv6),
		Port: int(s.config.Udp.BindPort),
		Zone: s.config.Udp.BindZone,
	}

	go s.listenUdp(udpAddressV6)

	connectAddressV4 := &net.TCPAddr{
		IP:   net.ParseIP(s.config.Tcp.BindIPv4),
		Port: int(s.config.Tcp.ConnectPort),
		Zone: s.config.Tcp.BindZone,
	}

	go s.listenConnect(connectAddressV4)

	connectAddressV6 := &net.TCPAddr{
		IP:   net.ParseIP(s.config.Tcp.BindIPv6),
		Port: int(s.config.Tcp.ConnectPort),
		Zone: s.config.Tcp.BindZone,
	}

	go s.listenConnect(connectAddressV6)

	bindAddressV4 := &net.TCPAddr{
		IP:   net.ParseIP(s.config.Tcp.BindIPv4),
		Port: int(s.config.Tcp.BindPort),
		Zone: s.config.Tcp.BindZone,
	}

	go s.listenBind(bindAddressV4)

	bindAddressV6 := &net.TCPAddr{
		IP:   net.ParseIP(s.config.Tcp.BindIPv6),
		Port: int(s.config.Tcp.BindPort),
		Zone: s.config.Tcp.BindZone,
	}

	s.listenBind(bindAddressV6)
}
