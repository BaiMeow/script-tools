package ssh

import (
	"golang.org/x/crypto/ssh"
)

type Client struct {
	*ssh.Client
}

func (c *Client) Run(cmd string) ([]byte, error) {
	session, err := c.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	return session.CombinedOutput(cmd)
}

func (c *Client) Send(cmd string) error {
	session, err := c.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()
	return session.Run(cmd)
}
