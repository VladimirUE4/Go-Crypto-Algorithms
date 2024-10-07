package main

import (
	"fmt"
	"strings"
)

var rotor1 = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
var rotor2 = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
var rotor3 = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
var reflector = "YRUHQSLDPXNGOKMIEBFZCWVJAT"
var rotorPositions = []int{0, 0, 0}

func shiftLetter(letter byte, shift int) byte {
	return byte(((int(letter-'A') + shift) % 26) + 'A')
}
func unshiftLetter(letter byte, shift int) byte {
	return byte(((int(letter-'A') - shift + 26) % 26) + 'A')
}
func applyRotor(rotor string, letter byte, position int) byte {
	shifted := shiftLetter(letter, position)
	return rotor[shifted-'A']
}
func reverseRotor(rotor string, letter byte, position int) byte {
	index := strings.IndexByte(rotor, shiftLetter(letter, position))
	return unshiftLetter(byte('A'+index), position)
}
func rotateRotors() {
	rotorPositions[0]++
	if rotorPositions[0] == 26 {
		rotorPositions[0] = 0
		rotorPositions[1]++
		if rotorPositions[1] == 26 {
			rotorPositions[1] = 0
			rotorPositions[2]++
			if rotorPositions[2] == 26 {
				rotorPositions[2] = 0
			}
		}
	}
}
func encodeLetter(letter byte) byte {
	if letter < 'A' || letter > 'Z' {
		return letter
	}
	rotateRotors()
	l := applyRotor(rotor1, letter, rotorPositions[0])
	l = applyRotor(rotor2, l, rotorPositions[1])
	l = applyRotor(rotor3, l, rotorPositions[2])
	l = reflector[l-'A']
	l = reverseRotor(rotor3, l, rotorPositions[2])
	l = reverseRotor(rotor2, l, rotorPositions[1])
	l = reverseRotor(rotor1, l, rotorPositions[0])
	return l
}
func encodeMessage(message string) string {
	var result strings.Builder
	message = strings.ToUpper(message)

	for i := 0; i < len(message); i++ {
		result.WriteByte(encodeLetter(message[i]))
	}
	return result.String()
}
func main() {
	message := "Hello World!"
	encoded := encodeMessage(message)
	fmt.Println("Crypted Message :", encoded)
}
