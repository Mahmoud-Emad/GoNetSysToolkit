package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on UDP port 8080
	conn, err := net.ListenPacket("udp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Listening on 127.0.0.1:8080")
	buffer := make([]byte, 1024)

	for {
		// Read from connection
		n, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}
		fmt.Printf("Received message from %s: %s\n", addr.String(), string(buffer[:n]))

		// Echo message back to sender
		_, err = conn.WriteTo(buffer[:n], addr)
		if err != nil {
			fmt.Println("Error writing:", err)
			continue
		}
	}
}
