package trpc

type Request struct {
	ServiceMethod string // format: "Service.Method"
	Seq           uint64 // sequence number chosen by client
}

type Response struct {
	ServiceMethod string // echoes that of the Request
	Seq           uint64 // echoes that of the request
	Error         string // error, if any.
}

type ServerCodec interface {
	ReadRequestHeader(*Request) error
	ReadRequestBody(interface{}) error
	WriteResponse(*Response, interface{}) error

	// Close can be called multiple times and must be idempotent.
	Close() error
}
