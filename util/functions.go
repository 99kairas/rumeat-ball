package util

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"strings"
)

func GenerateOrderID(prefix string) (string, error) {
	bytes := make([]byte, 4)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", err
	}

	return prefix + "-" + strings.ToUpper(hex.EncodeToString(bytes)), nil
}
