package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewUtils(t *testing.T) {
}

func TestAddressUtils_DetermineAddressType(t *testing.T) {
	tests := []struct {
		address string
		result  byte
		err     error
	}{
		{
			"1.2.3.4",
			1,
			nil,
		},
		{
			"1.2.3.4.5",
			0,
			CannotParseIpError,
		},
		{
			"ff02::2:114",
			4,
			nil,
		},
		{
			"ff02::2:114::443",
			0,
			CannotParseIpError,
		},
		{
			"[::1]",
			0,
			CannotParseIpError,
		},
	}

	addrUtils, err := NewUtils()

	require.NoError(t, err)

	for i, test := range tests {
		result, err := addrUtils.DetermineAddressType(test.address)

		require.ErrorIsf(t, err, test.err, "Error not equals (%d). ", i)

		require.Equalf(t, result, test.result, "Results not equals (%d). ", i)
	}
}

func TestAddressUtils_ParseAddress(t *testing.T) {
	tests := []struct {
		address string
		host    string
		port    int
		err     error
	}{
		{
			"1.2.3.4:443",
			"1.2.3.4",
			443,
			nil,
		},
		{
			"en.wikipedia.org:443",
			"en.wikipedia.org",
			443,
			nil,
		},
		{
			"[ff02::2:114]:443",
			"ff02::2:114",
			443,
			nil,
		},
		{
			"",
			"",
			0,
			CannotParseAddressError,
		},
	}

	addrUtils, err := NewUtils()

	require.NoError(t, err)

	for i, test := range tests {
		host, port, err := addrUtils.ParseAddress(test.address)

		require.ErrorIsf(t, err, test.err, "Error not equals (%d). ", i)

		require.Equalf(t, host, test.host, "Hosts not equals (%d). ", i)
		require.Equalf(t, port, test.port, "Ports not equals (%d). ", i)
	}
}

func TestAddressUtils_ConvertAddress(t *testing.T) {
	tests := []struct {
		addrType byte
		bytes    []byte
		address  string
		err      error
	}{
		{
			1,
			[]byte{1, 2, 3, 4},
			"1.2.3.4",
			nil,
		},
		{
			3,
			[]byte{114, 117, 46, 119, 105, 107, 105, 112, 101, 100, 105, 97, 46, 111, 114, 103},
			"ru.wikipedia.org",
			nil,
		},
		{
			4,
			[]byte{255, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 1, 20},
			"ff02::2:114",
			nil,
		},
	}

	addrUtils, err := NewUtils()

	require.NoError(t, err)

	for i, test := range tests {
		ip, err := addrUtils.ConvertAddress(test.addrType, test.bytes)

		require.ErrorIsf(t, err, test.err, "Error not equals (%d). ", i)

		require.Equalf(t, ip, test.address, "Ip not equals (%d). ", i)
	}
}
