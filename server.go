package main

import (
	"fmt"
	. "game/packet"
	"log"
	"net"
	"os"
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

	ch := make(chan Packet, 100000)

	for {
		ch <- packet
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
