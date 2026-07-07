package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

func SHA256(text string) string {
	hashByte := sha256.Sum256([]byte(text))

	return fmt.Sprintf("%x", hashByte)
}

func MD5(text string) string {
	hashBytes := md5.Sum([]byte(text))
	return fmt.Sprintf("%x", hashBytes)
}
