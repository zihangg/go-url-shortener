package helpers

import (
	"net/url"
	"os"
	"strings"
)

const (
	alphabet      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateEncodedURL(number uint64) string {
	length := len(alphabet)
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}

	return encodedBuilder.String()
}


func NonInternalDomain(urlString string) bool {
	if urlString == os.Getenv("DOMAIN") {
		return false 
	}

	u, err := url.Parse(urlString)
	if err != nil {
		return false
	}

	if u.Host == os.Getenv("DOMAIN") {
		return false
	}

	return true
}
