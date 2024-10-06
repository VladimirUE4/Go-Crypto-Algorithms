package main

import (
	"fmt"
)

func scytaleCipher(plainText string, key int) string {
	plainText = removeSpaces(plainText)
	length := len(plainText)
	cipherText := make([]byte, length)
	for i := 0; i < length; i++ {
		col := i % key
		row := i / key
		cipherText[i] = plainText[row+col*(length/key)]
	}
	return string(cipherText)
}
func removeSpaces(text string) string {
	result := ""
	for _, char := range text {
		if char != ' ' {
			result += string(char)
		}
	}
	return result
}
func main() {
	message := "Hello World!"
	key := 4
	encryptedMessage := scytaleCipher(message, key)
	fmt.Println("Crypted Messageé:", encryptedMessage)
}
