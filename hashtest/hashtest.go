package main

import (
	"fmt"
	"crypto/sha256"
	"crypto/sha1"
	"crypto/md5"
	"encoding/base64"
	"io"
)

func base64Encode ( src []byte ) []byte {
	return []byte( base64.StdEncoding.EncodeToString(src ) )
}

func base64Decode ( src []byte ) ( []byte, error ) {
	return base64.StdEncoding.DecodeString( string(src) )
}

func testBase64(){

	// encode
	hello := "this is the content to encode"
	debyte := base64Encode( []byte(hello) )
	fmt.Println("debyte=", string(debyte) )

	// decode
	enbyte, err := base64Decode( debyte )
	if err != nil {
		fmt.Println("base64Decode failed", err)
		return 
	}

	if hello != string(enbyte) {
		fmt.Println("hello is NOT EQUAL to enbyte")
	} else {
		fmt.Println("hello is EQUAL to enbyte")
	}

	fmt.Println( "decoded=", string(enbyte) )
}



func Password(){
	h := md5.New()
	io.WriteString( h, "this is password")

	pwmd5 := fmt.Sprintf("%x", h.Sum(nil) )
	fmt.Printf("pwmd5=%s\n", pwmd5)

	salt1 := "@#$%"
	salt2 := "^&*()"

	io.WriteString( h, salt1 )
	io.WriteString( h, "username" )
	io.WriteString( h, salt2)
	io.WriteString( h, pwmd5 )
	last := fmt.Sprintf("%x", h.Sum(nil ) )
	fmt.Println("last = ", last )
}

func main(){
	testBase64()

	Password()

	fmt.Println("Hello, world")

	h := sha256.New()
	io.WriteString( h, "abcdefg")
	fmt.Println("sha256:")
	fmt.Printf("% x\n", h.Sum(nil) )

	h2 := sha1.New()
	io.WriteString( h2, "abcdefg" )
	fmt.Println("sha1:")
	fmt.Printf("% x\n", h2.Sum(nil) )

	h3 := md5.New()
	io.WriteString(h3, "abcdefg")
	fmt.Println("md5:")
	fmt.Printf("%x\n", h3.Sum(nil) )

}

