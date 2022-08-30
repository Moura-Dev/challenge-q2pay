package services

import (
	"crypto/sha256"
	"fmt"
)

func SHA256ENCODER(text string) string {
	strEncoder := sha256.Sum256([]byte(text))

	return fmt.Sprintf("%x", strEncoder)
}
