package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type RequestChunk struct {
	Picture     byte
	AddressType byte
	Address     net.IP
	Port        uint16
}

type RequestHandler struct {
	config Config
	logger Logger
	sender PictureSender
}

func NewRequestHandler(config Config, logger Logger, sender PictureSender) (RequestHandler, error) {
	return RequestHandler{config: config, logger: logger, sender: sender}, nil
}

func (h RequestHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {
	chunk := RequestChunk{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&chunk)

	if err != nil {
		h.sendBadRequest(w)

		return
	}

	h.logger.Request(r.RemoteAddr, chunk.Picture)

	h.resolveLAddr(chunk.AddressType, chunk.Address, chunk.Port, chunk.Picture, w, r.RemoteAddr)
}

func (h RequestHandler) sendBadRequest(w http.ResponseWriter) {
	w.WriteHeader(400)
}

func (h RequestHandler) sendError(w http.ResponseWriter) {
	w.WriteHeader(500)
}

func (h RequestHandler) sendOk(w http.ResponseWriter) {
	w.WriteHeader(200)
}

func (h RequestHandler) resolveLAddr(addressType byte, ip net.IP, port uint16, picture byte, w http.ResponseWriter, address string) {
	var selfAddress string

	if addressType == 4 {
		selfAddress = fmt.Sprintf("[%s]:%d", ip, port)
	} else {
		selfAddress = fmt.Sprintf("%s:%d", ip, port)
	}

	lAddr, err := net.ResolveTCPAddr("tcp", selfAddress)

	if err != nil {
		h.logger.ResolveError(address, err)

		h.sendError(w)

		return
	}

	h.resolveRAddr(addressType, lAddr, picture, w, address)
}

func (h RequestHandler) resolveRAddr(addressType byte, lAddr *net.TCPAddr, picture byte, w http.ResponseWriter, address string) {
	var socksAddress string

	if addressType == 4 {
		socksAddress = fmt.Sprintf("[%s]:%d", h.config.Socks.IPv6, h.config.Socks.Port)
	} else {
		socksAddress = fmt.Sprintf("%s:%d", h.config.Socks.IPv4, h.config.Socks.Port)
	}

	rAddr, err := net.ResolveTCPAddr("tcp", socksAddress)

	if err != nil {
		h.logger.ResolveError(address, err)

		h.sendError(w)

		return
	}

	h.bind(lAddr, rAddr, picture, w, address)
}

func (h RequestHandler) bind(lAddr *net.TCPAddr, rAddr *net.TCPAddr, picture byte, w http.ResponseWriter, address string) {
	host, dialErr := net.DialTCP("tcp", lAddr, rAddr)

	if dialErr != nil {
		h.logger.DialError(address, dialErr)

		h.sendError(w)

		return
	}

	err := h.sender.Send(address, picture, host)

	if err != nil {
		h.sendError(w)

		return
	}

	h.sendOk(w)
}
