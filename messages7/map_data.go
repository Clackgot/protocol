package messages7

import (
	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type MapData struct {
	ChunkHeader *chunk7.ChunkHeader

	Data []byte
}

func (msg *MapData) MsgId() int {
	return network7.MsgSysMapData
}

func (msg *MapData) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *MapData) System() bool {
	return true
}

func (msg *MapData) Vital() bool {
	return true
}

func (msg *MapData) Pack() []byte {
	return msg.Data
}

func (msg *MapData) Unpack(u *packer.Unpacker) error {
	msg.Data = u.Rest()
	return nil
}

func (msg *MapData) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *MapData) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
