package demo

import "net/rpc"

type HelloSrv struct {
}

func (h HelloSrv) Hello(s *String, s2 *String) error {
	s2.Value = "hello" + s2.GetValue()
	return nil
}

func main() {

	RegisterHelloService(rpc.DefaultServer, HelloSrv{})
}
