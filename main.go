package main

import (
	"crypto/rand"
	"fmt"
)

// Define the runes to be used in the random string
var runes = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()-_=+[]{},.<>?;:|")

func main() {
	// Create a byte slice of length 32
	randomBytes := make([]byte, 32)

	// Fill the byte slice with random bytes
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error generating random bytes:", err)
		return
	}

	// Create a random string from the random bytes using allowed runes
	var randomString string
	for _, b := range randomBytes {
		idx := int(b) % len(runes)
		randomString += string(runes[idx])
	}

	fmt.Println("Random 32-byte string:", randomString)
}
