package hash

import (
	"fmt"
	"io"
	"crypto/md5"
	"code.google.com/p/go.crypto/ripemd160"
	"crypto/sha256"
	"github.com/miie/msjuvi/logger"
)

func MD5Hash(s string) (string, error) {
	h := md5.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		logger.LogWarning("error when io.WriteString for md5 hash. error: ", err)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func Ripemd160Hash(s string) (string, error) {
	h := ripemd160.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		logger.LogWarning("error when io.WriteString for ripemd160 hash. error: ", err)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func Sha256Hash(s string) (string, error) {
	h := sha256.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		logger.LogWarning("error when io.WriteString for sha256 hash. error: ", err)
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
