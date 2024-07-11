package models

import (
	"crypto/rand"
	"encoding/hex"
)

const attempts = 16

func generateRandom(size int) []byte {
	buffer := make([]byte, size)

	for i := 0; i < attempts; i++ {
		_, err := rand.Read(buffer)
		if err == nil {
			break
		}
	}

	return buffer
}

func GenerateRandomHexString(size int) string {
	return hex.EncodeToString(generateRandom(size))
}
