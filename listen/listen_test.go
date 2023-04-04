package listen

import (
	"net"
	"testing"
	"time"
)

func TestListener_ListenTCP(t *testing.T) {
	listener := Default()
	listener.CallBack = func(conn net.Conn) {
		_, _ = conn.Write([]byte("Hello, world!"))
		time.Sleep(time.Second * 1)
		rec := make([]byte, 1024)
		_, _ = conn.Read(rec)
		t.Log("rec:", string(rec), '\n')
	}
	err := listener.ListenTCP()
	if err != nil {
		return
	}
	time.Sleep(time.Second * 10)
	listener.Cancel()
	// test if listener closed
	_, err = net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   net.IP{127, 0, 0, 1},
		Port: 25001,
	})
	if err == nil {
		t.Log("conn not closed")
		t.Fail()
	}
}
