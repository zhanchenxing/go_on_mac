package main

import (
	"fmt"
	"net"
	"log"
	//"time"
	"crypto/aes"
	"crypto/cipher"
    "io"
    "aesrw"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06,
	0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

// step 1: aes decrypt on received any thing, and aes back to client

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

type AESReader struct {
    readFrom io.Reader

    left int
    buff []byte
    cbcdec cipher.BlockMode
}

func ( r *AESReader ) Init( toReadFrom io.Reader)error{
    r.buff = make( []byte, 1024*4 )
    r.readFrom = toReadFrom
    r.left = 0

	key_text := "1234567890123456"
	c, err := aes.NewCipher( []byte(key_text) )
    if err != nil {
        return err
    }

    r.cbcdec = cipher.NewCBCDecrypter( c, commonIV )
    return nil
}

func ( r *AESReader ) processLeft( to []byte)(n int){
    canDo := r.left/16*16
    willLeft := r.left-canDo

    for i := 0; i < canDo; i += 16 {
        r.cbcdec.CryptBlocks( to[i:i+16], r.buff[i:i+16] )
    }
    
    copy( r.buff[:willLeft], r.buff[canDo:] )
    r.left = willLeft
    return canDo
}

func ( r *AESReader) Read(b []byte) (n int, err error){
    l, err := r.readFrom.Read( r.buff[r.left:] )
    r.left += l
    if err != nil {
        return r.processLeft(b), err
    }

    return r.processLeft(b), nil
}


func handleConnection ( conn net.Conn ){
	buff := make( []byte, 1024 )
    ar := aesrw.AESReader{}
    ar.Init(conn)
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
