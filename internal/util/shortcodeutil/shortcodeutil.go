package shortcodeutil

import (
	"math/rand"
	"regexp"
)

const (
	characterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	shortcodeRegex = "^[0-9a-zA-Z_]{6}$"
)

func GenerateShortcode(n int) string {
	b := make([]byte, n)
    for i := range b {
        b[i] = characterBytes[rand.Intn(len(characterBytes))]
    }
    return string(b)
}

func ValidateShortcode(s string) bool {
	isValid, _ := regexp.MatchString(shortcodeRegex, s)
	return isValid
}
