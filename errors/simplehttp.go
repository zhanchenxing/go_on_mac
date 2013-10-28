package main

import (
	"net"
	"os"
	"bytes"
	"fmt"
	"io"
)

func main(){
	if len(os.Args) != 2{
		fmt.Fprintf( os.Stderr, "Usage: %s host:port", os.Args[0] )
		os.Exit(1)
	}

	service := os.Args[1]

	conn, err := net.Dial( "tcp", service )
	checkError( err )

	// Write(b []byte) (n int, err error)
	_, err = conn.Write( []byte( "HEAD / HTTP/1.0\r\n\r\n") )
	checkError( err )

	result, err := readFully(conn)
	checkError(err)

	fmt.Println( string(result) )
	os.Exit(0)
}

func checkError( err error ){
	if err != nil {
		fmt.Fprintf( os.Stderr, "Fatal error: %s", err.Error() )
		os.Exit(1)
	}
}

func readFully( conn net.Conn ) ( []byte, error ) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte // this is an array, definitely
	b := [2]string{"text 1", "text 2"}
	fmt.Println("b=",b)

	// this is an array, too!
	c := [...]string{"text 1", "text 2", "text 3", "text 4"}
	fmt.Println("c=",c)
	
	// The type specification for a slice is []T, 
	// The distinction is: have size or not

	// this is a slice
	// letters := []string{"a", "b", "c", "d", "e", "f", "g" }

	// slice can be make
	let := make( []string, 3, 6 )
	fmt.Println( "let=", let)

	// A slice can also be formed by "slicing" an existing slice or array


	for {
		n, err := conn.Read( buf[0:] )
		result.Write( buf[0:n] )
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return result.Bytes(), nil
}
