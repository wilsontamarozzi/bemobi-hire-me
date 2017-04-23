package helpers

import (
	"bytes"
	"math/rand"
	"time"
)

func RandomString(length int) string {
	characters := "abcdefghijklmnopqrstuvxyzABCDEFGHIJKLMNOPQRSTUVXYZ1234567890"
	maxSize := len(characters)
	rand.Seed(time.Now().UTC().UnixNano())
	var secret bytes.Buffer

	for i := 1; i <= length; i++ {
		secret.WriteString(string(characters[rand.Intn(maxSize)]))
	}

	return secret.String()
}

func Contains(array []string, text string) bool {
	for _, item := range array {
		if item == text {
			return true
		}
	}
	return false
}
