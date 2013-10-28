package main

import (
	"log"
	"net"
)

func main() {
	log.Print("Trying to connect...")
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected with", conn)

	conn.Write([]byte("hello, world"))

	buff := make([]byte, 256)
	received, err := conn.Read(buff)
	log.Print("received:[", string(buff[:received]), "]")
	if err != nil {
		log.Print("error is:", err)
	}
}
