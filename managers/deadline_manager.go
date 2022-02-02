package managers

import (
	"errors"
	"io"
	"time"
)

var (
	ReadChannelClosedError  = errors.New("Read channel is closed. ")
	WriteChannelClosedError = errors.New("Write channel is closed. ")
	ErrChannelClosedError   = errors.New("Err channel is closed. ")
	TimerChannelClosedError = errors.New("Timer channel is closed. ")
	TimeoutError            = errors.New("Timeout exceeded. ")
)

type DeadlineManager interface {
	Read(deadline uint, bufferLength int, reader io.Reader) ([]byte, error)
	Write(deadline uint, data []byte, writer io.Writer) error
}

type BaseDeadlineManager struct {
}

func NewBaseDeadlineManager() (BaseDeadlineManager, error) {
	return BaseDeadlineManager{}, nil
}

func (b BaseDeadlineManager) Read(deadline uint, bufferLength int, reader io.Reader) ([]byte, error) {
	closed := false

	read := make(chan []byte, 1)
	err := make(chan error, 1)

	timer := time.NewTimer(time.Second * time.Duration(deadline))

	go func() {
		buffer := make([]byte, bufferLength)

		i, readErr := reader.Read(buffer)

		if closed {
			return
		}

		if readErr != nil {
			err <- readErr

			return
		}

		read <- buffer[:i]
	}()

	select {
	case data, ok := <-read:
		closed = true

		b.readCleanUp(read, err, timer)

		if !ok {
			return nil, ReadChannelClosedError
		}

		return data, nil
	case _, ok := <-timer.C:
		closed = true

		b.readCleanUp(read, err, timer)

		if !ok {
			return nil, TimerChannelClosedError
		}

		return nil, TimeoutError
	case readErr, ok := <-err:
		closed = true

		b.readCleanUp(read, err, timer)

		if !ok {
			return nil, ErrChannelClosedError
		}

		return nil, readErr
	}
}

func (b BaseDeadlineManager) readCleanUp(read chan []byte, err chan error, timer *time.Timer) {
	close(read)
	close(err)

	timer.Stop()
}

func (b BaseDeadlineManager) Write(deadline uint, data []byte, writer io.Writer) error {
	closed := false

	done := make(chan bool, 1)
	err := make(chan error, 1)

	timer := time.NewTimer(time.Second * time.Duration(deadline))

	go func() {
		_, writeErr := writer.Write(data)

		if closed {
			return
		}

		if writeErr != nil {
			err <- writeErr

			return
		}

		done <- true
	}()

	select {
	case _, ok := <-done:
		closed = true

		b.writeCleanUp(done, err, timer)

		if !ok {
			return WriteChannelClosedError
		}

		return nil
	case _, ok := <-timer.C:
		closed = true

		b.writeCleanUp(done, err, timer)

		if !ok {
			return TimerChannelClosedError
		}

		return TimeoutError
	case readErr, ok := <-err:
		closed = true

		b.writeCleanUp(done, err, timer)

		if !ok {
			return ErrChannelClosedError
		}

		return readErr
	}
}

func (b BaseDeadlineManager) writeCleanUp(done chan bool, err chan error, timer *time.Timer) {
	close(done)
	close(err)

	timer.Stop()
}
