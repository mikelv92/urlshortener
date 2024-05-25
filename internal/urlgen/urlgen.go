package urlgen

import (
	"math/rand"
)

func GenerateURL() string {
	return generateRandomString(8)
}

func generateRandomString(length int) string {
	chars := "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	var s []byte

	for i := 0; i < length; i++ {
		index := rand.Intn(62)
		s = append(s, chars[index])
	}

	return string(s)
}
