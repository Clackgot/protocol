package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type RconLine struct {
	ChunkHeader *chunk7.ChunkHeader

	Line string
}

func (msg *RconLine) MsgId() int {
	return network7.MsgSysRconLine
}

func (msg *RconLine) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *RconLine) System() bool {
	return true
}

func (msg *RconLine) Vital() bool {
	return true
}

func (msg *RconLine) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Line),
	)
}

func (msg *RconLine) Unpack(u *packer.Unpacker) error {
	msg.Line, _ = u.GetString()
	return nil
}

func (msg *RconLine) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *RconLine) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
