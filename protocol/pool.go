package protocol

import "sync"

var msgPool = sync.Pool{
	New: func() interface{} {
		header := Header([headerSize]byte{})
		header[0] = magicNumber

		return &Message{
			Header: &header,
		}
	},
}

// GetPooledMsg gets a pooled message.
func GetPooledMsg() *Message {
	return msgPool.Get().(*Message)
}

// FreeMsg puts a msg into the pool.
func FreeMsg(msg *Message) {
	if msg != nil {
		msg.Reset()
		msgPool.Put(msg)
	}
}
