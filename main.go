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
	// Close connection if interrupted
	defer conn.Close()

	// Allocate 1024 bytes in memory for connection buffer
	buf := make([]byte, 1024)

	// While the connection is alive, echo back any messages received

	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Print back the message we just received
	fmt.Println(string(buf[:n]))

	// HTTP repsonse
	message := `
	HTTP/1.1 200 OK
	Date: Tue, 22 Jun 2024 18:30:00 GMT
	Content-Type: text/html; charset=UTF-8
	Content-Length: 1234

	<html>
		<body>
		<h1>Hello, World!</h1>
		</body>
	</html>`

	// Send the response back
	conn.Write([]byte(message))
	conn.Close()
}
