package vutils

import (
	"crypto/rand"
	"encoding/base64"
)

//RandomBytes return random bytes
func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

//RandomBase64 generate random Base64 string
func RandomBase64(n int) (string, error) {
	b, err := RandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}
