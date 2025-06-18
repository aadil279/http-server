package main

import (
	"fmt"
	"net"
	"strconv"
)

// Desired port
const PORT int = 8083

func main() {
	fmt.Println("Opening TCP server on port ", PORT, "...")

	// Open up a TCP socket on the desired port
	sock, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))

	if err != nil {
		fmt.Println(err)
		return
	}

	// Close socket after program exits
	defer sock.Close()
	fmt.Println("Successfully started TCP server!")

	// Handle incoming connections
	for {
		conn, err := sock.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		// Handle the connection asynchronously
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Allocate 1024 bytes in memory for connection buffer
	buf := make([]byte, 1024)

	// While the connection is alive, echo back any messages received
	for {
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println(err)
			return
		}

		// Echo back the message we just received
		message := string(buf[:n])
		conn.Write([]byte("Message received: " + message))
	}
}
