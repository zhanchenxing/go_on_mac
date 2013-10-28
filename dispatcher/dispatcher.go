package main

import (
	"fmt"
	"net"
	"log"
)

func main(){
	fmt.Println("hello, this is dispatcher!")

	ln, err := net.Listen( "tcp", ":8080" )
	if err != nil {
		log.Fatal( err )
	}

	log.Print( "Listening..." )

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Print( err )
			continue
		}

		// log.Print( conn, "connected!" )
		go handleConnection( conn ) 
	}
}

func handleConnection ( conn net.Conn ){
	buff := make( []byte, 1024 )
	received, err := conn.Read( buff )
	if received > 0 {
		log.Print("received:[", string(buff[0:received]), "]" )
		conn.Write( buff[0:received] )
	} else {
		log.Print("nothing received!")
	}

	if err != nil {
		log.Print( "error happened:", err )
	}
	conn.Close()
}
