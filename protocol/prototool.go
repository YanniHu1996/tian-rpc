package protocol

import (
	"fmt"
	"io"
)

const (
	magicNumber byte = 0x11
)

func MagicNumber() byte {
	return magicNumber
}

type Message struct {
	*Header
	Metadata
	Payload []byte
	data    []byte

	//ServicePath   string
	//ServiceMethod string
}

func (m *Message) Encode() ([]byte, error) {
	// 协议头

	// 元数据

	// 负载
}
func (m *Message) Decode(r io.Reader) error {
	// 解析协议头
	_, err := io.ReadFull(r, m.Header[:1])
	if err != nil {
		return err
	}
	if !m.Header.CheckMagicNumber() {
		return fmt.Errorf("wrong magic number: %v", m.Header[0])
	}

	// 解析元数据
	if l := m.Header.MetadataSize(); l > 0 {
		metadataData := make([]byte, m.Header.MetadataSize())
		m.Metadata, err = decodeMetadata(m.Header.MetadataSize(), metadataData)
		if err != nil {
			return err
		}
	}

	// 解析负载
	m.Payload = make([]byte, m.Header.PayloadSize())
	_, err = io.ReadFull(r, m.Payload)
	if err != nil {
		return err
	}
	return nil
}

// Reset clean data of this message but keep allocated data
func (m *Message) Reset() {
	resetHeader(m.Header)
	m.Metadata = nil
	m.Payload = []byte{}
	m.data = m.data[:0]
}
