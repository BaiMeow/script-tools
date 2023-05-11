package pwn

import (
	"fmt"
	"log"
	"net"
)

type Pwn struct {
	net.Conn
}

func NewRemote(addr string) net.Conn {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	return Pwn{conn}
}

func (p Pwn) SendLine(line []byte) {
	_, err := p.Conn.Write(line)
	if err != nil {
		log.Println(err)
		return
	}
}

func (p Pwn) RecLine() []byte {
	var buf []byte
	read := make([]byte, 1024)
	for {
		n, err := p.Conn.Read(read)
		if err != nil {
			log.Println(err)
			return nil
		}
		if n == 1024 {
			buf = append(buf, read...)
			read = make([]byte, len(read)*2)
			continue
		}
		break
	}
	return buf
}

func (p Pwn) after(flag []byte) {
	var abyte [1]byte
	_, err := p.Conn.Read(abyte[:])
	if err != nil {
		log.Println(err)
	}
	var i int
	for i = 0; i < len(flag); i++ {
		if abyte[0] == flag[i] {
			i++
		} else {
			i = 0
		}
	}
	return
}

func (p Pwn) RecAfter(after []byte) []byte {
	return p.RecLine()
}

func (p Pwn) SendAfter(after []byte, line []byte) {
	p.after(after)
	p.SendLine(line)
}

func (p Pwn) Interactive() {
	go func() {
		for {
			fmt.Println("Input: ", string(p.RecLine()))
		}
	}()
	go func() {
		for {
			var str string
			fmt.Scanln(&str)
			p.SendLine([]byte(str))
			fmt.Println("Send: ", str)
		}
	}()
}
