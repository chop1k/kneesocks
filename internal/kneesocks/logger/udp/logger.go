package udp

type Logger struct {
	Errors ErrorsLogger
	Listen ListenLogger
	Packet PacketLogger
}

func NewLogger(
	errors ErrorsLogger,
	listen ListenLogger,
	packet PacketLogger,
) Logger {
	return Logger{
		Errors: errors,
		Listen: listen,
		Packet: packet,
	}
}
