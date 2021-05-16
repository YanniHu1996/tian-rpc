package protocol

import "encoding/binary"

const headerSize = 13

type Header [headerSize]byte

// MessageType is message type of requests and responses.
type MessageType byte

const (
	// Request is message type of request
	Request MessageType = iota
	// Response is message type of response
	Response
)

// SerializeType defines serialization type of payload.
type SerializeType byte

const (
	// SerializeNone uses raw []byte and don't serialize/deserialize
	SerializeNone SerializeType = iota
	// JSON for payload.
	JSON
	// ProtoBuffer for payload.
	ProtoBuffer
)

// CheckMagicNumber  0~15魔术位，
func (h Header) CheckMagicNumber() bool {
	return h[0] == magicNumber
}

// PayloadSize  16~47整体长度
func (h Header) PayloadSize() uint32 {
	return binary.BigEndian.Uint32(h[2:6])
}
func (h Header) SetPayloadSize(size uint32) {
	binary.BigEndian.PutUint32(h[2:6], size)
}

// MetadataSize 48~63头长度
func (h Header) MetadataSize() uint32 {
	return binary.BigEndian.Uint32(h[6:9])
}

func (h Header) SetMetadataSize(size uint32) {
	binary.BigEndian.PutUint32(h[6:9], size)
}

// Version 64~71 协议版本，
func (h Header) Version() byte {
	return h[9]
}

func (h *Header) SetVersion(v byte) {
	h[9] = v
}

// MessageType 72~79 消息类型，
func (h Header) MessageType() MessageType {
	return MessageType(h[9])
}

func (h *Header) SetMessageType(mt MessageType) {
	h[9] = byte(mt)
}

// SerializeType 80~87 序列化方式
func (h Header) SerializeType() SerializeType {
	return SerializeType(h[10])
}

func (h *Header) SetSerializeType(st SerializeType) {
	h[10] = byte(st)
}

// Seq 88~103 消息ID
func (h Header) Seq() uint64 {
	return binary.BigEndian.Uint64(h[11:13])
}

func (h *Header) SetSeq(seq uint64) {
	binary.BigEndian.PutUint64(h[11:13], seq)
}

var zeroHeaderArray Header
var zeroHeader = zeroHeaderArray[1:]

func resetHeader(h *Header) {
	copy(h[1:], zeroHeader)
}
