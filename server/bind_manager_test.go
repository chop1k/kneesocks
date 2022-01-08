package server

import (
	"github.com/stretchr/testify/require"
	"net"
	"testing"
)

func TestNewBindManager(t *testing.T) {
	bindManager := NewBindManager()

	require.Zero(t, len(bindManager.addresses), 1, "BindManager constructor returned instance with non-empty map. ")
}

func TestBindManager_IsBound(t *testing.T) {
	bindManager := NewBindManager()

	require.False(t, bindManager.IsBound("test"))

	bindManager.addresses["test"] = bundle{}

	require.True(t, bindManager.IsBound("test"))
}

func TestBindManager_Bind(t *testing.T) {
	bindManager := NewBindManager()

	_, ok := bindManager.addresses["test"]

	require.False(t, ok)

	bindManager.Bind("test")

	_, ok = bindManager.addresses["test"]

	require.True(t, ok)
}

func TestBindManager_Remove(t *testing.T) {
	bindManager := NewBindManager()

	bindManager.addresses["test"] = bundle{
		client: make(chan net.Conn),
		host:   make(chan net.Conn),
	}

	bindManager.Remove("test")

	_, ok := bindManager.addresses["test"]

	require.False(t, ok)
}

func TestBindManager_RemoveClosesChannels(t *testing.T) {
	bindManager := NewBindManager()

	clientChan := make(chan net.Conn)
	hostChan := make(chan net.Conn)

	bindManager.addresses["test"] = bundle{
		client: clientChan,
		host:   hostChan,
	}

	bindManager.Remove("test")

	_, ok := <-clientChan

	require.False(t, ok)

	_, ok = <-hostChan

	require.False(t, ok)
}

func TestBindManager_SendHost(t *testing.T) {
	bindManager := NewBindManager()

	clientChan := make(chan net.Conn)
	hostChan := make(chan net.Conn)

	bindManager.addresses["test"] = bundle{
		client: clientChan,
		host:   hostChan,
	}

	go func() {
		_ = bindManager.SendHost("test", nil)
	}()

	conn := <-hostChan

	require.Nil(t, conn)
}

func TestBindManager_SendHostReturnsNotExistsError(t *testing.T) {
	bindManager := NewBindManager()

	_, ok := bindManager.addresses["test"]

	require.False(t, ok)

	err := bindManager.SendHost("test", nil)

	require.ErrorIs(t, err, AddressNotExistsError)
}

func TestBindManager_SendClient(t *testing.T) {
	bindManager := NewBindManager()

	clientChan := make(chan net.Conn)
	hostChan := make(chan net.Conn)

	bindManager.addresses["test"] = bundle{
		client: clientChan,
		host:   hostChan,
	}

	go func() {
		_ = bindManager.SendClient("test", nil)
	}()

	conn := <-clientChan

	require.Nil(t, conn)
}

func TestBindManager_SendClientReturnsNotExistsError(t *testing.T) {
	bindManager := NewBindManager()

	_, ok := bindManager.addresses["test"]

	require.False(t, ok)

	err := bindManager.SendClient("test", nil)

	require.ErrorIs(t, err, AddressNotExistsError)
}

func TestBindManager_ReceiveClient(t *testing.T) {
	bindManager := NewBindManager()

	clientChan := make(chan net.Conn)
	hostChan := make(chan net.Conn)

	bindManager.addresses["test"] = bundle{
		client: clientChan,
		host:   hostChan,
	}

	go func() {
		clientChan <- nil
	}()

	client, err := bindManager.ReceiveClient("test")

	require.NoError(t, err)

	require.Nil(t, client)
}

func TestBindManager_ReceiveClientReturnsNotExistsError(t *testing.T) {
	bindManager := NewBindManager()

	_, ok := bindManager.addresses["test"]

	require.False(t, ok)

	_, err := bindManager.ReceiveClient("test")

	require.ErrorIs(t, err, AddressNotExistsError)
}

func TestBindManager_ReceiveClientReturnsClosedChannelError(t *testing.T) {
	bindManager := NewBindManager()

	clientChan := make(chan net.Conn)
	hostChan := make(chan net.Conn)

	bindManager.addresses["test"] = bundle{
		client: clientChan,
		host:   hostChan,
	}

	close(clientChan)
	close(hostChan)

	_, err := bindManager.ReceiveClient("test")

	require.ErrorIs(t, err, ClientChannelClosedError)
}

func TestBindManager_ReceiveHost(t *testing.T) {
	bindManager := NewBindManager()

	clientChan := make(chan net.Conn)
	hostChan := make(chan net.Conn)

	bindManager.addresses["test"] = bundle{
		client: clientChan,
		host:   hostChan,
	}

	go func() {
		hostChan <- nil
	}()

	host, err := bindManager.ReceiveHost("test")

	require.NoError(t, err)

	require.Nil(t, host)
}

func TestBindManager_ReceiveHostReturnsNotExistsError(t *testing.T) {
	bindManager := NewBindManager()

	_, ok := bindManager.addresses["test"]

	require.False(t, ok)

	_, err := bindManager.ReceiveHost("test")

	require.ErrorIs(t, err, AddressNotExistsError)
}

func TestBindManager_ReceiveHostReturnsClosedChannelError(t *testing.T) {
	bindManager := NewBindManager()

	clientChan := make(chan net.Conn)
	hostChan := make(chan net.Conn)

	bindManager.addresses["test"] = bundle{
		client: clientChan,
		host:   hostChan,
	}

	close(clientChan)
	close(hostChan)

	_, err := bindManager.ReceiveHost("test")

	require.ErrorIs(t, err, HostChannelClosedError)
}
