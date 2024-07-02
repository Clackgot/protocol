package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type RconCmdRem struct {
	ChunkHeader *chunk7.ChunkHeader

	Name string
}

func (msg *RconCmdRem) MsgId() int {
	return network7.MsgSysRconCmdRem
}

func (msg *RconCmdRem) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *RconCmdRem) System() bool {
	return true
}

func (msg *RconCmdRem) Vital() bool {
	return true
}

func (msg *RconCmdRem) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Name),
	)
}

func (msg *RconCmdRem) Unpack(u *packer.Unpacker) error {
	msg.Name, _ = u.GetString()
	return nil
}

func (msg *RconCmdRem) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *RconCmdRem) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
