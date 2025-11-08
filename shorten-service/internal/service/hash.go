package service

import (
	"crypto/sha256"
	"encoding/base64"
	"time"
)

func createShortKey(longURL string) string {
	hash := sha256.Sum256([]byte(longURL))
	encoded := base64.RawURLEncoding.EncodeToString(hash[:])
	return encoded[:6]
}

func createSaltedShortKey(longURL string) string {
	salt := time.Now().UnixNano()
	input := longURL + "|" + string(rune(salt))

	hash := sha256.Sum256([]byte(input))
	encoded := base64.RawURLEncoding.EncodeToString(hash[:])
	return encoded[:6]

}
