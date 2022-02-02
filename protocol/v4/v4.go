package v4

import (
	"net"
)

type RequestChunk struct {
	SocksVersion    uint8
	CommandCode     uint8
	DestinationPort uint16
	DestinationIp   net.IP
	UserId          string
}

type ResponseChunk struct {
	SocksVersion    uint8
	CommandCode     uint8
	DestinationPort uint16
	DestinationIp   net.IP
}
