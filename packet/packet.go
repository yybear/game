package packet

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

const (
	PACKET_HEADER_LEN = 17
)

type Packet struct {
	Start      byte
	Version    byte
	PacketType byte
	Secret     byte
	Timestamp  uint64
	PacketSize uint16
	Padding    byte
	End        byte
	Cmd        byte
	Data       []byte
}

func (p *Packet) Decode(conn net.Conn) {
	header := make([]byte, PACKET_HEADER_LEN)

	for {
		n, err := io.ReadFull(conn, header)
		if err != nil {
			log.Println("error receiving header, bytes:", n, "reason:", err)
			break
		}
		if n == PACKET_HEADER_LEN {
			log.Println("get packet header")
			break
		}

	}

	p.Start = header[0]
	p.Version = header[1]
	p.PacketType = header[2]
	p.Secret = header[3]

	p.Timestamp = binary.BigEndian.Uint64(header[4:12])
	p.PacketSize = binary.BigEndian.Uint16(header[12:14])

	p.Padding = header[14]
	p.End = header[15]
	p.Cmd = header[16]

	data := make([]byte, p.PacketSize-1)
	for {
		n, err := io.ReadFull(conn, data)

		if err != nil {
			log.Println("error receiving msg, bytes:", n, "reason:", err)
			break
		}
		if uint16(n) == p.PacketSize {
			log.Println("get packet data size is:", n)
			break
		}
	}

	p.Data = data
}

/*func (p Packet) Encode() []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)

	err := enc.Encode(p)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	return network.Bytes()
}*/
