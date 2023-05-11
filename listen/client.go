package listen

import (
	"io"
	"net"
)

type Client struct {
	net.Conn
}

func (c Client) Run(cmd string) ([]byte, error) {
	_, err := c.Read([]byte{})
	if err != nil {
		return nil, err
	}
	_, err = c.Write([]byte(cmd))
	if err != nil {
		return nil, err
	}
	return io.ReadAll(c)
}

func (c Client) Send(cmd string) error {
	_, err := c.Write([]byte(cmd))
	if err != nil {
		return err
	}
	return nil
}
