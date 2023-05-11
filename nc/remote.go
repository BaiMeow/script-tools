package nc

import "net"

func Remote(addr string) (net.Conn, error) {
	return net.Dial("tcp", addr)
}
