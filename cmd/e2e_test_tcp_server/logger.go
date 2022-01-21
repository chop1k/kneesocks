package main

import "github.com/rs/zerolog"

type Logger struct {
	logger zerolog.Logger
}

func (l Logger) Listen(address string) {

}

func (l Logger) Connection(address string) {

}

func (l Logger) AcceptError(address string, err error) {

}

func (l Logger) PictureRequest(address string, picture byte) {

}

func (l Logger) Error(address string, err error) {

}

func (l Logger) DialError(address string, err error) {

}

func (l Logger) FileError(err error) {

}

func (l Logger) ResolveError(address string, err error) {

}

func (l Logger) IOError(address string, err error) {

}
