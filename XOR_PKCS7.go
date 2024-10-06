package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

const blockSize = 16

func xorBlock(plainText, key []byte) []byte {
	cipherText := make([]byte, len(plainText))
	for i := range plainText {
		cipherText[i] = plainText[i] ^ key[i]
	}
	return cipherText
}
func pkcs7Pad(data []byte) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}
func pkcs7Unpad(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("empty data")
	}
	padding := int(data[length-1])
	if padding > blockSize || padding == 0 {
		return nil, fmt.Errorf("padding PKCS7 invalide")
	}
	for i := 0; i < padding; i++ {
		if data[length-1-i] != byte(padding) {
			return nil, fmt.Errorf("padding PKCS7 incorrect")
		}
	}
	return data[:length-padding], nil
}
func main() {
	plainText := []byte("Hello World!")
	paddedText := pkcs7Pad(plainText)
	fmt.Println("Data with Padding (hex):", fmt.Sprintf("%x", paddedText))
	key := []byte("SimpleAESKey1234")
	if len(paddedText) != 16 || len(key) != 16 {
		fmt.Println("The data and key should be 16 bytes.")
		return
	}
	cipherText := xorBlock(paddedText, key)
	fmt.Println("Crypted Data (hex):", hex.EncodeToString(cipherText))
	decryptedText := xorBlock(cipherText, key)
	fmt.Println("Decrypted Data with padding:", string(decryptedText))
	unpaddedText, err := pkcs7Unpad(decryptedText)
	if err != nil {
		fmt.Println("Error while depadding:", err)
		return
	}
	fmt.Println("Decrypted Data without padding:", string(unpaddedText))
}
