package socks

import (
	"context"
	"crypto/tls"
	"net"

	"github.com/lucas-clemente/quic-go"
	"github.com/pkg/errors"
)

type Client struct {
	address   string
	password  []byte
	tlsConfig *tls.Config
	session   quic.Session
}

func NewClient(address string, tlsConfig *tls.Config, password string) (*Client, error) {
	c := Client{
		address:   address,
		password:  []byte(password),
		tlsConfig: tlsConfig,
	}
	if len(c.password) > 32 {
		return nil, errors.New("password size > 32")
	}
	err := c.dial()
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Client) dial() error {
	var err error
	c.session, err = quic.DialAddr(c.address, c.tlsConfig, nil)
	return err
}

// Connect
func (c *Client) Connect(host string, port uint16) (net.Conn, error) {
	stream, err := c.session.OpenStreamSync(context.Background())
	if err != nil {
		if ne, ok := err.(net.Error); ok && ne.Timeout() {
			// reconnect
			err = c.dial()
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	hostData, err := packHostData(host, port)
	if err != nil {
		return nil, err
	}
	conn := newConn(c.session, stream)

	// send request
	_, err = conn.Write(append(c.password, hostData...))
	if err != nil {
		return nil, err
	}

	// receive response
	resp := make([]byte, 1)
	_, err = conn.Read(resp)
	if err != nil {
		return nil, err
	}
	if resp[0] != respOK {
		return nil, Response(resp[0])
	}
	return conn, nil
}

func (c *Client) Close() {
	_ = c.session.Close()
}