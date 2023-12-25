package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"os"

	"github.com/pkg/errors"
)

func GetCheckSum(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", errors.Wrap(err, "os.ReadFile")
	}

	var h = sha256.New()
	if _, err := h.Write(data); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
