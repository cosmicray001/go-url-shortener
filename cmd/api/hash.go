package main

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(text string) string {
	if text == "" {
		return ""
	}
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
