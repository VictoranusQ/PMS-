package utils

import (
	"crypto/md5"
	"encoding/hex"
)

type Crypto struct {
}

func (c *Crypto) MD5(s string) string {
	h := md5.New()
	h.Write([]byte(s))

	return hex.EncodeToString(h.Sum(nil))
}
