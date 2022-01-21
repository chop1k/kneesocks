package main

import "net"

type Server struct {
	config            Config
	connectionHandler ConnectionHandler
	logger            Logger
}

func NewServer(config Config, connectionHandler ConnectionHandler, logger Logger) Server {
	return Server{
		config:            config,
		connectionHandler: connectionHandler,
		logger:            logger,
	}
}

func (s Server) listenTcp() {
	address := &net.TCPAddr{
		IP:   net.ParseIP(s.config.BindIP),
		Port: int(s.config.BindPort),
		Zone: s.config.BindZone,
	}

	listener, err := net.ListenTCP("tcp", address)

	if err != nil {
		panic(err)
	}

	s.logger.Listen(address.String())

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

func (s Server) Start() {
}
