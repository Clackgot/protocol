package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/teeworlds-go/teeworlds/messages7"
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/protocol7"
)

const (
	maxPacksize = 1400
)

func getConnection() (net.Conn, error) {
	conn, err := net.Dial("udp", "127.0.0.1:8303")
	if err != nil {
		fmt.Printf("Some error %v", err)
	}
	return conn, err
}

func readNetwork(ch chan<- []byte, conn net.Conn) {
	buf := make([]byte, maxPacksize)

	for {
		len, err := bufio.NewReader(conn).Read(buf)
		packet := make([]byte, len)
		copy(packet[:], buf[:])
		if err == nil {
			ch <- packet
		} else {
			fmt.Printf("Some error %v\n", err)
			break
		}
	}

	conn.Close()
}

func main() {
	ch := make(chan []byte, maxPacksize)

	conn, err := getConnection()
	if err != nil {
		fmt.Printf("error connecting %v\n", err)
		os.Exit(1)
	}

	client := &protocol7.Connection{
		ClientToken: [4]byte{0x01, 0x02, 0x03, 0x04},
		ServerToken: [4]byte{0xff, 0xff, 0xff, 0xff},
		Ack:         0,
		Players:     make([]protocol7.Player, network7.MaxClients),
	}

	go readNetwork(ch, conn)

	packet := client.CtrlToken()
	conn.Write(packet.Pack(client))

	for {
		time.Sleep(10_000_000)
		select {
		case msg := <-ch:
			packet, err = client.OnPacket(msg)
			if err != nil {
				panic(err)
			}
			if packet != nil {

				// example of modifying outgoing traffic
				for i, msg := range packet.Messages {
					if msg.MsgId() == network7.MsgCtrlConnect {
						if connect, ok := packet.Messages[0].(messages7.CtrlConnect); ok {
							connect.Token = [4]byte{0xaa, 0xaa, 0xaa, 0xaa}
							packet.Messages[i] = connect
						}
					}
				}

				conn.Write(packet.Pack(client))
			}
		default:
			// do nothing
		}
	}

}
