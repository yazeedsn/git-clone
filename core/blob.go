package core

import (
	"encoding/hex"
	"fmt"
	"github/yazeedsn/gogit/util/compress"
	"github/yazeedsn/gogit/util/hash"
	"os"
	"path/filepath"
)

type Blob struct {
	repository *Repository
	hash       string
}

func (b Blob) Repository() *Repository {
	return b.repository
}

func (b Blob) Hash() string {
	return b.hash
}

func (b Blob) Type() string {
	return "Blob"
}

func (b *Blob) New(Repo *Repository, RelativePath string) error {
	// assign b a reporsitory
	b.repository = Repo

	// open the file of the blob
	path := filepath.Join(b.Repository().WorkingDir, RelativePath)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return err
	}
	if info.IsDir() {
		return fmt.Errorf("%s is a directory, not a file", RelativePath)
	}

	// hash and assign hash to blob
	hash, err := hash.HashFile(path)
	if err != nil {
		return err
	}
	b.hash = hex.EncodeToString(hash)

	// if blob is already stored in objects return
	if Exists(b) {
		return nil
	}

	// create and store the blob in the appropriate reference in objects
	if err := makeBlobFile(b, path); err != nil {
		return err
	}

	return nil
}

func makeBlobFile(b *Blob, filePath string) error {
	root, file, err := GetReference(b)
	if err != nil {
		return err
	}

	if _, err := os.Stat(root); os.IsNotExist(err) {
		if err = os.MkdirAll(root, 0755); err != nil {
			return err
		}
	}

	if err := compress.CompressFile(filePath, file); err != nil {
		return err
	}

	return nil
}
