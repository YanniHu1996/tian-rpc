package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (HelloService) Hello(req string, reply *string) error {
	*reply = "hello " + req
	return nil
}

func main() {
	rpc.RegisterName("Say", HelloService{})

	listener, _ := net.Listen("tcp", ":1234")

	for {
		conn, _ := listener.Accept()
		rpc.ServeConn(conn)
	}
}
