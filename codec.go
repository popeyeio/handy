package handy

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

func BuildMD5(parts ...string) string {
	hash := md5.New()
	for _, part := range parts {
		hash.Write(Str2Bytes(part))
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func BuildHMAC1(key, data string) []byte {
	return BuildHMAC(key, data, sha1.New)
}

func BuildHMAC256(key, data string) []byte {
	return BuildHMAC(key, data, sha256.New)
}

type hashFunc func() hash.Hash

func BuildHMAC(key, data string, hf hashFunc) []byte {
	hash := hmac.New(hf, Str2Bytes(key))
	hash.Write(Str2Bytes(data))
	return hash.Sum(nil)
}

func EncodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
