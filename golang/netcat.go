package main

import (
	"io"
	"net"
	"time"
)

func main() {
	// create listener
	listener, _ := net.Listen("tcp", "localhost:8080")

	// wait for conn && accept conn
	for {
		conn, _ := listener.Accept()
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	//close conn
	defer c.Close()
	for {
		// sent string to the c object
		_, err := io.WriteString(c, time.Now().Format("15:05:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
