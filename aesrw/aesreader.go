package aesrw

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06,
	0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

type AESReader struct {
	readFrom io.Reader

	buff   []byte
	cbcdec cipher.BlockMode
    readTo []byte
}

func MakeAESReader(toReadFrom io.Reader) (*AESReader, error) {
	key_text := "1234567890123456"
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		return nil, err
	}

	ret := &AESReader{}
	ret.buff = make([]byte, 0, 1024*4)
	ret.readTo = make([]byte, 1024*4)
	ret.readFrom = toReadFrom
	ret.cbcdec = cipher.NewCBCDecrypter(c, commonIV)
	return ret, nil
}

func Min( p1, p2 int ) int {
    if p1 < p2 {
        return p1
    }
    return p2
}

func (r *AESReader) processLeft(to []byte) (n int) {
	canDo := Min( len(r.buff) / 16 * 16, len(to)/16*16 )

    i := 0
	for ; i < canDo; i += 16 {
		r.cbcdec.CryptBlocks(to[i:i+16], r.buff[i:i+16])
	}

    r.buff = r.buff[i:]
	return i
}

func (r *AESReader) Read(b []byte) (n int, err error) {
	l, err := r.readFrom.Read(r.readTo)
    r.buff = append( r.buff, r.readTo[:l]...)
	return r.processLeft(b), err
}
