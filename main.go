package main

import (
	"fmt"
	"net"
)

func main() {
	// Set up a UDP listen address and port
	listenAddr := &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: 8888}

	// Create a UDP listener on the listen address
	conn, err := net.ListenUDP("udp", listenAddr)
	if err != nil {
		panic(err)
	}

	// Loop forever, receiving and printing messages
	for {
		msg := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(msg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Received message: %s\n", string(msg[:n]))
	}
}
