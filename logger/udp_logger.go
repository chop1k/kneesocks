package logger

type UdpLogger interface {
	PacketReceived()
	PacketDenied()
	PacketSent()
}
