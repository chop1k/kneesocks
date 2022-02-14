package transfer

import (
	"net"
	"socks/config/tree"
	rate2 "socks/transfer/rate"
)

type ConnectHandler struct {
	handler Handler
}

func NewConnectHandler(handler Handler) (ConnectHandler, error) {
	return ConnectHandler{handler: handler}, nil
}

func (b ConnectHandler) HandleClient(rate tree.RateRestrictions, client net.Conn, host net.Conn) {
	limitedClient := rate2.NewBaseLimitedConn(rate.ClientWriteBuffersPerSecond, rate.ClientReadBuffersPerSecond, client)
	limitedHost := rate2.NewBaseLimitedConn(rate.HostWriteBuffersPerSecond, rate.HostReadBuffersPerSecond, host)

	go b.handler.TransferToHost(limitedClient, limitedHost)
	b.handler.TransferToClient(limitedClient, limitedHost)
}
