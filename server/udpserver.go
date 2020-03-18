package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3334"
	CONN_TYPE = "udp"
)

// Listen Start listening on udp port
func main() {
	// Listen for incoming connections.
	conn, err := net.ListenPacket(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer conn.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		// Handle connections in a new goroutine.
		handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.PacketConn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, raddr, err := conn.ReadFrom(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	tmpbuff := bytes.NewBuffer(buf)
	tmpstruct := new(Cron)
	gobobj := gob.NewDecoder(tmpbuff)
	gobobj.Decode(tmpstruct)
	fmt.Printf("%+v", string(tmpstruct.Output))
	fmt.Printf("%+v", tmpstruct)
	// Send a response back to person contacting us.
	conn.WriteTo([]byte("Message received."), raddr)
	// Close the connection when you're done with it.
}
