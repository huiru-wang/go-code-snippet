package encryptutil

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// Sha1 sha1签名
func Sha1(content string) string {
	hash := sha1.New()
	hash.Write([]byte(content))
	return hex.EncodeToString(hash.Sum(nil))
}

// HmacSha256 HmacSha256签名
func HmacSha256(key string, content string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(content))
	return hex.EncodeToString(hash.Sum(nil))
}
