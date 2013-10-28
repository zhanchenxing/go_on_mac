package main

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"os"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06,
	0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}


func main(){
	fmt.Println("Hello, world")

	plaintext := []byte( "My name is Astaxie90123456789012")

	if len(os.Args) > 1 {
		plaintext = []byte( os.Args[1] )
	}

	// aes 
	key_text := "1234567890123456"
	if len(os.Args) > 2 {
		key_text = os.Args[2]
	}

	fmt.Println( "key_text", key_text )

	c, err := aes.NewCipher( []byte(key_text) )
	if err != nil {
		fmt.Println("error on aes.NewCipher:", err)
		os.Exit(-1)
	}

	cfb := cipher.NewCFBEncrypter( c, commonIV )
	ciphertext := make( []byte, len(plaintext) )
	cfb.XORKeyStream( ciphertext, plaintext )
	fmt.Printf("%s=>%x\n", plaintext, ciphertext )

	//
	cfbdec := cipher.NewCFBDecrypter( c, commonIV )
	plaintextCopy := make( []byte, len(plaintext) )
	cfbdec.XORKeyStream( plaintextCopy, ciphertext )
	fmt.Printf("%x=>%s\n", ciphertext, plaintextCopy )

    cbcenc := cipher.NewCBCEncrypter( c, commonIV )
    dest1 := make( []byte, len(plaintext) )
    cbcenc.CryptBlocks( dest1, plaintext )
    fmt.Printf("CBC mode %x=>%x\n", plaintext, dest1 )
    dest2 := make( []byte, len(plaintext) )
    cbcenc.CryptBlocks( dest2, plaintext )
    fmt.Printf("CBC mode %x=>%x\n", plaintext, dest2 )

    cbcdec := cipher.NewCBCDecrypter( c, commonIV )

    decDest1 := make( []byte, len(plaintext) )
    cbcdec.CryptBlocks( decDest1, dest1 )
    fmt.Printf("CBC mode dec %x=>%x\n", dest1, decDest1 )
    decDest2 := make( []byte, len(plaintext) )
    cbcdec.CryptBlocks( decDest2, dest2 )
    fmt.Printf("CBC mode dec %x=>%x\n", dest2, decDest2 )

}



