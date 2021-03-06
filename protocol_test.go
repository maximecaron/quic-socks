package socks

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPackHostData(t *testing.T) {
	testData := map[string][]byte{
		"1.1.1.1": {typeIPv4,
			0x01, 0x01, 0x01, 0x01,
			0x01, 0xBB}, // port
		"::1": {typeIPv6,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01,
			0x01, 0xBB},
		"google.com": {typeFQDN,
			0x0A, // FQDN size
			0x67, 0x6F, 0x6F, 0x67, 0x6C, 0x65, 0x2E, 0x63, 0x6F, 0x6D,
			0x01, 0xBB},
	}
	for host, expected := range testData {
		b, err := packHostData(host, 443)
		require.NoError(t, err)
		require.Equal(t, expected, b)
	}
}
