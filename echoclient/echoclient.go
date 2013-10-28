package main

import (
	"fmt"
	"net"
	"log"
	// "encoding/binary"
)

func main(){
	x := []int{1,2,3}
	y := []int{4,5,6}
	x = append(x, y...)
	fmt.Println(x)


	s := make( []byte, 12 )
	fmt.Println("len(s)=", len(s) )

	fmt.Println("echoserver is listening on 8080")
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal( err )
	}

	go handleConnection( conn )

	for {
		var text string
		fmt.Scan(&text)

		if text=="exit"{
			return
		}

		if len(text)>0{
			log.Print( "Sending ", text )
			conn.Write([]byte(text))
		}
	}
}

func handleConnection ( conn net.Conn ){
    buf := make( []byte, 1024 )
	// buff := make( []byte, 1024 )
	for {
		// var num int32 
        /*
		err := binary.Read( conn, binary.LittleEndian, buf )
		if err != nil {
			fmt.Println("read uint32 failed: ", err )
		} else {
			fmt.Println("read uint32 value:", buf )
		}*/


		received, err := conn.Read( buf )
		if received > 0 {
			log.Print("received:[", string(buf[0:received]), "]" )
		} else {
			log.Print("nothing received!")
		}

		if err != nil {
			log.Print( "error happened:", err )
			conn.Close()
			return
		}
	}
}
