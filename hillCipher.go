package main

import (
	"fmt"
	"strings"
)

func charToInt(c rune) int {
	return int(c - 'A')
}
func intToChar(n int) rune {
	return rune(n + 'A')
}
func hillCipher(plainText string, keyMatrix [2][2]int) string {
	plainText = strings.ToUpper(plainText)
	plainText = strings.ReplaceAll(plainText, " ", "")
	cipherText := ""
	if len(plainText)%2 != 0 {
		plainText += "X"
	}
	for i := 0; i < len(plainText); i += 2 {
		firstLetter := charToInt(rune(plainText[i]))
		secondLetter := charToInt(rune(plainText[i+1]))
		c1 := (keyMatrix[0][0]*firstLetter + keyMatrix[0][1]*secondLetter) % 26
		c2 := (keyMatrix[1][0]*firstLetter + keyMatrix[1][1]*secondLetter) % 26
		cipherText += string(intToChar(c1)) + string(intToChar(c2))
	}

	return cipherText
}
func main() {
	keyMatrix := [2][2]int{
		{3, 3},
		{2, 5},
	}
	message := "Hello World!"
	encryptedMessage := hillCipher(message, keyMatrix)
	fmt.Println("Crypted Message:", encryptedMessage)
}
