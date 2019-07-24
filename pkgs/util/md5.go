package util

import (
	"crypto/md5"
	"encoding/hex"
)

// string to md5
// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
