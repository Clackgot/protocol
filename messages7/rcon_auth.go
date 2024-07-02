package messages7

import (
	"slices"

	"github.com/teeworlds-go/protocol/chunk7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/packer"
)

type RconAuth struct {
	ChunkHeader *chunk7.ChunkHeader

	Password string
}

func (msg *RconAuth) MsgId() int {
	return network7.MsgSysRconAuth
}

func (msg *RconAuth) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *RconAuth) System() bool {
	return true
}

func (msg *RconAuth) Vital() bool {
	return true
}

func (msg *RconAuth) Pack() []byte {
	return slices.Concat(
		packer.PackStr(msg.Password),
	)
}

func (msg *RconAuth) Unpack(u *packer.Unpacker) error {
	msg.Password, _ = u.GetString()
	return nil
}

func (msg *RconAuth) Header() *chunk7.ChunkHeader {
	return msg.ChunkHeader
}

func (msg *RconAuth) SetHeader(header *chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
