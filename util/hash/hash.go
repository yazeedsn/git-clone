package hash

import (
	"crypto/sha1"
	"hash"
	"io"
	"os"
)

func NewHasher() hash.Hash {
	return sha1.New()
}

func HashFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	h := sha1.New()
	if _, err = io.Copy(h, file); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}
