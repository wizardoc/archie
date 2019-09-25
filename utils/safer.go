package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
)

func Hash(str string) string {
	SHA1 := sha1.New()
	SHA1.Write([]byte(str))

	return hex.EncodeToString(SHA1.Sum(getSalt()))
}

func getSalt() []byte {
	return []byte("Wizard Salt")
}

/** 生成验证码 */
func CreateVerifyCode() string {
	var verifyCodes []string
	const VERIFY_CODE_COUNT = 6

	for i := 0; i < VERIFY_CODE_COUNT; i++ {
		verifyCodes = append(verifyCodes, strconv.Itoa(rand.Intn(10)))
	}

	return strings.Join(verifyCodes, "")
}
