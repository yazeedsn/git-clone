package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Repository struct {
	WorkingDir string
	GitDir     string
}

func (r *Repository) New(workingDir string, create bool) error {
	info, err := os.Stat(workingDir)
	if os.IsNotExist(err) || !info.IsDir() {
		log.Print("dir doesn't exist")
		return err
	} else if err != nil {
		return err
	}

	gitDir := filepath.Join(workingDir, GIT_DIR_NAME)
	info, err = os.Stat(gitDir)
	if err == nil && info.IsDir() {
		log.Print("repository already exists")
	} else if create {
		err = os.Mkdir(gitDir, 0775)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	r.WorkingDir = workingDir
	r.GitDir = gitDir
	return nil
}

func (r *Repository) Add(rPath string) error {
	path := filepath.Join(r.WorkingDir, rPath)
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if info.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			ePath := filepath.Join(rPath, entry.Name())
			r.Add(ePath)
		}
	} else {
		addFile(r, rPath)
	}

	return nil
}

func addFile(r *Repository, rFilePath string) error {
	var b Blob
	if err := b.New(r, rFilePath); err != nil {
		return err
	}

	if err := writeIndex(r, b.Hash(), rFilePath); err != nil {
		return err
	}

	return nil
}

func writeIndex(r *Repository, hash string, rPath string) error {
	indexPath := filepath.Join(r.GitDir, "index")
	file, err := os.OpenFile(indexPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	row := fmt.Sprintf("%s %s\n", hash, rPath)
	if _, err := file.WriteString(row); err != nil {
		return err
	}

	return nil
}
