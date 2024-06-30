package snapshot7_test

// same connection as three_tees_localhost_big_snap_test.go

import (
	"testing"

	"github.com/teeworlds-go/go-teeworlds-protocol/internal/testutils/require"
	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/object7"
	"github.com/teeworlds-go/go-teeworlds-protocol/protocol7"
)

// localhost first connection
// 2 players already connected (id 0, id 1)
// map bridge_pickups

func TestFirstFewSnaps(t *testing.T) {
	t.Parallel()
	// snapshot captured with tcpdump
	// generated by a vanilla teeworlds 0.7.5 server
	// used https://github.com/ChillerDragon/teeworlds/tree/hacking-on-protocol client to connect
	// 0.7 vanilla based client with debug prints
	//
	// libtw2 dissector details
	// Teeworlds 0.7 Protocol packet
	//     Flags: compressed (..01 00..)
	//     Acknowledged sequence number: 4 (.... ..00 0000 0100)
	//     Number of chunks: 5
	//     Token: 57d3edf3
	//     Compressed payload (469 bytes)
	// Teeworlds 0.7 Protocol chunk: game.sv_game_info
	//     Header (vital: 9)
	//     Message: game.sv_game_info
	//     Game flags: none (0x0)
	//     Score limit: 20
	//     Time limit: 0
	//     Match num: 0
	//     Match current: 1
	// Teeworlds 0.7 Protocol chunk: game.sv_client_info
	//     Header (vital: 10)
	//     Message: game.sv_client_info
	//     Client id: 0
	//     Local: false
	//     Team: red
	//     Name: "dommy"
	//     Clan: "|*KoG*|"
	//     Country: 64
	//     Skin part names: "greensward"
	//     Skin part names: "duodonny"
	//     Skin part names: ""
	//     Skin part names: "standard"
	//     Skin part names: "standard"
	//     Skin part names: "standard"
	//     Use custom colors: true
	//     Use custom colors: true
	//     Use custom colors: false
	//     Use custom colors: false
	//     Use custom colors: false
	//     Use custom colors: false
	//     Skin part colors: 5635840
	//     Skin part colors: 5635860
	//     Skin part colors: 65408
	//     Skin part colors: 65408
	//     Skin part colors: 65408
	//     Skin part colors: 65408
	//     Silent: false
	// Teeworlds 0.7 Protocol chunk: game.sv_client_info
	//     Header (vital: 11)
	//     Message: game.sv_client_info
	//     Client id: 1
	//     Local: false
	//     Team: red
	//     Name: "ChillerDragon.*"
	//     Clan: "|*KoG*|"
	//     Country: 64
	//     Skin part names: "greensward"
	//     Skin part names: "duodonny"
	//     Skin part names: ""
	//     Skin part names: "standard"
	//     Skin part names: "standard"
	//     Skin part names: "standard"
	//     Use custom colors: true
	//     Use custom colors: true
	//     Use custom colors: false
	//     Use custom colors: false
	//     Use custom colors: false
	//     Use custom colors: false
	//     Skin part colors: 5635840
	//     Skin part colors: 5635860
	//     Skin part colors: 65408
	//     Skin part colors: 65408
	//     Skin part colors: 65408
	//     Skin part colors: 65408
	//     Silent: false
	// Teeworlds 0.7 Protocol chunk: game.sv_client_info
	//     Header (vital: 12)
	//     Message: game.sv_client_info
	//     Client id: 2
	//     Local: true
	//     Team: red
	//     Name: "ChillerDragon"
	//     Clan: ""
	//     Country: -1
	//     Skin part names: "greensward"
	//     Skin part names: "duodonny"
	//     Skin part names: ""
	//     Skin part names: "standard"
	//     Skin part names: "standard"
	//     Skin part names: "standard"
	//     Use custom colors: true
	//     Use custom colors: true
	//     Use custom colors: false
	//     Use custom colors: false
	//     Use custom colors: false
	//     Use custom colors: false
	//     Skin part colors: 5635840
	//     Skin part colors: -11141356
	//     Skin part colors: 65408
	//     Skin part colors: 65408
	//     Skin part colors: 65408
	//     Skin part colors: 65408
	//     Silent: false
	// Teeworlds 0.7 Protocol chunk: sys.snap_single
	//     Header (non-vital)
	//     Message: sys.snap_single
	//     Tick: 1020
	//     Delta tick: 1021
	//     Crc: 18302
	//     Data (149 bytes)

	dump := []byte{
		0x10, 0x04, 0x05, 0x57, 0xd3, 0xed, 0xf3,
		0x4a, 0x36, 0x4c, 0xed, 0xe1, 0x47, 0xde, 0x2e, 0xe9, 0x61, 0x9c,
		0x13, 0x6d, 0x11, 0x1d, 0xd1, 0x11, 0x94, 0xb7, 0x25, 0x12, 0x55, 0xdb, 0x22, 0x38, 0x24, 0xda,
		0x92, 0x00, 0xe9, 0x1c, 0x9c, 0x38, 0x81, 0x13, 0x5c, 0x3b, 0x9d, 0x2a, 0x6e, 0x39, 0x0a, 0x4e,
		0x9c, 0x98, 0x13, 0x6d, 0xa9, 0x2d, 0xe2, 0x44, 0x5b, 0xe4, 0xda, 0x5d, 0x3b, 0x94, 0x4f, 0xa7,
		0x50, 0xc2, 0x51, 0xae, 0x9d, 0x13, 0x39, 0x0a, 0x4e, 0x9c, 0x98, 0x4e, 0xa1, 0x84, 0xa3, 0x5c,
		0x3b, 0x27, 0x72, 0x14, 0x9c, 0x38, 0x31, 0x9d, 0x42, 0x09, 0x47, 0xb9, 0x76, 0x4e, 0xe4, 0x28,
		0x38, 0x71, 0x62, 0xc4, 0x07, 0xa6, 0xd5, 0x77, 0x76, 0xc3, 0xb4, 0xfa, 0xce, 0x0e, 0x34, 0xd8,
		0x0d, 0x68, 0xb0, 0x1b, 0xd0, 0x60, 0x37, 0xa0, 0xc1, 0x6e, 0x79, 0xeb, 0xf8, 0xc3, 0xc0, 0x19,
		0x77, 0xad, 0xae, 0xc5, 0x95, 0x70, 0x25, 0x38, 0x01, 0x4e, 0x9a, 0xc1, 0x89, 0xa3, 0xe8, 0x9c,
		0x2d, 0x72, 0xed, 0x2e, 0x4a, 0x6c, 0x4b, 0x24, 0xaa, 0xb6, 0x45, 0x70, 0x48, 0xb4, 0x25, 0x01,
		0xd2, 0x39, 0x38, 0x71, 0x02, 0x27, 0xb8, 0x76, 0x3a, 0x55, 0xdc, 0x72, 0x14, 0x9c, 0x38, 0x31,
		0x27, 0xda, 0x52, 0x5b, 0xc4, 0x89, 0xb6, 0xc8, 0xb5, 0xbb, 0x76, 0x28, 0x9f, 0x4e, 0xa1, 0x84,
		0xa3, 0x5c, 0x3b, 0x27, 0x72, 0x14, 0x9c, 0x38, 0x31, 0x9d, 0x42, 0x09, 0x47, 0xb9, 0x76, 0x4e,
		0xe4, 0x28, 0x38, 0x71, 0x62, 0x3a, 0x85, 0x12, 0x8e, 0x72, 0xed, 0x9c, 0xc8, 0x51, 0x70, 0xe2,
		0xc4, 0x88, 0x0f, 0x4c, 0xab, 0xef, 0xec, 0x86, 0x69, 0xf5, 0x9d, 0x1d, 0x68, 0xb0, 0x1b, 0xd0,
		0x60, 0x37, 0xa0, 0xc1, 0x6e, 0x40, 0x83, 0xdd, 0xf2, 0x76, 0x49, 0x0e, 0x43, 0x60, 0xc6, 0x5d,
		0xab, 0x6b, 0x71, 0x25, 0x5c, 0x09, 0x4e, 0x80, 0x93, 0x66, 0x70, 0xe2, 0x28, 0x3a, 0x67, 0x8b,
		0x5c, 0xfb, 0x4a, 0x74, 0x0e, 0x4e, 0x9c, 0xc0, 0x09, 0xae, 0x9d, 0x4e, 0x15, 0xb7, 0x1c, 0x05,
		0x27, 0x4e, 0xcc, 0x89, 0xb6, 0xd4, 0x16, 0x71, 0xa2, 0x2d, 0x72, 0xed, 0xae, 0x1d, 0xca, 0xa7,
		0x53, 0x28, 0xe1, 0x28, 0xd7, 0xce, 0x89, 0x1c, 0x05, 0x27, 0x4e, 0x4c, 0xa7, 0x50, 0xc2, 0x51,
		0xae, 0x9d, 0x13, 0x39, 0x0a, 0x4e, 0x9c, 0x98, 0x4e, 0xa1, 0x84, 0xa3, 0x5c, 0x3b, 0x27, 0x72,
		0x14, 0x9c, 0x38, 0x31, 0xe2, 0x03, 0xd3, 0xea, 0x3b, 0x3b, 0x55, 0xfb, 0x4d, 0xea, 0xf5, 0x40,
		0x83, 0xdd, 0x80, 0x06, 0xbb, 0x01, 0x0d, 0x76, 0x03, 0x1a, 0xec, 0x56, 0xb8, 0x64, 0x8a, 0x97,
		0x3b, 0x17, 0xdc, 0xd1, 0xf5, 0xb1, 0x50, 0x4c, 0x94, 0xf7, 0x10, 0xd6, 0x37, 0x93, 0x1e, 0x7a,
		0xc2, 0x24, 0x0d, 0x93, 0x5e, 0xcf, 0xc2, 0x24, 0x1f, 0x93, 0x9e, 0xf0, 0x3c, 0x93, 0x10, 0x93,
		0x9e, 0x85, 0x67, 0x37, 0xc9, 0xd3, 0xa4, 0xe7, 0xd9, 0x84, 0xea, 0xb6, 0xbe, 0x7b, 0xdd, 0x99,
		0x94, 0xfb, 0xc1, 0x1e, 0x7a, 0xf4, 0x91, 0x58, 0x69, 0x93, 0x72, 0x3f, 0xd8, 0x3f, 0x5e, 0x8f,
		0x86, 0x3b, 0x93, 0x8e, 0x1f, 0xec, 0x5f, 0x69, 0x93, 0x8e, 0x1f, 0xec, 0x1f, 0xaf, 0x17, 0xee,
		0x75, 0x67, 0x52, 0x2e, 0xcd, 0x4b, 0xd5, 0xa6, 0xae, 0xb4, 0x49, 0x39, 0xac, 0xb7, 0xbe, 0x1e,
		0xef, 0x9f, 0xec, 0x8f, 0x64, 0x7f, 0x41, 0x56, 0xdc, 0x00,
	}

	packet := protocol7.Packet{}
	err := packet.Unpack(dump)
	require.NoError(t, err)

	// TODO: not working yet
	// conn := protocol7.Session{}
	// conn.Ack = packet.Header.Ack
	// repack := packet.Pack(&conn)
	// require.Equal(t, dump, repack)

	// content
	require.Equal(t, 5, len(packet.Messages))
	require.Equal(t, network7.MsgGameSvGameInfo, packet.Messages[0].MsgId())
	require.Equal(t, network7.MsgGameSvClientInfo, packet.Messages[1].MsgId())
	require.Equal(t, network7.MsgGameSvClientInfo, packet.Messages[2].MsgId())
	require.Equal(t, network7.MsgGameSvClientInfo, packet.Messages[3].MsgId())
	require.Equal(t, network7.MsgSysSnapSingle, packet.Messages[4].MsgId())
	msg, ok := packet.Messages[4].(*messages7.SnapSingle)
	require.Equal(t, true, ok)
	require.Equal(t, 1020, msg.GameTick)
	require.Equal(t, 1021, msg.DeltaTick)
	require.Equal(t, 18302, msg.Crc)

	// verified with hacking on protocol print
	require.Equal(t, 12, msg.Snapshot.NumItemDeltas)
	require.Equal(t, 0, msg.Snapshot.NumRemovedItems)
	require.Equal(t, 12, len(msg.Snapshot.Items))

	require.Equal(t, 18302, msg.Snapshot.Crc)

	// verified with hacking on protocol
	item := msg.Snapshot.Items[0]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok := item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 1, pickup.Id())
	require.Equal(t, 1392, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[1]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 2, pickup.Id())
	require.Equal(t, 1424, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupHealth, pickup.Type)

	item = msg.Snapshot.Items[2]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 3, pickup.Id())
	require.Equal(t, 1488, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupGrenade, pickup.Type)

	item = msg.Snapshot.Items[3]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 4, pickup.Id())
	require.Equal(t, 1552, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupShotgun, pickup.Type)

	item = msg.Snapshot.Items[4]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 5, pickup.Id())
	require.Equal(t, 1616, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupLaser, pickup.Type)

	item = msg.Snapshot.Items[5]
	require.Equal(t, network7.ObjGameData, item.TypeId())
	gameData, ok := item.(*object7.GameData)
	require.Equal(t, true, ok)
	require.Equal(t, 0, gameData.Id())
	require.Equal(t, 500, gameData.GameStartTick)
	require.Equal(t, 0, gameData.FlagsRaw)
	require.Equal(t, 0, gameData.GameStateEndTick)

	item = msg.Snapshot.Items[6]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok := item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 0, character.Id())
	require.Equal(t, 1019, character.Tick)
	require.Equal(t, 784, character.X)
	require.Equal(t, 337, character.Y)
	require.Equal(t, 0, character.VelX)
	require.Equal(t, 109, character.VelY)
	require.Equal(t, 137, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, -1, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 784, character.HookX)
	require.Equal(t, 337, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 0, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 0, character.AmmoCount)
	require.Equal(t, network7.WeaponGun, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[7]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok = item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 1, character.Id())
	require.Equal(t, 980, character.Tick)
	require.Equal(t, 848, character.X)
	require.Equal(t, 337, character.Y)
	require.Equal(t, 0, character.VelX)
	require.Equal(t, 0, character.VelY)
	require.Equal(t, 0, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, -1, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 848, character.HookX)
	require.Equal(t, 337, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 0, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 0, character.AmmoCount)
	require.Equal(t, network7.WeaponGun, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[8]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok = item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 2, character.Id())
	require.Equal(t, 1019, character.Tick)
	require.Equal(t, 784, character.X)
	require.Equal(t, 300, character.Y)
	require.Equal(t, 0, character.VelX)
	require.Equal(t, -1132, character.VelY)
	require.Equal(t, 0, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, -1, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 784, character.HookX)
	require.Equal(t, 304, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 10, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 10, character.AmmoCount)
	require.Equal(t, network7.WeaponGun, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[9]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok := item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 0, playerInfo.Id())
	require.Equal(t, 8, playerInfo.PlayerFlags)
	require.Equal(t, 0, playerInfo.Score)
	require.Equal(t, 0, playerInfo.Latency)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[10]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok = item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 1, playerInfo.Id())
	require.Equal(t, 8, playerInfo.PlayerFlags)
	require.Equal(t, 0, playerInfo.Score)
	require.Equal(t, 0, playerInfo.Latency)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[11]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok = item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 2, playerInfo.Id())
	require.Equal(t, 8, playerInfo.PlayerFlags)
	require.Equal(t, 0, playerInfo.Score)
	require.Equal(t, 0, playerInfo.Latency)

	// libtw2 dissector details
	// Teeworlds 0.7 Protocol packet
	//     Flags: compressed (..01 00..)
	//     Acknowledged sequence number: 4 (.... ..00 0000 0100)
	//     Number of chunks: 1
	//     Token: 57d3edf3
	//     Compressed payload (111 bytes)
	// Teeworlds 0.7 Protocol chunk: sys.snap_single
	//     Header (non-vital)
	//     Message: sys.snap_single
	//     Tick: 1030
	//     Delta tick: 1031
	//     Crc: 18307
	//     Data (149 bytes)
	dump = []byte{
		0x10, 0x04, 0x01, 0x57, 0xd3, 0xed, 0xf3,
		0xc2, 0x25, 0x53, 0x52, 0xd4, 0xae, 0x56, 0xfb, 0x8d, 0xb3, 0x50,
		0x4c, 0x94, 0xf7, 0x10, 0xd6, 0x37, 0x93, 0x1e, 0x7a, 0xc2, 0x24, 0x0d, 0x93, 0x5e, 0xcf, 0xc2,
		0x24, 0x1f, 0x93, 0x9e, 0xf0, 0x3c, 0x93, 0x10, 0x93, 0x9e, 0x85, 0x67, 0x37, 0xc9, 0xd3, 0xa4,
		0xe7, 0xd9, 0x84, 0xea, 0xb6, 0xbe, 0x7b, 0xdd, 0x99, 0x94, 0xfb, 0xc1, 0x1e, 0x7a, 0xf4, 0x91,
		0x58, 0x69, 0x93, 0x72, 0x3f, 0xd8, 0x3f, 0x5e, 0x8f, 0x86, 0x3b, 0x93, 0x8e, 0x1f, 0xec, 0x5f,
		0x69, 0x93, 0x8e, 0x1f, 0xec, 0x1f, 0xed, 0xeb, 0x85, 0x7b, 0xdd, 0x99, 0x94, 0x4b, 0xf3, 0x52,
		0xb5, 0xa9, 0x2b, 0x6d, 0x52, 0x0e, 0xeb, 0xad, 0xaf, 0xc7, 0xfb, 0x27, 0xfb, 0x23, 0xd9, 0x5f,
		0x90, 0x15, 0x37, 0x00,
	}

	packet = protocol7.Packet{}
	err = packet.Unpack(dump)
	require.NoError(t, err)

	// TODO: not working yet
	// conn := protocol7.Session{}
	// conn.Ack = packet.Header.Ack
	// repack := packet.Pack(&conn)
	// require.Equal(t, dump, repack)

	// content
	require.Equal(t, 1, len(packet.Messages))
	require.Equal(t, network7.MsgSysSnapSingle, packet.Messages[0].MsgId())
	msg, ok = packet.Messages[0].(*messages7.SnapSingle)
	require.Equal(t, true, ok)
	require.Equal(t, 1030, msg.GameTick)
	require.Equal(t, 1031, msg.DeltaTick)
	require.Equal(t, 18307, msg.Crc)

	// verified with hacking on protocol print
	// this is interesting the server sent the full 12 items again
	// it also has different crc
	// because character with id 1 is now blinking its eyes (yes thats the only difference lmao)
	// the crc changed from 18302 to 18307 (18302+5) because emote blink (5) is 5 more than emote normal (0)
	// the client had no time yet to ack the first snapshot thats why the server resend it all
	//
	//            No. Time      Source Destination Protocol Length  Info
	//            12  0.060259  8303   42069       TW7      119     sys.server_info
	// 1st snap   13  0.113440  8303   42069       TW7      538     game.sv_game_info, game.sv_client_info, game.sv_client_info, game.sv_client_info, sys.snap_single
	// 2nd snap   14  0.313050  8303   42069       TW7      180     sys.snap_single
	//            15  0.513597  8303   42069       TW7      178     sys.snap_single
	//            16  0.528518  42069  8303        TW7      81      sys.input
	require.Equal(t, 12, msg.Snapshot.NumItemDeltas)
	require.Equal(t, 0, msg.Snapshot.NumRemovedItems)
	require.Equal(t, 12, len(msg.Snapshot.Items))

	require.Equal(t, 18307, msg.Snapshot.Crc)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[0]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 1, pickup.Id())
	require.Equal(t, 1392, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[1]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 2, pickup.Id())
	require.Equal(t, 1424, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupHealth, pickup.Type)

	item = msg.Snapshot.Items[2]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 3, pickup.Id())
	require.Equal(t, 1488, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupGrenade, pickup.Type)

	item = msg.Snapshot.Items[3]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 4, pickup.Id())
	require.Equal(t, 1552, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupShotgun, pickup.Type)

	item = msg.Snapshot.Items[4]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 5, pickup.Id())
	require.Equal(t, 1616, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupLaser, pickup.Type)

	item = msg.Snapshot.Items[5]
	require.Equal(t, network7.ObjGameData, item.TypeId())
	gameData, ok = item.(*object7.GameData)
	require.Equal(t, true, ok)
	require.Equal(t, 0, gameData.Id())
	require.Equal(t, 500, gameData.GameStartTick)
	require.Equal(t, 0, gameData.FlagsRaw)
	require.Equal(t, 0, gameData.GameStateEndTick)

	item = msg.Snapshot.Items[6]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok = item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 0, character.Id())
	require.Equal(t, 1019, character.Tick)
	require.Equal(t, 784, character.X)
	require.Equal(t, 337, character.Y)
	require.Equal(t, 0, character.VelX)
	require.Equal(t, 109, character.VelY)
	require.Equal(t, 137, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, -1, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 784, character.HookX)
	require.Equal(t, 337, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 0, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 0, character.AmmoCount)
	require.Equal(t, network7.WeaponGun, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[7]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok = item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 1, character.Id())
	require.Equal(t, 980, character.Tick)
	require.Equal(t, 848, character.X)
	require.Equal(t, 337, character.Y)
	require.Equal(t, 0, character.VelX)
	require.Equal(t, 0, character.VelY)
	require.Equal(t, 0, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, -1, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 848, character.HookX)
	require.Equal(t, 337, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 0, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 0, character.AmmoCount)
	require.Equal(t, network7.WeaponGun, character.Weapon)
	require.Equal(t, network7.EyeEmoteBlink, character.Emote) // this is the only difference to the previous snap
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[8]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok = item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 2, character.Id())
	require.Equal(t, 1019, character.Tick)
	require.Equal(t, 784, character.X)
	require.Equal(t, 300, character.Y)
	require.Equal(t, 0, character.VelX)
	require.Equal(t, -1132, character.VelY)
	require.Equal(t, 0, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, -1, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 784, character.HookX)
	require.Equal(t, 304, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 10, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 10, character.AmmoCount)
	require.Equal(t, network7.WeaponGun, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[9]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok = item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 0, playerInfo.Id())
	require.Equal(t, 8, playerInfo.PlayerFlags)
	require.Equal(t, 0, playerInfo.Score)
	require.Equal(t, 0, playerInfo.Latency)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[10]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok = item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 1, playerInfo.Id())
	require.Equal(t, 8, playerInfo.PlayerFlags)
	require.Equal(t, 0, playerInfo.Score)
	require.Equal(t, 0, playerInfo.Latency)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[11]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok = item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 2, playerInfo.Id())
	require.Equal(t, 8, playerInfo.PlayerFlags)
	require.Equal(t, 0, playerInfo.Score)
	require.Equal(t, 0, playerInfo.Latency)

	// libtw2 dissector details
	// Teeworlds 0.7 Protocol packet
	//     Flags: compressed (..01 00..)
	//     Acknowledged sequence number: 4 (.... ..00 0000 0100)
	//     Number of chunks: 1
	//     Token: 57d3edf3
	//     Compressed payload (109 bytes)
	// Teeworlds 0.7 Protocol chunk: sys.snap_single
	//     Header (non-vital)
	//     Message: sys.snap_single
	//     Tick: 1040
	//     Delta tick: 1041
	//     Crc: 18420
	//     Data (149 bytes)
	dump = []byte{
		0x10, 0x04, 0x01, 0x57, 0xd3, 0xed, 0xf3,
		0xc2, 0x25, 0x53, 0x26, 0xd5, 0x7e, 0x50, 0x43, 0xe5, 0x54, 0x28,
		0x26, 0xca, 0x7b, 0x08, 0xeb, 0x9b, 0x49, 0x0f, 0x3d, 0x61, 0x92, 0x86, 0x49, 0xaf, 0x67, 0x61,
		0x92, 0x8f, 0x49, 0x4f, 0x78, 0x9e, 0x49, 0x88, 0x49, 0xcf, 0xc2, 0xb3, 0x9b, 0xe4, 0x69, 0xd2,
		0xf3, 0x6c, 0x42, 0x75, 0x5b, 0x9f, 0x31, 0xb5, 0x49, 0xb9, 0x1f, 0xec, 0xa1, 0x47, 0x1f, 0x89,
		0x95, 0x36, 0x29, 0xf7, 0x83, 0xfd, 0xe3, 0xf5, 0x68, 0xb8, 0x33, 0xe9, 0xf8, 0xc1, 0xfe, 0x95,
		0x36, 0xe9, 0xf8, 0xc1, 0xfe, 0xf1, 0x7a, 0xc1, 0x98, 0xda, 0xa4, 0x9c, 0x7a, 0x2f, 0xdf, 0xab,
		0xaf, 0xb4, 0x49, 0xb9, 0xef, 0xbc, 0xf5, 0xf5, 0x78, 0xff, 0x64, 0x7f, 0x24, 0xfb, 0x0b, 0xb2,
		0xe2, 0x06,
	}

	packet = protocol7.Packet{}
	err = packet.Unpack(dump)
	require.NoError(t, err)

	// TODO: not working yet
	// conn := protocol7.Session{}
	// conn.Ack = packet.Header.Ack
	// repack := packet.Pack(&conn)
	// require.Equal(t, dump, repack)

	// content
	require.Equal(t, 1, len(packet.Messages))
	require.Equal(t, network7.MsgSysSnapSingle, packet.Messages[0].MsgId())
	msg, ok = packet.Messages[0].(*messages7.SnapSingle)
	require.Equal(t, true, ok)
	require.Equal(t, 1040, msg.GameTick)
	require.Equal(t, 1041, msg.DeltaTick)
	require.Equal(t, 18420, msg.Crc)

	// verified with hacking on protocol print
	//
	//            No. Time      Source Destination Protocol Length  Info
	//            12  0.060259  8303   42069       TW7      119     sys.server_info
	//            13  0.113440  8303   42069       TW7      538     game.sv_game_info, game.sv_client_info, game.sv_client_info, game.sv_client_info, sys.snap_single
	//            14  0.313050  8303   42069       TW7      180     sys.snap_single
	//  ------->  15  0.513597  8303   42069       TW7      178     sys.snap_single
	//            16  0.528518  42069  8303        TW7      81      sys.input
	require.Equal(t, 12, msg.Snapshot.NumItemDeltas)
	require.Equal(t, 0, msg.Snapshot.NumRemovedItems)
	require.Equal(t, 12, len(msg.Snapshot.Items))

	require.Equal(t, 18420, msg.Snapshot.Crc)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[0]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 1, pickup.Id())
	require.Equal(t, 1392, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupArmor, pickup.Type)

	item = msg.Snapshot.Items[1]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 2, pickup.Id())
	require.Equal(t, 1424, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupHealth, pickup.Type)

	item = msg.Snapshot.Items[2]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 3, pickup.Id())
	require.Equal(t, 1488, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupGrenade, pickup.Type)

	item = msg.Snapshot.Items[3]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 4, pickup.Id())
	require.Equal(t, 1552, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupShotgun, pickup.Type)

	item = msg.Snapshot.Items[4]
	require.Equal(t, network7.ObjPickup, item.TypeId())
	pickup, ok = item.(*object7.Pickup)
	require.Equal(t, true, ok)
	require.Equal(t, 5, pickup.Id())
	require.Equal(t, 1616, pickup.X)
	require.Equal(t, 272, pickup.Y)
	require.Equal(t, network7.PickupLaser, pickup.Type)

	item = msg.Snapshot.Items[5]
	require.Equal(t, network7.ObjGameData, item.TypeId())
	gameData, ok = item.(*object7.GameData)
	require.Equal(t, true, ok)
	require.Equal(t, 0, gameData.Id())
	require.Equal(t, 500, gameData.GameStartTick)
	require.Equal(t, 0, gameData.FlagsRaw)
	require.Equal(t, 0, gameData.GameStateEndTick)

	item = msg.Snapshot.Items[6]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok = item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 0, character.Id())
	require.Equal(t, 1039, character.Tick)
	require.Equal(t, 784, character.X)
	require.Equal(t, 337, character.Y)
	require.Equal(t, 0, character.VelX)
	require.Equal(t, 109, character.VelY)
	require.Equal(t, 137, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, -1, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 784, character.HookX)
	require.Equal(t, 337, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 0, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 0, character.AmmoCount)
	require.Equal(t, network7.WeaponGun, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[7]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok = item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 1, character.Id())
	require.Equal(t, 980, character.Tick)
	require.Equal(t, 848, character.X)
	require.Equal(t, 337, character.Y)
	require.Equal(t, 0, character.VelX)
	require.Equal(t, 0, character.VelY)
	require.Equal(t, 0, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, -1, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 848, character.HookX)
	require.Equal(t, 337, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 0, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 0, character.AmmoCount)
	require.Equal(t, network7.WeaponGun, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[8]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok = item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 2, character.Id())
	require.Equal(t, 1039, character.Tick)
	require.Equal(t, 784, character.X)
	require.Equal(t, 299, character.Y)
	require.Equal(t, 0, character.VelX)
	require.Equal(t, -1052, character.VelY)
	require.Equal(t, 0, character.Angle)
	require.Equal(t, 0, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, -1, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 784, character.HookX)
	require.Equal(t, 303, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	require.Equal(t, 10, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 10, character.AmmoCount)
	require.Equal(t, network7.WeaponGun, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[9]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok = item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 0, playerInfo.Id())
	require.Equal(t, 8, playerInfo.PlayerFlags)
	require.Equal(t, 0, playerInfo.Score)
	require.Equal(t, 0, playerInfo.Latency)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[10]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok = item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 1, playerInfo.Id())
	require.Equal(t, 8, playerInfo.PlayerFlags)
	require.Equal(t, 0, playerInfo.Score)
	require.Equal(t, 0, playerInfo.Latency)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[11]
	require.Equal(t, network7.ObjPlayerInfo, item.TypeId())
	playerInfo, ok = item.(*object7.PlayerInfo)
	require.Equal(t, true, ok)
	require.Equal(t, 2, playerInfo.Id())
	require.Equal(t, 8, playerInfo.PlayerFlags)
	require.Equal(t, 0, playerInfo.Score)
	require.Equal(t, 0, playerInfo.Latency)

	// --------------------------------------------
	// first delta snapshot
	// --------------------------------------------

	// libtw2 dissector details
	// Teeworlds 0.7 Protocol packet
	//     Flags: compressed (..01 00..)
	//     Acknowledged sequence number: 4 (.... ..00 0000 0100)
	//     Number of chunks: 2
	//     Token: 57d3edf3
	//     Compressed payload (31 bytes)
	// Teeworlds 0.7 Protocol chunk: sys.input_timing
	//     Header (non-vital)
	//     Message: sys.input_timing
	//     Input pred tick: 1056
	//     Time left: 21
	// Teeworlds 0.7 Protocol chunk: sys.snap_single
	//     Header (non-vital)
	//     Message: sys.snap_single
	//     Tick: 1056
	//     Delta tick: 2
	//     Crc: 20704
	//     Data (30 bytes)
	//
	//              No. Time      Source Destination Protocol Length  Info
	//              12  0.060259  8303   42069       TW7      119     sys.server_info
	//              13  0.113440  8303   42069       TW7      538     game.sv_game_info, game.sv_client_info, game.sv_client_info, game.sv_client_info, sys.snap_single
	//              14  0.313050  8303   42069       TW7      180     sys.snap_single
	//              15  0.513597  8303   42069       TW7      178     sys.snap_single
	//              16  0.528518  42069  8303        TW7      81      sys.input
	// ...
	//              30  0.811001  42069  8303        TW7      84      sys.input
	// 1st delta -> 31  0.833344  8303   42069       TW7      100     sys.input_timing, sys.snap_single
	//              32  0.852673  42069  8303        TW7      85      sys.input
	//
	dump = []byte{
		0x10, 0x04, 0x02, 0x57, 0xd3, 0xed, 0xf3,
		0x3d, 0xdf, 0xf8, 0xa9, 0x7d, 0xab, 0xdd, 0x14, 0x3f, 0xb5, 0xe0,
		0x67, 0xcb, 0x42, 0x39, 0xd6, 0x0b, 0x53, 0xe8, 0x01, 0x1b, 0xa0, 0xdb, 0xd7, 0x3d, 0xfc, 0x8c,
		0xff, 0x57, 0xdc, 0x00,
	}

	packet = protocol7.Packet{}
	err = packet.Unpack(dump)
	require.NoError(t, err)

	// TODO: not working yet
	// conn := protocol7.Session{}
	// conn.Ack = packet.Header.Ack
	// repack := packet.Pack(&conn)
	// require.Equal(t, dump, repack)

	// content
	require.Equal(t, 2, len(packet.Messages))
	require.Equal(t, network7.MsgSysInputTiming, packet.Messages[0].MsgId())
	inputTiming, ok := packet.Messages[0].(*messages7.InputTiming)
	require.Equal(t, true, ok)
	require.Equal(t, 1056, inputTiming.IntendedPredTick)
	require.Equal(t, 21, inputTiming.TimeLeft)

	require.Equal(t, network7.MsgSysSnapSingle, packet.Messages[1].MsgId())
	msg, ok = packet.Messages[1].(*messages7.SnapSingle)
	require.Equal(t, true, ok)
	require.Equal(t, 1056, msg.GameTick)
	require.Equal(t, 2, msg.DeltaTick)
	require.Equal(t, 20704, msg.Crc)

	require.Equal(t, 1, msg.Snapshot.NumItemDeltas)
	require.Equal(t, 0, msg.Snapshot.NumRemovedItems)
	require.Equal(t, 1, len(msg.Snapshot.Items))

	// TODO: delta against the previous one to get the correct crc
	// require.Equal(t, 20704, msg.Snapshot.Crc)

	// verified with hacking on protocol
	item = msg.Snapshot.Items[0]
	require.Equal(t, network7.ObjCharacter, item.TypeId())
	character, ok = item.(*object7.Character)
	require.Equal(t, true, ok)
	require.Equal(t, 2, character.Id())
	require.Equal(t, 17, character.Tick)
	require.Equal(t, 1, character.X)
	require.Equal(t, 4, character.Y)
	require.Equal(t, 384, character.VelX)
	require.Equal(t, 2176, character.VelY)
	require.Equal(t, -295, character.Angle)
	require.Equal(t, 1, character.Direction)
	require.Equal(t, 0, character.Jumped)
	require.Equal(t, 0, character.HookedPlayer)
	require.Equal(t, 0, character.HookState)
	require.Equal(t, 0, character.HookTick)
	require.Equal(t, 0, character.HookX)
	require.Equal(t, -4, character.HookY)
	require.Equal(t, 0, character.HookDx)
	require.Equal(t, 0, character.HookDy)
	// health 0 is verified with hacking on protocol
	// but the value is still odd
	// because in the non delta snapshots the health was 10
	// its character with id 2 so that has be us
	// there were two characters (id 0 and id 1) already on the server
	// so when we connect with the hacking on protocol client
	// we get id 2 ( see also Local: true in the details further up )
	// and we should be able to see our own health to display it in the hud
	// and we also did in the first 3 non delta snapshots
	// but the delta then claims that we have a health of 0
	// iirc i did not die in the first second after joining
	// so what is going on here?
	require.Equal(t, 0, character.Health)
	require.Equal(t, 0, character.Armor)
	require.Equal(t, 0, character.AmmoCount)
	require.Equal(t, network7.WeaponHammer, character.Weapon)
	require.Equal(t, network7.EyeEmoteNormal, character.Emote)
	require.Equal(t, 0, character.AttackTick)
	require.Equal(t, 0, character.TriggeredEvents)
}