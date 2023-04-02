// Package listen is a package that provides a listener which accept many connections in one port.
package listen

import (
	"log"
	"net"
)

type Listener struct {
	Port     int
	BuffLen  int
	CallBack func(conn net.Conn)
}

// Default returns a Listener listening on port 25001
func Default() *Listener {
	return &Listener{
		Port:     25001,
		BuffLen:  1024,
		CallBack: nil,
	}
}

func (c *Listener) ListenTCP() error {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP:   net.IP{0, 0, 0, 0},
		Port: c.Port,
	})
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		log.Println("New connection from: ", conn.RemoteAddr().String())
		if err != nil {
			log.Println(err)
		}
		go c.handleConnection(conn)
	}
}

func (c *Listener) handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		_ = conn.Close()
		log.Println("Connection closed: ", conn.RemoteAddr().String())
	}(conn)
	if c.CallBack != nil {
		c.CallBack(conn)
	}
}
