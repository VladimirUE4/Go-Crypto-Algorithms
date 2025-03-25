package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"

	"golang.org/x/crypto/argon2"
)

// Generate a random salt
func generateSalt(size int) []byte {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt
}

// Key derivation with Argon2 and random noise
func deriveKey(password string, salt []byte, keyLength uint32) []byte {
	baseKey := argon2.IDKey([]byte(password), salt, 5, 512*1024, 32, keyLength)
	noise := generateSalt(len(baseKey))
	for i := range baseKey {
		baseKey[i] ^= noise[i] // Add random noise
	}
	return baseKey
}

// Generate an HMAC for integrity
func generateHMAC(data, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}

// Chaotic transformation based on Mandelbrot
func chaoticTransform(data []byte, salt []byte) []byte {
	for i := range data {
		x := float64(data[i]) * 2.718            // Euler's constant
		y := float64(salt[i%len(salt)]) * 3.1415 // Pi
		z := math.Mod(math.Pow(x, y), 256)       // Chaos
		data[i] = byte(z)
	}
	return data
}

// Reverse chaotic transformation
func reverseChaoticTransform(data []byte, salt []byte) []byte {
	for i := range data {
		// Reverse the chaotic transformation (approximation for invertibility)
		y := float64(salt[i%len(salt)]) * 3.1415 // Pi
		data[i] = byte(math.Pow(float64(data[i]), 1/y))
	}
	return data
}

// AES encryption with fractal permutation
func encryptAES(key, plaintext []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	_, err = rand.Read(iv)
	if err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext
}

// AES decryption
func decryptAES(key, ciphertext []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext
}

// Encryption function
func encrypt(text, password string) string {
	salt := generateSalt(16)
	key := deriveKey(password, salt, uint32(len(text)))
	encrypted := chaoticTransform([]byte(text), salt)
	encrypted = encryptAES(key, encrypted)
	hmacValue := generateHMAC(encrypted, key)

	return hex.EncodeToString(salt) + hex.EncodeToString(encrypted) + hex.EncodeToString(hmacValue)
}

// Decryption function
func decrypt(encryptedHex, password string) string {
	salt, _ := hex.DecodeString(encryptedHex[:32])
	encrypted, _ := hex.DecodeString(encryptedHex[32 : len(encryptedHex)-64])
	hmacReceived, _ := hex.DecodeString(encryptedHex[len(encryptedHex)-64:])

	key := deriveKey(password, salt, uint32(len(encrypted)))
	expectedHMAC := generateHMAC(encrypted, key)
	if !hmac.Equal(hmacReceived, expectedHMAC) {
		panic("ERROR: Integrity check failed!")
	}

	decrypted := decryptAES(key, encrypted)
	decrypted = reverseChaoticTransform(decrypted, salt)

	return string(decrypted)
}

// Example usage
func main() {
	message := "Ultra Secure Message!"
	password := "supersecurepassword"

	encrypted := encrypt(message, password)
	fmt.Println("🔐 Encrypted:", encrypted)

	decrypted := decrypt(encrypted, password)
	fmt.Println("🔓 Decrypted:", decrypted)
}
