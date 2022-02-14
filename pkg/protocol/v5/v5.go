package v5

type MethodsChunk struct {
	SocksVersion byte
	Methods      []byte
}

type MethodSelectionChunk struct {
	SocksVersion byte
	Method       byte
}

type RequestChunk struct {
	SocksVersion byte
	CommandCode  byte
	AddressType  byte
	Address      string
	Port         uint16
}

type ResponseChunk struct {
	SocksVersion byte
	ReplyCode    byte
	AddressType  byte
	Address      string
	Port         uint16
}

type UdpRequest struct {
	Fragment    byte
	AddressType byte
	Address     string
	Port        uint16
	Data        []byte
}
