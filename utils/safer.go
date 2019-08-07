package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Hash(str string) string {
	SHA1 := sha1.New()
	SHA1.Write([]byte(str))

	return hex.EncodeToString(SHA1.Sum(getSalt()))
}

func getSalt() []byte {
	return []byte("Wizard Salt")
}
