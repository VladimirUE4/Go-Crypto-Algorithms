package main

import (
	"bytes"
	"fmt"
)

const blockSize = 16

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
	fmt.Println("Data with padding (hex):", fmt.Sprintf("%x", paddedText))
	unpaddedText, err := pkcs7Unpad(paddedText)
	if err != nil {
		fmt.Println("Error during depadding:", err)
		return
	}
	fmt.Println("Data after depadding:", string(unpaddedText))
}
