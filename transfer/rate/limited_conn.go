package rate

import (
	"errors"
	"net"
	"time"
)

var (
	ImpossibleError = errors.New("Now < previously saved now. ")
)

type BaseLimitedConn struct {
	writesPerSecond int
	readsPerSecond  int
	lastRead        int64
	readsCollected  int
	lastWrite       int64
	writesCollected int
	conn            net.Conn
}

func NewBaseLimitedConn(writesPerSecond int, readsPerSecond int, conn net.Conn) *BaseLimitedConn {
	now := time.Now().UTC()

	return &BaseLimitedConn{
		writesPerSecond: writesPerSecond,
		readsPerSecond:  readsPerSecond,
		lastRead:        now.Unix(),
		readsCollected:  0,
		lastWrite:       now.Unix(),
		writesCollected: 0,
		conn:            conn,
	}
}

func (b BaseLimitedConn) LocalAddr() net.Addr {
	return b.conn.LocalAddr()
}

func (b BaseLimitedConn) RemoteAddr() net.Addr {
	return b.conn.RemoteAddr()
}

func (b BaseLimitedConn) SetDeadline(t time.Time) error {
	return b.conn.SetDeadline(t)
}

func (b BaseLimitedConn) SetReadDeadline(t time.Time) error {
	return b.conn.SetReadDeadline(t)
}

func (b BaseLimitedConn) SetWriteDeadline(t time.Time) error {
	return b.conn.SetWriteDeadline(t)
}

func (b *BaseLimitedConn) Write(p []byte) (n int, err error) {
	if b.writesPerSecond <= 0 {
		return b.conn.Write(p)
	}

	now := time.Now().UTC()

	unix := now.Unix()

	if unix > b.lastWrite {

		b.lastWrite = unix
		b.writesCollected = 0

		return b.conn.Write(p)
	}

	if unix == b.lastWrite {
		if b.writesCollected >= b.writesPerSecond {
			nextSecond := time.Date(
				now.Year(),
				now.Month(),
				now.Day(),
				now.Hour(),
				now.Minute(),
				now.Second()+1,
				0,
				time.UTC,
			)

			time.Sleep(nextSecond.Sub(now))

			b.lastWrite = nextSecond.Unix()
			b.writesCollected = 0

			return b.conn.Write(p)
		}

		b.lastWrite = unix
		b.writesCollected = b.writesCollected + 1

		return b.conn.Write(p)
	}

	return 0, ImpossibleError
}

func (b *BaseLimitedConn) Read(p []byte) (n int, err error) {
	if b.readsPerSecond <= 0 {
		return b.conn.Read(p)
	}

	now := time.Now().UTC()

	unix := now.Unix()

	if unix > b.lastRead {
		b.lastRead = unix
		b.readsCollected = 0

		return b.conn.Read(p)
	}

	if unix == b.lastRead {
		if b.readsCollected >= b.readsPerSecond {
			nextSecond := time.Date(
				now.Year(),
				now.Month(),
				now.Day(),
				now.Hour(),
				now.Minute(),
				now.Second()+1,
				0,
				time.UTC,
			)

			time.Sleep(nextSecond.Sub(now))

			b.lastRead = nextSecond.Unix()
			b.readsCollected = 0

			return b.conn.Read(p)
		}

		b.lastRead = unix
		b.readsCollected++

		return b.conn.Read(p)
	}

	return 0, ImpossibleError
}

func (b BaseLimitedConn) Close() error {
	return b.conn.Close()
}
