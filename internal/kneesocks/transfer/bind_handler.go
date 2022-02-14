package transfer

import (
	"net"
	"socks/internal/kneesocks/managers"
	"socks/internal/kneesocks/transfer/rate"
)

type BindHandler struct {
	bindRate managers.BindRateManager
	handler  Handler
}

func NewBindHandler(
	bindRate managers.BindRateManager,
	handler Handler,
) (BindHandler, error) {
	return BindHandler{
		bindRate: bindRate,
		handler:  handler,
	}, nil
}

func (b BindHandler) HandleClient(client net.Conn, host net.Conn) {
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

func (b BindHandler) HandleHost(client net.Conn, host net.Conn) {
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
