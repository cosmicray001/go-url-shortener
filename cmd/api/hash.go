package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"
)

func (app *application) generateShortUrl(actualUrl string) (string, error) {
	for i := 0; i < 5; i++ {
		text := actualUrl + time.Now().String()
		hashedStr := GetMD5Hash(text)[:8]
		ok, err := app.urlBank.CheckExistUrl(hashedStr)
		if err != nil {
			return "", err
		}
		if !ok {
			return hashedStr, nil
		}
	}
	return "", errors.New("limit finished")
}

func GetMD5Hash(text string) string {
	if text == "" {
		return ""
	}
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
