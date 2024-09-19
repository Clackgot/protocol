package snapshot7_test

import (
	"testing"

	"github.com/teeworlds-go/protocol/internal/testutils/require"
	"github.com/teeworlds-go/protocol/messages7"
	"github.com/teeworlds-go/protocol/network7"
	"github.com/teeworlds-go/protocol/protocol7"
	"github.com/teeworlds-go/protocol/teeworlds7"
)

// go client connected to public ~40 player ddnet Linear server
// the ddnet server for some reason sent part 1 before part 0

func Test2PartSnap(t *testing.T) {
	t.Parallel()
	// snapshot captured with tcpdump
	// libtw2 dissector details
	//
	// Frame 23: 576 bytes on wire (4608 bits), 576 bytes captured (4608 bits)
	// Linux cooked capture v2
	// Internet Protocol Version 4
	// User Datagram Protocol, Src Port: 8304, Dst Port: 44074
	// Teeworlds 0.7 Protocol packet
	// Teeworlds 0.7 Protocol chunk: sys.snap
	//     Header (non-vital)
	//     Message: sys.snap
	//     Tick: 12847240
	//     Delta tick: 12847241
	//     Num parts: 2
	//     Part: 1
	//     Crc: -2136653645
	//     Data (605 bytes)
	dumpPart1 := []byte{
		0x10, 0x04, 0x01, 0x01, 0x02, 0x03, 0x04,
		0x4c, 0xcb, 0x1c, 0x2d, 0x3e, 0xf3, 0x93, 0xfb, 0xc8,
		0x67, 0x7e, 0x72, 0x81, 0x8a, 0x4f, 0x9d, 0xa6, 0x01, 0x45, 0xee, 0x7c, 0xcc, 0xb4, 0x60, 0x2d,
		0x30, 0x1f, 0xd6, 0x28, 0x1d, 0x7a, 0x5d, 0xc9, 0xfd, 0x19, 0x55, 0xea, 0x82, 0xa4, 0x05, 0x23,
		0xb1, 0xbf, 0xee, 0x50, 0xdd, 0x37, 0x77, 0x41, 0xd2, 0x42, 0xb7, 0x98, 0x8f, 0x6e, 0x74, 0xb5,
		0x09, 0x72, 0x2a, 0xf7, 0x37, 0x3d, 0x54, 0x29, 0x6e, 0x5d, 0x90, 0xb4, 0x30, 0x5d, 0x60, 0x3e,
		0xa6, 0xa3, 0xdf, 0xa0, 0x0c, 0x26, 0x7f, 0x7f, 0x18, 0x7d, 0x07, 0x45, 0x6e, 0x5d, 0x90, 0xb4,
		0x80, 0x21, 0x30, 0x1f, 0x18, 0xc8, 0xa9, 0x2e, 0xf9, 0x58, 0xee, 0xaf, 0xa8, 0x21, 0xcd, 0x54,
		0x17, 0x24, 0x2d, 0x14, 0x89, 0xf9, 0x28, 0x42, 0x83, 0x3a, 0x39, 0x95, 0xfb, 0xeb, 0xd4, 0xc2,
		0x9f, 0x74, 0x41, 0xd2, 0x42, 0x47, 0x60, 0x3e, 0x3a, 0xc8, 0x40, 0xd2, 0x0f, 0x72, 0x7f, 0x89,
		0xc2, 0x36, 0x0b, 0xff, 0x2e, 0x48, 0x5a, 0x48, 0x14, 0xf3, 0x91, 0x88, 0x26, 0x92, 0xa9, 0x42,
		0xee, 0x4f, 0xbb, 0x74, 0x5f, 0x49, 0xba, 0x20, 0x69, 0x41, 0x5b, 0xec, 0x6f, 0x59, 0x8d, 0xee,
		0xb5, 0x77, 0x41, 0xd2, 0xc2, 0x32, 0x31, 0x1f, 0xcb, 0x50, 0x0b, 0x43, 0x34, 0xe5, 0xfe, 0xae,
		0xa7, 0xa8, 0x4b, 0x36, 0x5d, 0x90, 0xb4, 0x70, 0x5d, 0xcc, 0xc7, 0x75, 0xa4, 0xde, 0x55, 0x9c,
		0xe5, 0xfe, 0x74, 0x68, 0xd2, 0x95, 0xb5, 0xea, 0x82, 0xa4, 0x05, 0x1d, 0x04, 0xd6, 0x43, 0x80,
		0xcf, 0xfc, 0xe4, 0xb0, 0x16, 0x54, 0x54, 0x07, 0x82, 0x6f, 0x6c, 0x56, 0x1a, 0xd6, 0x02, 0x56,
		0xf5, 0x7a, 0x7d, 0xb4, 0xcf, 0x4b, 0xd2, 0x02, 0xa2, 0x07, 0xc0, 0xa6, 0xf8, 0x95, 0xf4, 0x76,
		0x29, 0x3e, 0xf3, 0x93, 0xd3, 0x94, 0xfb, 0x41, 0x03, 0xd8, 0x57, 0x7a, 0x2b, 0x85, 0x49, 0x37,
		0x4d, 0xb9, 0x1f, 0x34, 0xe0, 0x0a, 0x0d, 0xb6, 0xf8, 0x2a, 0xa4, 0x86, 0x53, 0x79, 0x5e, 0x92,
		0x16, 0x76, 0x3d, 0x00, 0x36, 0x45, 0x53, 0xd7, 0x82, 0xb6, 0x6c, 0xa1, 0xa7, 0xed, 0x1e, 0x9f,
		0xf9, 0xc9, 0xfd, 0xe6, 0xef, 0x07, 0x0d, 0x28, 0x55, 0xdc, 0xed, 0x2b, 0x45, 0xbf, 0xf9, 0xfb,
		0x41, 0x03, 0x7e, 0x44, 0xa9, 0xb3, 0x38, 0x36, 0xe9, 0x6a, 0x7e, 0xf2, 0xbc, 0x24, 0x2d, 0x68,
		0xeb, 0x01, 0xb0, 0x29, 0xc2, 0x5f, 0x16, 0xdf, 0xda, 0xe8, 0x19, 0x71, 0xaa, 0x21, 0x3f, 0x39,
		0x15, 0x77, 0xfc, 0x5d, 0xe1, 0x5b, 0x9f, 0xed, 0x2b, 0xd9, 0x75, 0x45, 0x15, 0xe8, 0x37, 0x35,
		0xf5, 0x7c, 0x41, 0x8f, 0x1a, 0xc4, 0xe2, 0x18, 0xaa, 0x06, 0xe5, 0x54, 0x9e, 0x97, 0xa4, 0x05,
		0x23, 0x3d, 0x00, 0x36, 0x45, 0xdf, 0x2e, 0x13, 0xb5, 0xd8, 0xe8, 0x4b, 0xe1, 0xc7, 0x4f, 0x4e,
		0xcd, 0xee, 0x07, 0x56, 0xfb, 0x4a, 0xb1, 0x52, 0x59, 0xd4, 0xec, 0x7e, 0x60, 0xa5, 0xf9, 0x82,
		0x63, 0xc4, 0xf1, 0xbc, 0x24, 0x2d, 0xea, 0x01, 0xb0, 0x29, 0x10, 0xc6, 0xa5, 0x79, 0x3c, 0xe6,
		0xe9, 0x75, 0x6c, 0xf3, 0xe3, 0x27, 0xa7, 0x62, 0xe3, 0x07, 0x56, 0xeb, 0xba, 0xcd, 0x4a, 0x15,
		0x50, 0xb1, 0xf1, 0x03, 0x2b, 0x7e, 0x82, 0xc7, 0x2c, 0xde, 0x2e, 0xdd, 0x47, 0x2c, 0xe4, 0x79,
		0x49, 0x5a, 0xe8, 0xe8, 0x01, 0xb0, 0x29, 0x50, 0xd6, 0x21, 0x1c, 0x6f, 0xd7, 0xbb, 0x51, 0xa0,
		0x21, 0x3f, 0x39, 0x15, 0x1b, 0xea, 0x29, 0xa4, 0x4d, 0x11, 0xa9, 0x5e, 0xbe, 0x92, 0x9d, 0x1a,
		0xdf, 0x52, 0x80, 0x70, 0x6d, 0x54, 0x66, 0x0b, 0x49, 0x8b, 0xb7, 0xfb, 0x18, 0xae, 0xe1, 0x9f,
		0x97, 0xa4, 0x85, 0x9b, 0x1e, 0x00, 0x9b, 0x02, 0x75, 0xae, 0xeb, 0x36, 0x7c, 0x29, 0x6e, 0x00,
	}

	// 2nd part -> 23	1.922887	8304	44074	TW7	576	sys.snap
	//             24	1.927097	44074	8304	TW7	57	ctrl.disconnect
	//             25	1.930487	8304	44074	TW7	912	sys.snap

	packet := protocol7.Packet{}
	err := packet.Unpack(dumpPart1)
	require.NoError(t, err)

	// TODO: not working yet
	// conn := protocol7.Session{}
	// conn.Ack = packet.Header.Ack
	// repack := packet.Pack(&conn)
	// require.Equal(t, dump, repack)

	// content
	require.Equal(t, 1, len(packet.Messages))
	require.Equal(t, network7.MsgSysSnap, packet.Messages[0].MsgId())
	part1, ok := packet.Messages[0].(*messages7.Snap)
	require.Equal(t, true, ok)
	require.Equal(t, 12847240, part1.GameTick)
	require.Equal(t, 12847241, part1.DeltaTick)
	require.Equal(t, 2, part1.NumParts)
	require.Equal(t, 1, part1.Part)
	require.Equal(t, -2136653645, part1.Crc)

	// Frame 25: 912 bytes on wire (7296 bits), 912 bytes captured (7296 bits)
	// Linux cooked capture v2
	// Internet Protocol Version 4
	// User Datagram Protocol, Src Port: 8304, Dst Port: 44074
	// Teeworlds 0.7 Protocol packet
	// Teeworlds 0.7 Protocol chunk: sys.snap
	//     Header (non-vital)
	//     Message: sys.snap
	//     Tick: 12847240
	//     Delta tick: 12847241
	//     Num parts: 2
	//     Part: 0
	//     Crc: -2136653645
	//     Data (900 bytes)
	dumpPart0 := []byte{
		0x10, 0x04, 0x01, 0x01, 0x02, 0x03, 0x04, 0xf4, 0x34, 0x1c, 0x2d, 0x3e, 0xf3, 0x93, 0xfb, 0xc8,
		0x67, 0x7e, 0x72, 0x51, 0xf1, 0xa9, 0xd3, 0x34, 0xa0, 0xc8, 0x9d, 0x46, 0xfa, 0x86, 0x58, 0x65,
		0x49, 0x0b, 0xcf, 0x34, 0x14, 0xcf, 0xa4, 0xe2, 0x3b, 0xa6, 0xc1, 0x4b, 0x21, 0x47, 0x3c, 0xa9,
		0x81, 0x42, 0xbf, 0x69, 0x10, 0xbd, 0x7b, 0x18, 0xf0, 0x2d, 0x55, 0xd8, 0xa4, 0x2b, 0x69, 0xe1,
		0x29, 0x9e, 0xaf, 0x4b, 0x3d, 0x5d, 0xc1, 0x11, 0xb8, 0x1c, 0xc3, 0xd9, 0x86, 0xe3, 0xe8, 0x4a,
		0xa6, 0x89, 0xf5, 0x8a, 0x69, 0x10, 0x17, 0xe0, 0xfa, 0x77, 0x41, 0xd2, 0xc2, 0xf3, 0xa3, 0x2b,
		0x43, 0xa1, 0x34, 0x8f, 0x7a, 0xc5, 0xdd, 0x4f, 0x8a, 0x82, 0xba, 0x15, 0x6a, 0x48, 0x53, 0x0e,
		0xc5, 0xae, 0xe6, 0x24, 0xc5, 0xe8, 0xf3, 0x92, 0xb4, 0xf0, 0x74, 0xa5, 0x98, 0xe6, 0xfb, 0x30,
		0x7a, 0x5f, 0x29, 0xc3, 0x12, 0x47, 0x49, 0xa5, 0x06, 0xeb, 0xba, 0xcf, 0x7a, 0x3a, 0x81, 0x62,
		0x13, 0xd3, 0x75, 0xba, 0x57, 0xd2, 0xc2, 0xe3, 0x4f, 0x55, 0xae, 0x36, 0x0d, 0xb9, 0x2b, 0xc3,
		0xe4, 0x43, 0xae, 0x8c, 0x8a, 0x6f, 0x22, 0x1f, 0x01, 0x7a, 0xfe, 0x52, 0x6f, 0x1a, 0x52, 0xe7,
		0x2e, 0x07, 0x49, 0x0b, 0x0f, 0x8a, 0x15, 0xcf, 0x8f, 0xae, 0x41, 0x68, 0x74, 0x0d, 0x26, 0xc8,
		0x98, 0xbf, 0x09, 0x98, 0x56, 0x8e, 0xf8, 0xc1, 0x61, 0x60, 0xb9, 0x25, 0x49, 0x79, 0x97, 0x49,
		0x5a, 0x78, 0x3e, 0xfb, 0x2c, 0x75, 0x3e, 0xb6, 0x53, 0xf4, 0x21, 0xa9, 0x65, 0xe1, 0x48, 0x8d,
		0xaf, 0x74, 0xa5, 0xb6, 0xe0, 0x8e, 0x31, 0x3f, 0xf8, 0x30, 0x8e, 0x75, 0x97, 0xa2, 0xa4, 0x85,
		0xe7, 0x32, 0xdb, 0xa6, 0x95, 0x02, 0x87, 0xc7, 0x38, 0x2b, 0x9e, 0x4c, 0x61, 0x02, 0x70, 0xa9,
		0xca, 0x57, 0x3a, 0x26, 0x96, 0xe1, 0x1e, 0xaa, 0x76, 0xd7, 0xe5, 0x92, 0x16, 0x9e, 0x74, 0xaa,
		0x34, 0x01, 0x5f, 0x47, 0xdf, 0x51, 0xa5, 0x8f, 0x14, 0xda, 0xa9, 0x8a, 0x97, 0x6b, 0xd2, 0x75,
		0x02, 0xd7, 0x16, 0x34, 0x28, 0xc7, 0x21, 0x15, 0x4a, 0x5a, 0x78, 0x52, 0xa8, 0x5e, 0x05, 0x75,
		0x5a, 0xd0, 0x52, 0x3c, 0xd9, 0x2d, 0xc9, 0x6d, 0x81, 0x93, 0x15, 0x1f, 0x2e, 0x53, 0x13, 0x75,
		0xfb, 0x4a, 0x8e, 0x45, 0x95, 0x25, 0x2d, 0x5a, 0xb8, 0x0a, 0x56, 0x65, 0x1f, 0x86, 0x4c, 0x80,
		0x36, 0x0f, 0x69, 0xa1, 0x8d, 0x23, 0x9e, 0xf0, 0xfc, 0x97, 0x66, 0x90, 0x2e, 0x48, 0x5a, 0x14,
		0xf3, 0x89, 0x2c, 0x18, 0x33, 0xb1, 0xdc, 0x1f, 0xc3, 0xba, 0xdc, 0xa6, 0x0b, 0x92, 0x16, 0x28,
		0xb0, 0xbf, 0xa8, 0xc5, 0x16, 0xdc, 0x75, 0x41, 0xd2, 0x42, 0x88, 0xf9, 0x08, 0x84, 0x9e, 0xba,
		0x65, 0xc8, 0xfd, 0x2d, 0x1a, 0x92, 0x36, 0x5d, 0x90, 0xb4, 0xb0, 0x10, 0xf3, 0xb1, 0x40, 0xcb,
		0xa9, 0xf7, 0xb1, 0xdc, 0xdf, 0x0b, 0xe8, 0xc4, 0xad, 0x0b, 0x92, 0x16, 0x9e, 0xc0, 0x7c, 0x3c,
		0xe4, 0xe0, 0x4a, 0xa5, 0xcb, 0xfd, 0xed, 0xa5, 0x53, 0xa5, 0x5b, 0x17, 0x24, 0x2d, 0xec, 0x62,
		0x7f, 0x9b, 0x14, 0x61, 0x6d, 0xba, 0x20, 0x69, 0x61, 0x23, 0xe6, 0x63, 0x83, 0x0a, 0xa9, 0x97,
		0x2e, 0xf7, 0x77, 0x4b, 0x51, 0x6a, 0xdd, 0xba, 0x20, 0x69, 0xe1, 0x26, 0x30, 0x1f, 0x37, 0xf4,
		0x9b, 0xae, 0x19, 0xfe, 0xfe, 0xa4, 0x46, 0x5d, 0x72, 0xeb, 0x82, 0xa4, 0x05, 0x11, 0xf3, 0x21,
		0xa8, 0x21, 0x07, 0xe9, 0x72, 0x7f, 0x66, 0xe9, 0xbc, 0x60, 0x74, 0x41, 0xd2, 0x82, 0x29, 0xe6,
		0xc3, 0x44, 0x03, 0xf0, 0x93, 0xfb, 0xeb, 0x53, 0x74, 0x35, 0xe9, 0x82, 0xa4, 0x85, 0x5e, 0xcc,
		0x47, 0x8f, 0xee, 0x75, 0x19, 0x67, 0xb9, 0xbf, 0x7f, 0x5e, 0xbe, 0xb6, 0x4d, 0x17, 0x24, 0x2d,
		0xfc, 0x05, 0xf6, 0x97, 0x37, 0xf8, 0xd6, 0xa6, 0x0b, 0x92, 0x16, 0x72, 0x31, 0x1f, 0x39, 0x2a,
		0x84, 0x62, 0x13, 0xcb, 0xfd, 0x9d, 0xc0, 0x72, 0xd2, 0x05, 0x49, 0x0b, 0x47, 0xcc, 0xc7, 0x41,
		0x97, 0xbb, 0xc0, 0x59, 0xee, 0x4f, 0x5f, 0x0b, 0x4c, 0xdc, 0x2e, 0x48, 0x5a, 0xd0, 0x0b, 0x8b,
		0x7c, 0xe8, 0xd1, 0x72, 0x95, 0x35, 0xdf, 0xdf, 0xdf, 0x5d, 0x8d, 0x30, 0xe9, 0xbb, 0x20, 0x69,
		0xe1, 0x4e, 0xec, 0xaf, 0xce, 0x8f, 0x02, 0xb7, 0x2e, 0x48, 0x5a, 0xa8, 0xc5, 0x7c, 0xd4, 0x08,
		0x97, 0x23, 0xb2, 0xcb, 0xfd, 0x4d, 0x95, 0x06, 0xc5, 0x0a, 0xba, 0x20, 0x69, 0x61, 0x4a, 0xcc,
		0xc7, 0x14, 0x32, 0x64, 0x0b, 0x9c, 0xca, 0xfd, 0xb9, 0x29, 0xf2, 0xb8, 0xbd, 0x0b, 0x92, 0x16,
		0x5c, 0x31, 0x1f, 0x2e, 0xca, 0xee, 0x6a, 0x7e, 0x72, 0x7f, 0x55, 0x79, 0xf1, 0x38, 0xe9, 0x82,
		0xa4, 0x85, 0x2a, 0x31, 0x1f, 0x55, 0xe8, 0x47, 0x99, 0xfc, 0xe4, 0xfe, 0xf0, 0xa5, 0x73, 0xc4,
		0x87, 0x77, 0x41, 0xd2, 0x02, 0x5e, 0x60, 0x7f, 0xdf, 0xf2, 0xe2, 0xa5, 0xef, 0x82, 0xa4, 0x85,
		0x6f, 0x62, 0x3e, 0xbe, 0x21, 0x2f, 0x3f, 0xf0, 0x93, 0xfb, 0xd3, 0x48, 0xcb, 0x55, 0x6e, 0x5d,
		0x90, 0xb4, 0xa0, 0x21, 0xe6, 0x43, 0x03, 0x5d, 0xae, 0x21, 0x3f, 0xb9, 0x3f, 0x9f, 0x06, 0x14,
		0x38, 0x5d, 0x90, 0xb4, 0xe0, 0x23, 0xe6, 0xc3, 0x07, 0x5d, 0xd0, 0xc4, 0xa9, 0xdc, 0x1f, 0x52,
		0xa9, 0x0b, 0x92, 0x16, 0x10, 0x81, 0xfd, 0x79, 0x96, 0x66, 0x62, 0xb7, 0x2e, 0x48, 0x5a, 0xf0,
		0x14, 0xf3, 0xe1, 0x89, 0xee, 0x91, 0xce, 0x4f, 0xee, 0x2f, 0x21, 0x2c, 0xd5, 0xb8, 0x75, 0x41,
		0xd2, 0x42, 0x82, 0x98, 0x8f, 0x04, 0xf4, 0xb1, 0xaf, 0xcd, 0xa9, 0xdc, 0x1f, 0x5a, 0xa8, 0x9c,
		0x4c, 0xba, 0x20, 0x69, 0x01, 0x4d, 0xec, 0xaf, 0xa0, 0x46, 0xaa, 0xa6, 0xef, 0x82, 0xa4, 0x85,
		0x02, 0x31, 0x1f, 0x05, 0x48, 0x3d, 0xf4, 0x3e, 0x96, 0xfb, 0x9b, 0xc3, 0x52, 0xa7, 0x4d, 0x17,
		0x24, 0x2d, 0xcc, 0x62, 0x3e, 0x66, 0x74, 0x8f, 0x1f, 0x38, 0x95, 0xfb, 0x2b, 0x97, 0xc6, 0x90,
		0xbd, 0x0b, 0x92, 0x16, 0xca, 0x62, 0x3e, 0xca, 0xa8, 0x18, 0x63, 0xfc, 0xe4, 0xfe, 0x2e, 0xd5,
		0xc8, 0x99, 0x74, 0x41, 0xd2, 0xc2, 0x25, 0x81, 0xfd, 0x59, 0x57, 0xea, 0x82, 0xa4, 0xe2, 0x06,
	}

	//             23	1.922887	8304	44074	TW7	576	sys.snap
	//             24	1.927097	44074	8304	TW7	57	ctrl.disconnect
	// 1st part -> 25	1.930487	8304	44074	TW7	912	sys.snap

	packet = protocol7.Packet{}
	err = packet.Unpack(dumpPart0)
	require.NoError(t, err)

	// content
	require.Equal(t, 1, len(packet.Messages))
	require.Equal(t, network7.MsgSysSnap, packet.Messages[0].MsgId())
	part0, ok := packet.Messages[0].(*messages7.Snap)
	require.Equal(t, true, ok)
	require.Equal(t, 12847240, part0.GameTick)
	require.Equal(t, 12847241, part0.DeltaTick)
	require.Equal(t, 2, part0.NumParts)
	require.Equal(t, 0, part0.Part)
	require.Equal(t, -2136653645, part0.Crc)

	// ------------------------------------
	// client with state and delta unpacker
	// ------------------------------------

	client := teeworlds7.NewClient()

	// part 0-2
	client.SnapshotStorage.AddIncomingData(part1.Part, part1.NumParts, part1.Data)
	client.SnapshotStorage.AddIncomingData(part0.Part, part0.NumParts, part0.Data) // out of order

	// TODO: this crashes

	// // we don't have the actual prev snap here
	// // just use an empty snap should be fine too
	// // then the final values will be wrong but it should still parse correctly i think
	// prevSnap, found := client.SnapshotStorage.Get(snapshot7.EmptySnapTick)
	// require.True(t, found)

	// u := &packer.Unpacker{}
	// u.Reset(client.SnapshotStorage.IncomingData())

	// newFullSnap, err := snapshot7.UnpackDelta(prevSnap, u)
	// require.NoError(t, err)

	// // TODO: should this be part0 here?
	// err = client.SnapshotStorage.Add(part0.GameTick, newFullSnap)
	// require.NoError(t, err)

	// // TODO:
	// // require.Equal(t, 999999, len(newFullSnap.Items))
}
