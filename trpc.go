package trpc

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

type option struct {
	server Server
}

var defaultOption = &option{
	server: &TcpSerer{Addr: ":1234"},
}

type optionFunc func(*option)

type TRpc struct {
	option *option
	Server Server
}

func NewTRpc(optionFuncs ...optionFunc) *TRpc {
	for _, f := range optionFuncs {
		f(defaultOption)
	}

	r := &TRpc{
		option: defaultOption,
		Server: defaultOption.server,
	}

	defaultOption.server.Register(r.Handler)

	return r
}

func (t *TRpc) Handler(conn Conn) {
	fmt.Println(conn.RemoteAddr().String())
	// TODO
	for {
		data := [3]byte{}

		rd := bufio.NewReader(conn)
		l, err := io.ReadFull(rd, data[:])
		l, err := conn.Read(data[:])
		if err != nil {
			log.Println(err)
			return
		}
		if l != 0 {
			fmt.Printf("size: %v,data: %v\n", l, string(data[:]))
		}
	}
}

func (t *TRpc) Start() error {
	return t.Server.Start()
}
