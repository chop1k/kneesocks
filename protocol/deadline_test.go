package protocol

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var testError = errors.New("Test error. ")

type ReaderMock struct {
	Wait int
	Err  error
}

func (r ReaderMock) Read(_ []byte) (n int, err error) {
	time.Sleep(time.Second * time.Duration(r.Wait))

	return 1, r.Err
}

func TestBaseDeadline_Read(t *testing.T) {
	manager, err := NewBaseDeadline()

	require.NoError(t, err)

	reader := ReaderMock{
		Wait: 1,
		Err:  nil,
	}

	i, readErr := manager.Read(2, 5, reader)

	require.NoError(t, readErr)

	require.Equal(t, []byte{0}, i)
}

func TestBaseDeadline_ReadReturnsTimeoutError(t *testing.T) {
	manager, err := NewBaseDeadline()

	require.NoError(t, err)

	reader := ReaderMock{
		Wait: 2,
		Err:  nil,
	}

	i, readErr := manager.Read(1, 5, reader)

	require.ErrorIs(t, TimeoutError, readErr)

	require.Nil(t, i)
}

func TestBaseDeadline_ReadReturnsError(t *testing.T) {
	manager, err := NewBaseDeadline()

	require.NoError(t, err)

	reader := ReaderMock{
		Wait: 0,
		Err:  testError,
	}

	i, readErr := manager.Read(2, 5, reader)

	require.ErrorIs(t, testError, readErr)

	require.Nil(t, i)
}

type WriterMock struct {
	Wait uint
	Err  error
}

func (w WriterMock) Write(_ []byte) (n int, err error) {
	time.Sleep(time.Second * time.Duration(w.Wait))

	return 0, w.Err
}

func TestBaseDeadline_Write(t *testing.T) {
	manager, err := NewBaseDeadline()

	require.NoError(t, err)

	writer := WriterMock{
		Wait: 0,
		Err:  nil,
	}

	writeErr := manager.Write(2, []byte{}, writer)

	require.NoError(t, writeErr)
}

func TestBaseDeadline_WriteReturnsTimeoutError(t *testing.T) {
	manager, err := NewBaseDeadline()

	require.NoError(t, err)

	writer := WriterMock{
		Wait: 2,
		Err:  nil,
	}

	writeErr := manager.Write(2, []byte{}, writer)

	require.ErrorIs(t, TimeoutError, writeErr)
}

func TestBaseDeadline_WriteReturnsError(t *testing.T) {
	manager, err := NewBaseDeadline()

	require.NoError(t, err)

	writer := WriterMock{
		Wait: 0,
		Err:  testError,
	}

	writeErr := manager.Write(2, []byte{}, writer)

	require.ErrorIs(t, testError, writeErr)
}
