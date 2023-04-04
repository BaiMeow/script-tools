// Package listen is a package that provides a listener which accept many connections in one port.
package listen

import (
	"context"
	"log"
	"net"
	"strconv"
)

type Listener struct {
	Port     int
	CallBack func(conn net.Conn)
	ctx      context.Context
	cancel   context.CancelFunc
}

// Default returns a Listener listening on port 25001
func Default() *Listener {
	ctx, cancel := context.WithCancel(context.Background())
	return &Listener{
		Port:     25001,
		CallBack: nil,
		ctx:      ctx,
		cancel:   cancel,
	}
}

func (c *Listener) ListenTCP() error {
	cfg := &net.ListenConfig{}
	listener, err := cfg.Listen(c.ctx, "tcp", ":"+strconv.Itoa(c.Port))
	if err != nil {
		return err
	}
	go func() {
		defer listener.Close()
		for {
			select {
			case <-c.ctx.Done():
				return
			default:
				conn, err := listener.Accept()
				if err != nil {
					log.Println(err)
					continue
				}
				log.Println("New connection from: ", conn.RemoteAddr().String())
				go c.handleConnection(conn)
			}
		}
	}()
	return nil
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

func (c *Listener) Cancel() {
	if c.cancel != nil {
		c.cancel()
	}
}
