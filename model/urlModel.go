package model

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
)

type URLData struct {
	OriginalURL  string
	ShortenedURL string
}

func (u *URLData) Validate() error {
	_, err := url.ParseRequestURI(u.OriginalURL)
	if err != nil {
		return errors.New("invalid URL format")
	}
	return nil
}

func GenerateHash(url string, counter int64) string {
	uniqueString := fmt.Sprintf("%s%d", url, counter)
	hashBytes := sha256.Sum256([]byte(uniqueString))
	encodedHash := base64.URLEncoding.EncodeToString(hashBytes[:])
	return encodedHash[:8]
}
