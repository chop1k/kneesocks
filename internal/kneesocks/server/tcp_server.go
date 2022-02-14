package server

import (
	"fmt"
	"net"
	"socks/internal/kneesocks/config/tcp"
	"socks/internal/kneesocks/handlers"
	tcp2 "socks/internal/kneesocks/logger/tcp"
)

type TcpServer struct {
	connectionHandler handlers.ConnectionHandler
	logger            tcp2.Logger
	bind              tcp.BindConfig
}

func NewTcpServer(
	connectionHandler handlers.ConnectionHandler,
	tcpLogger tcp2.Logger,
	bind tcp.BindConfig,
) (TcpServer, error) {
	return TcpServer{
		connectionHandler: connectionHandler,
		logger:            tcpLogger,
		bind:              bind,
	}, nil
}

func (s TcpServer) Listen() {
	address := fmt.Sprintf("%s:%d", s.bind.Address, s.bind.Port)

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
