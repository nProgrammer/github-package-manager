package tools

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncSHA(str string) string {
	hash := sha256.Sum256([]byte(str))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}
