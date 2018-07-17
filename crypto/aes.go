package main

import (
	"crypto/aes"
	"fmt"
	"bytes"
	"crypto/cipher"
)

// http://www.jianshu.com/p/9c1c8958b279
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func main(){
	key := "qwertyuiouqwerty"
	data := "ffffffffffffffhhhhhhhhhhhfhfhfhffhfhfhfh"
	res,err := AesEncrypt([]byte(data), []byte(key))
	if err!=nil{
		fmt.Println("error:", err)}
	println("res:", string(res))
	res2,err := AesDecrypt(res, []byte(key))
	if err!=nil{
		fmt.Println("error:", err)}
	println("res2:", string(res2))
	

}
