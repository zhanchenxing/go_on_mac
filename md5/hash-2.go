package main

import (
	"io"
	"fmt"
	"os"
	"crypto/md5"
	"crypto/sha1"
)

func main(){
	testFile := "123.txt"
	infile, inerr := os.Open( testFile )
	if inerr != nil {
		fmt.Println(err)
		return
	}

	md5Code := md5.New()
	io.Copy( md5Code, infile )
	fmt.Print("%x  	a
