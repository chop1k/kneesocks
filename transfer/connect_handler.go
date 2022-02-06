package transfer

import (
	"net"
	"socks/config/tree"
	rate2 "socks/transfer/rate"
)

type ConnectHandler interface {
	HandleClient(rate tree.RateRestrictions, client net.Conn, host net.Conn)
}

type BaseConnectHandler struct {
	handler BaseHandler
}

func NewBaseConnectHandler(handler BaseHandler) (BaseConnectHandler, error) {
	return BaseConnectHandler{handler: handler}, nil
}

func (b BaseConnectHandler) HandleClient(rate tree.RateRestrictions, client net.Conn, host net.Conn) {
	limitedClient := rate2.NewBaseLimitedConn(rate.ClientWriteBuffersPerSecond, rate.ClientReadBuffersPerSecond, client)
	limitedHost := rate2.NewBaseLimitedConn(rate.HostWriteBuffersPerSecond, rate.HostReadBuffersPerSecond, host)

	go b.handler.TransferToHost(limitedClient, limitedHost)
	b.handler.TransferToClient(limitedClient, limitedHost)
}
