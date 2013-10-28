package main

import (
	"fmt"
	"bytes"
	"encoding/binary"
	"encoding/base64"
	"encoding/hex"
	"math"
	"os"
	"io"
)


func main (){
	fmt.Println("testing buffer")
	fmt.Println("Buffer is a variable-sized buffer of bytes with Read and Write methods. The zero value of Buffer is an Empty buffer ready to use")
	var b bytes.Buffer
	b.Write( []byte("hello") )
	b.Write( []byte("world") )
	b.WriteByte( 'c' )
	fmt.Println("Dumping b:")
	fmt.Println( hex.Dump( b.Bytes() ) )
	fmt.Println(">>>>>>>>>>>")

	fmt.Println("A Buffer can turn a string or a []byte into an io.Reader\n")
	buf2 := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder( base64.StdEncoding, buf2 )
	io.Copy( os.Stdout, dec )
	fmt.Println("---------------")
	return

	fmt.Println("Hello, world!")

	buf := new( bytes.Buffer )
	var pi float64 = math.Pi
	err := binary.Write( buf, binary.LittleEndian, pi )
	
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x\n", buf.Bytes() )
	return
}
