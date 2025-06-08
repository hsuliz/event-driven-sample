package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func Encode(raw []byte) string {
	checksum := sha256.Sum256(raw)
	return hex.EncodeToString(checksum[:])
}
