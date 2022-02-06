package transfer

import (
	"net"
	"socks/managers"
	"socks/transfer/rate"
)

type BindHandler interface {
	HandleClient(client net.Conn, host net.Conn)
	HandleHost(client net.Conn, host net.Conn)
}

type BaseBindHandler struct {
	bindRate managers.BindRateManager
	handler  BaseHandler
}

func NewBaseBindHandler(
	bindRate managers.BindRateManager,
	handler BaseHandler,
) (BaseBindHandler, error) {
	return BaseBindHandler{
		bindRate: bindRate,
		handler:  handler,
	}, nil
}

func (b BaseBindHandler) HandleClient(client net.Conn, host net.Conn) {
	restrictions, err := b.bindRate.Get(client.RemoteAddr().String())

	if err != nil {
		_ = client.Close()
		_ = host.Close()

		return
	}

	limitedClient := rate.NewBaseLimitedConn(restrictions.ClientWriteBuffersPerSecond, restrictions.ClientReadBuffersPerSecond, client)
	limitedHost := rate.NewBaseLimitedConn(restrictions.HostWriteBuffersPerSecond, restrictions.HostReadBuffersPerSecond, host)

	b.handler.TransferToHost(limitedClient, limitedHost)
}

func (b BaseBindHandler) HandleHost(client net.Conn, host net.Conn) {
	restrictions, err := b.bindRate.Get(client.RemoteAddr().String())

	if err != nil {
		_ = client.Close()
		_ = host.Close()

		return
	}

	limitedClient := rate.NewBaseLimitedConn(restrictions.ClientWriteBuffersPerSecond, restrictions.ClientReadBuffersPerSecond, client)
	limitedHost := rate.NewBaseLimitedConn(restrictions.HostWriteBuffersPerSecond, restrictions.HostReadBuffersPerSecond, host)

	b.handler.TransferToClient(limitedClient, limitedHost)
}
