package server

import (
	"fmt"
	"net"
	"socks/config/tcp"
	"socks/handlers"
	tcp2 "socks/logger/tcp"
)

type TcpServer struct {
	connectionHandler handlers.ConnectionHandler
	logger            tcp2.Logger
	config            tcp.BindConfig
}

func NewTcpServer(
	connectionHandler handlers.ConnectionHandler,
	tcpLogger tcp2.Logger,
	tcpConfig tcp.BindConfig,
) (TcpServer, error) {
	return TcpServer{
		connectionHandler: connectionHandler,
		logger:            tcpLogger,
		config:            tcpConfig,
	}, nil
}

func (s TcpServer) getAddress() string {
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

func (s TcpServer) Listen() {
	address := s.getAddress()

	listener, err := net.Listen("tcp", address)

	if err != nil {
		s.logger.Errors.ListenError(address, err)

		panic(err)
	}

	s.logger.Listen.Listen(address)

	for {
		conn, err := listener.Accept()

		if err != nil {
			continue
		}

		s.logger.Connection.Accepted(conn.RemoteAddr().String())

		go s.connectionHandler.HandleConnection(conn)
	}
}
