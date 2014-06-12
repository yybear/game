package main

import (
	"encoding/json"
	. "game/event"
	. "game/packet"
	"log"
	"net"
	"os"
	//"strconv"
)

func main() {
	service := ":" + os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	log.Println("server start")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	var packet Packet
	packet.Decode(conn)

	var req interface{}
	err := json.Unmarshal(packet.Data, &req)
	checkError(err)

	m := req.(map[string]interface{})
	log.Println("req is:", m)

	//sessionId := m["session_id"].(string)
	//appId := int(m["app_id"].(float64))
	code := int(m["type"].(float64))
	switch code {
	case LOGIN_GAME:
		log.Println("user login")
	default:
		log.Println("user req type: ", m["type"])
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
