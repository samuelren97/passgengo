package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashSHA256(b []byte) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
