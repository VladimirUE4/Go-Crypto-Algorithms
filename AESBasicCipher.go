// Example of the AES Basic implementation in Go using XOR encryption with Padding. without SubBytes, ShiftRows, MixColumns, only one round.

package main

import (
	"encoding/hex"
	"fmt"
)

func xorBlock(plainText, key []byte) []byte {
	cipherText := make([]byte, len(plainText))
	for i := range plainText {
		cipherText[i] = plainText[i] ^ key[i]
	}
	return cipherText
}
func padText(text string) string {
	if len(text) < 16 {
		padding := 16 - len(text)
		for i := 0; i < padding; i++ {
			text += "X"
		}
	}
	return text
}
func main() {
	plainText := "Hello World!"
	plainText = padText(plainText)
	key := []byte("SimpleAESKey1234")
	if len(plainText) != 16 || len(key) != 16 {
		fmt.Println("The text and the key should be 16 bytes.")
		return
	}
	cipherText := xorBlock([]byte(plainText), key)
	fmt.Println("Crypted Message (hex):", hex.EncodeToString(cipherText))
	decryptedText := xorBlock(cipherText, key)
	fmt.Println("Decrypted Message (with padding):", string(decryptedText))
}
