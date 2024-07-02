package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type RconCmd struct {
	ChunkHeader *chunk7.ChunkHeader

	Command string
}

func (msg *RconCmd) MsgId() int {
	return network7.MsgSysRconCmd
}

func (msg *RconCmd) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *RconCmd) System() bool {
	return true
}

func (msg *RconCmd) Vital() bool {
	return true
}

func (msg *RconCmd) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Command),
	)
}

func (msg *RconCmd) Unpack(u *packer.Unpacker) error {
	msg.Command, _ = u.GetString()
	return nil
}

func (msg *RconCmd) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *RconCmd) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
