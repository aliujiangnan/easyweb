package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

var privKey = []byte("123456789")

func padding(src []byte,blocksize int) []byte {
	padnum:=blocksize-len(src)%blocksize
	pad:=bytes.Repeat([]byte{byte(padnum)},padnum)
	return append(src,pad...)
}
 
func unpadding(src []byte) []byte {
	n:=len(src)
	unpadnum:=int(src[n-1])
	return src[:n-unpadnum]
}
 
func DncryptAES(src interface{}) interface{} {
	block,_:=aes.NewCipher(privKey)
	srcByte:=padding(src.([]byte),block.BlockSize())
	blockmode:=cipher.NewCBCEncrypter(block,privKey)
	blockmode.CryptBlocks(srcByte,srcByte)
	return src
}
 
func DecryptAES(src interface{}) interface{} {
	block,_:=aes.NewCipher(privKey)
	blockmode:=cipher.NewCBCDecrypter(block,privKey)
	srcByte := src.([]byte)
	blockmode.CryptBlocks(srcByte,srcByte)
	src=unpadding(srcByte)
	return src
}