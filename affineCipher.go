package main

import (
	"fmt"
)

func affineCipher(plainText string, a, b int) string {
	cipherText := ""
	for _, char := range plainText {
		if char >= 'A' && char <= 'Z' {
			cipherText += string(((a*(int(char)-'A') + b) % 26) + 'A')
		} else if char >= 'a' && char <= 'z' {
			cipherText += string(((a*(int(char)-'a') + b) % 26) + 'a')
		} else {
			cipherText += string(char)
		}
	}
	return cipherText
}
func main() {
	message := "Hello World!"
	a, b := 5, 8
	encryptedMessage := affineCipher(message, a, b)
	fmt.Println("Crypted Message:", encryptedMessage)
}
