package aesrw

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
)

type AESWriter struct {
	writeTo io.Writer

	buff   []byte
	cbcenc cipher.BlockMode
}

func MakeAESWriter(toWriteTo io.Writer) (*AESWriter, error) {
	key_text := "1234567890123456"
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		return nil, err
	}

	r := &AESWriter{}
	r.buff = make([]byte, 0, 1024*4)
	r.writeTo = toWriteTo
	r.cbcenc = cipher.NewCBCEncrypter(c, commonIV)
	return r, nil
}

func (r *AESWriter) Write(b []byte) (int, error) {
	r.buff = append(r.buff, b...)
	canDo := len(r.buff) / 16 * 16

	i := 0
    var err error
	for ; i < canDo; i += 16 {
		r.cbcenc.CryptBlocks(r.buff[i:i+16], r.buff[i:i+16])
		_, err := r.writeTo.Write(r.buff[i : i+16])
		if err != nil {
			break
		}
	}
	r.buff = r.buff[i:]
	return i, err
}
