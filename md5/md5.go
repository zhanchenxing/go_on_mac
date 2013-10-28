package main

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
)

func main(){
	TestString :="Hi, pandaman!"
	Md5Inst := md5.New()
	Md5Inst.Write([]byte( TestString ) )
	Result := Md5Inst.Sum( []byte("") )
	fmt.Printf("%x\n\n", Result)

	sha1Code := sha1.New()
	sha1Code.Write( []byte( TestString ) )
	Result = sha1Code.Sum( []byte("") )
	fmt.Printf("%x\n\n", Result )
}
