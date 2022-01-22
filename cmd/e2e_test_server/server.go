package main

import (
	"fmt"
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

func (s Server) listenTcp() {
	address := &net.TCPAddr{
		IP:   net.ParseIP(s.config.Tcp.BindIP),
		Port: int(s.config.Tcp.BindPort),
		Zone: s.config.Tcp.BindZone,
	}

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

		s.logger.Connection(conn.RemoteAddr().String())

		go s.connectionHandler.HandleConnection(conn)
	}
}

func (s Server) listenUdp() {
	address := fmt.Sprintf("%s:%d", s.config.Udp.BindIP, s.config.Udp.BindPort)

	listener, err := net.ListenPacket("udp", address)

	if err != nil {
		panic(err)
	}

	s.logger.ListenUdp(address)

	for {
		buf := make([]byte, 1024)

		_, addr, err := listener.ReadFrom(buf)

		if err != nil {
			s.logger.AcceptPacketError(address, err)

			continue
		}

		go s.packetHandler.HandlePacket(buf[0], addr, listener)
	}
}

func (s Server) Start() {
	go s.listenUdp()

	s.listenTcp()
}
