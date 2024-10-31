package core

import (
	"fmt"
	"os"
	"path/filepath"
)

type Object interface {
	Repository() *Repository
	Hash() string
	Type() string
}

type PathObject struct {
	Object
	Path string
}

func GetReference(object Object) (root string, file string, err error) {
	if len(object.Hash()) < 20 {
		return "", "", fmt.Errorf("Object is does not have a valid hash (%s)", object.Hash())
	}

	root = filepath.Join(object.Repository().GitDir, "Objects", object.Hash()[:2])
	file = filepath.Join(root, object.Hash()[2:len(object.Hash())])
	return root, file, nil
}

func Exists(object Object) bool {
	root, file, err := GetReference(object)
	if err != nil {
		return false
	}
	if _, err := os.Stat(root); os.IsNotExist(err) {
		return false
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
