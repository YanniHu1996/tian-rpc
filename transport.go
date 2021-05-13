package trpc

import (
	"io"
	"log"
	"net"
	"sync"
)

type (
	Conn io.ReadWriteCloser

	Handler func(conn Conn)

	Server interface {
		Register(handler Handler)

		Start() error

		Stop() error
	}
)

type TcpSerer struct {
	Addr    string
	handler func(conn Conn)
	wg      sync.WaitGroup
}

func (t *TcpSerer) Register(f Handler) {
	t.handler = f
}

func (t *TcpSerer) Start() error {
	l, err := net.Listen("tcp", t.Addr)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return err
		}

		t.wg.Add(1)
		go func() {
			defer t.wg.Wait()
			t.handler(conn)
		}()
	}
}

func (t *TcpSerer) Stop() error {
	t.wg.Wait()
	return nil
}
