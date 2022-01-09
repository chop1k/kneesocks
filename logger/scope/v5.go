package scope

import (
	v5 "socks/protocol/v5"
	"strconv"
	"time"
)

type SocksV5Scope struct {
	event      string
	parameters map[string]string
	logs       map[string]map[string]string
}

func NewSocksV5Scope() SocksV5Scope {
	return SocksV5Scope{
		parameters: map[string]string{},
	}
}

func (s SocksV5Scope) setTime() {
	now := time.Now()

	s.parameters["now_unix"] = strconv.FormatInt(now.Unix(), 10)
	s.parameters["now"] = now.Format("2006-01-02 15:04:05")
	s.parameters["now_date"] = now.Format("2006-01-02")
	s.parameters["now_time"] = now.Format("15:04:05")
}

func (s SocksV5Scope) setChunk(chunk v5.RequestChunk) {
	s.parameters["chunk.CommandCode"] = string(chunk.CommandCode)
	s.parameters["chunk.Address"] = chunk.Address
	s.parameters["chunk.Port"] = strconv.Itoa(int(chunk.Port))
	s.parameters["chunk.SocksVersion"] = string(chunk.SocksVersion)
}

func (s SocksV5Scope) setClientAddress(client string) {
	s.parameters["client"] = client
}

func (s SocksV5Scope) setHostAddress(host string) {
	s.parameters["host"] = host
}
