package main

import (
	"fmt"
)

func xorCipher(plainText, key string) string {
	cipherText := ""
	keyLen := len(key)

	for i, char := range plainText {
		cipherText += string(char ^ rune(key[i%keyLen]))
	}

	return cipherText
}
func main() {
	message := "Hello World!"
	key := "KEY"
	encryptedMessage := xorCipher(message, key)
	fmt.Println("Crypted Message:", encryptedMessage)
}
