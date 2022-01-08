package server

import (
	"fmt"
	"net"
	config2 "socks/config/tree"
	"socks/logger"
)

type Server struct {
	config            config2.Config
	connectionHandler ConnectionHandler
	packetHandler     PacketHandler
	tcpLogger         logger.TcpLogger
}

func NewServer(
	config config2.Config,
	connectionHandler ConnectionHandler,
	packetHandler PacketHandler,
	tcpLogger logger.TcpLogger,
) Server {
	return Server{
		config:            config,
		connectionHandler: connectionHandler,
		packetHandler:     packetHandler,
		tcpLogger:         tcpLogger,
	}
}

func (s Server) listenTcp() {
	addr := &net.TCPAddr{
		IP:   net.ParseIP(s.config.Tcp.BindIp),
		Port: int(s.config.Tcp.BindPort),
		Zone: s.config.Tcp.BindZone,
	}

	listener, err := net.ListenTCP("tcp", addr)

	if err != nil {
		panic(err)
	}

	go s.tcpLogger.Listen(addr.String())

	for {
		conn, acceptErr := listener.AcceptTCP()

		if acceptErr != nil {
			continue
		}

		//keepAliveErr := conn.SetKeepAlive(true)
		//
		//if keepAliveErr != nil {
		//	go s.errorHandler.HandleError(keepAliveErr)
		//
		//	continue
		//}

		/*writeBuffErr := conn.SetWriteBuffer(4090)

		if writeBuffErr != nil {
			go s.errorHandler.HandleError(writeBuffErr)

			continue
		}*/

		//readBuffErr := conn.SetReadBuffer(4090)
		//
		//if readBuffErr != nil {
		//	go s.errorHandler.HandleError(readBuffErr)
		//
		//	continue
		//}

		remoteAddr := conn.RemoteAddr()

		go s.tcpLogger.ConnectionAccepted(remoteAddr.String())

		go s.connectionHandler.HandleConnection(conn)
	}
}

func (s Server) listenUdp() {
	addr := fmt.Sprintf("%s:%d", s.config.Udp.BindIp, s.config.Udp.BindPort)

	packet, err := net.ListenPacket("udp", addr)

	if err != nil {
		panic(packet)
	}

	for {
		payload := make([]byte, 65535)

		i, address, readErr := packet.ReadFrom(payload)

		if readErr != nil {
			panic(readErr)
		}

		go s.packetHandler.HandlePacket(payload[:i], address, packet)
	}
}

func (s Server) listenUnix() {
	select {}
}

func (s Server) Start() {
	go s.listenTcp()
	go s.listenUdp()

	s.listenUnix()
}
