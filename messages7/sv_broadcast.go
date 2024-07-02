package messages7

import (
	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type SvBroadcast struct {
	ChunkHeader *chunk7.ChunkHeader

	Message string
}

func (msg *SvBroadcast) MsgId() int {
	return network7.MsgGameSvBroadcast
}

func (msg *SvBroadcast) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvBroadcast) System() bool {
	return false
}

func (msg *SvBroadcast) Vital() bool {
	return true
}

func (msg *SvBroadcast) Pack() []byte {
	return packer.PackStr(msg.Message)
}

func (msg *SvBroadcast) Unpack(u *packer.Unpacker) error {
	msg.Message, _ = u.GetString()
	return nil
}

func (msg *SvBroadcast) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *SvBroadcast) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
