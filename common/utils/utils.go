package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

func GenCode(length int) string {
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = byte(rand.Intn(10)) + '0'
	}
	return string(code)
}

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
