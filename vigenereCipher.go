package main

import (
	"fmt"
)

func vigenereCipher(plainText, key string) string {
	cipherText := ""
	keyLength := len(key)

	for i, char := range plainText {
		shift := rune(key[i%keyLength] - 'A')

		if char >= 'A' && char <= 'Z' {
			cipherText += string((char-'A'+shift)%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			cipherText += string((char-'a'+shift)%26 + 'a')
		} else {
			cipherText += string(char)
		}
	}
	return cipherText
}
func main() {
	message := "Hello World!"
	key := "KEY"
	encryptedMessage := vigenereCipher(message, key)
	fmt.Println("Crypted Message:", encryptedMessage)
}
