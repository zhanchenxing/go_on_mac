package main

import (
	"fmt"
	"net"
	"log"
	//"time"
)

func main(){
	fmt.Println("echoserver is listening on 8080")

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

		log.Printf("[%v] connected!", conn.RemoteAddr() )
		go handleConnection( conn ) 
	}
}

func handleConnection ( conn net.Conn ){
	buff := make( []byte, 1024 )
	for {
		received, err := conn.Read( buff )
		if received > 0 {
			log.Print("received:[", string(buff[0:received]), "]" )
            /*
			for i:=0; i<received; i++ {
				conn.Write( buff[i:i+1] )
				time.Sleep( 3 * time.Second )
			}*/
            conn.Write( buff[0:received] )
		} else {
			log.Print("nothing received!")
		}

		if err != nil {
			log.Printf("[%v] error happened:[%v]", conn.RemoteAddr(), err )
			conn.Close()
			return
		}
	}
}
