package main

import (
	"fmt"
)

func caesarCipher(plainText string, shift int) string {
	cipherText := ""
	for _, char := range plainText {
		if char >= 'A' && char <= 'Z' {
			cipherText += string((char-'A'+rune(shift))%26 + 'A')
		} else if char >= 'a' && char <= 'z' {
			cipherText += string((char-'a'+rune(shift))%26 + 'a')
		} else {
			cipherText += string(char)
		}
	}
	return cipherText
}
func main() {
	message := "Hello World!"
	shift := 26
	encryptedMessage := caesarCipher(message, shift)
	fmt.Println("Crypted Message:", encryptedMessage)
}
