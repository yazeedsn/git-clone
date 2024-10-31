package core

import (
	"encoding/hex"
	"fmt"
	"github/yazeedsn/gogit/util/compress"
	"github/yazeedsn/gogit/util/hash"
	"os"
	"path/filepath"
	"strings"
)

type Tree struct {
	repository *Repository
	hash       string
	Children   []PathObject
}

func (t Tree) Repository() *Repository {
	return t.repository
}

func (t Tree) Hash() string {
	return t.hash
}

func (t Tree) Type() string {
	return "Tree"
}

func (t *Tree) New(r *Repository, relativePath string) error {
	t.repository = r
	rootTreePath := filepath.Join(r.WorkingDir, relativePath)
	entries, err := os.ReadDir(rootTreePath)
	t.Children = make([]PathObject, len(entries))
	if err != nil {
		return err
	}

	var builder strings.Builder
	hasher := hash.NewHasher()
	for _, entry := range entries {
		var child PathObject
		child.Path = filepath.Join(relativePath, entry.Name())
		if entry.IsDir() {
			var tree Tree
			if err := tree.New(r, child.Path); err != nil {
				return err
			}
			child.Object = tree
		} else {
			var blob Blob
			if err := blob.New(r, child.Path); err != nil {
				return err
			}
			child.Object = blob
		}
		t.Children = append(t.Children, child)
		_, err = hasher.Write(append([]byte(child.Hash()), []byte(child.Path)...))
		if err != nil {
			return err
		}

		_, err = builder.WriteString(fmt.Sprintf("%4s %20x %s\n", child.Type(), child.Hash(), child.Path))
		if err != nil {
			return err
		}
	}
	t.hash = hex.EncodeToString(hasher.Sum(nil))

	if Exists(t) {
		return nil
	}

	if err := makeTreeFile(t, builder.String()); err != nil {
		return err
	}

	return nil
}

func makeTreeFile(t *Tree, content string) error {
	root, filePath, err := GetReference(t)
	if err != nil {
		return err
	}

	if _, err := os.Stat(root); os.IsNotExist(err) {
		if err = os.MkdirAll(root, 0755); err != nil {
			return err
		}
	}

	err = compress.CompressContent(content, filePath)
	if err != nil {
		return err
	}
	return nil
}
