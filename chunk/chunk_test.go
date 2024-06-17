package chunk

import (
	"reflect"
	"testing"
)

func TestBrokenNonVitalHeader(t *testing.T) {
	// this is a real vital header at a wrong offset
	// so it creates a actually non vital header with a size that is bigger than usual
	// verified results with teeworlds-network/twnet_parser python lib

	header := ChunkHeader{}
	//            {0x40, 0x3a, 0x01}
	header.Unpack([]byte{0x3a, 0x01})

	want := ChunkHeader {
		Flags: ChunkFlags {
			Vital: false,
			Resend: false,
		},
		Size: 3713,
		Seq: 0,
	}

	if !reflect.DeepEqual(header, want) {
		t.Errorf("got %v, wanted %v", header, want)
	}
}

func TestVitalHeaderMapChange(t *testing.T) {
	// generated by vanilla teeworlds 0.7 server
	// verified with libtw2 wireshark dissector

	header := ChunkHeader{}
	header.Unpack([]byte{0x40, 0x3a, 0x01})

	want := ChunkHeader {
		Flags: ChunkFlags {
			Vital: true,
			Resend: false,
		},
		Size: 58,
		Seq: 1,
	}

	if !reflect.DeepEqual(header, want) {
		t.Errorf("got %v, wanted %v", header, want)
	}
}

func TestVitalHeader(t *testing.T) {
	header := ChunkHeader{}
	header.Unpack([]byte{0x40, 0x10, 0x0a})
	want := ChunkHeader {
		Flags: ChunkFlags {
			Vital: true,
			Resend: false,
		},
		Size: 16,
		Seq: 10,
	}

	if !reflect.DeepEqual(header, want) {
		t.Errorf("got %v, wanted %v", header, want)
	}
}
