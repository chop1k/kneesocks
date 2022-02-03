package v4a

import (
	"net"
)

type RequestChunk struct {
	SocksVersion    byte
	CommandCode     byte
	DestinationPort uint16
	DestinationIp   net.IP
	Domain          string
}

type ResponseChunk struct {
	SocksVersion    byte
	CommandCode     byte
	DestinationPort uint16
	DestinationIp   net.IP
}
