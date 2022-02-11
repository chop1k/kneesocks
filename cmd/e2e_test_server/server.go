package main

import (
	"fmt"
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

func (s Server) listenConnect() {
	address := fmt.Sprintf("%s:%d", s.config.Tcp.BindAddress, s.config.Tcp.ConnectPort)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		panic(err)
	}

	s.logger.ListenConnect(address)

	for {
		conn, err := listener.Accept()

		if err != nil {
			s.logger.AcceptError(conn.RemoteAddr().String(), err)

			continue
		}

		s.logger.Connection(conn.RemoteAddr().String(), address)

		go s.connectHandler.HandleConnect(conn)
	}
}

func (s Server) listenBind() {
	address := fmt.Sprintf("%s:%d", s.config.Tcp.BindAddress, s.config.Tcp.BindPort)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		panic(err)
	}

	s.logger.ListenBind(address)

	for {
		conn, err := listener.Accept()

		if err != nil {
			s.logger.AcceptError(conn.RemoteAddr().String(), err)

			continue
		}

		s.logger.Connection(conn.RemoteAddr().String(), address)

		go s.bindHandler.HandleBind(conn)
	}
}

func (s Server) listenUdp() {
	address := fmt.Sprintf("%s:%d", s.config.Udp.BindAddress, s.config.Udp.BindPort)

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

		s.logger.PacketAccepted(addr.String(), address)

		go s.packetHandler.HandlePacket(buf[0], addr, listener)
	}
}

func (s Server) Start() {
	go s.listenUdp()
	go s.listenBind()

	s.listenConnect()
}
