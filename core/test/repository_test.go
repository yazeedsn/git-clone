package test

import (
	"github/yazeedsn/gogit/core"
	"os"
	"path/filepath"
	"testing"
)

func TestNotExistRepo(t *testing.T) {
	path := "TestNotExistRepoTemp"
	_, err := os.Stat(path)
	if os.IsExist(err) {
		os.Remove(path)
	}
	repo := new(core.Repository)
	err = repo.New(path, false)
	if err == nil {
		t.Fatal("Falid to recongnize the non-existance of a repository in the working dir")
	}
}

func TestCreateRepo(t *testing.T) {
	path := "TestCreateRepoTemp"
	defer os.RemoveAll(path)
	err := os.Mkdir(path, 0755)
	if err != nil {
		t.Fatal(err)
	}

	repo := new(core.Repository)
	err = repo.New(path, true)
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.Stat(filepath.Join(path, core.GIT_DIR_NAME))
	if err != nil {
		t.Fatal(err)
	}
}

func TestExistingRepo(t *testing.T) {
	path := "TestExistingRepoTemp"
	defer os.RemoveAll(path)
	err := os.MkdirAll(filepath.Join(path, core.GIT_DIR_NAME), 0755)
	if err != nil {
		t.Fatal(err)
	}

	repo := new(core.Repository)
	err = repo.New(path, false)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddFile(t *testing.T) {
	path := "TestCreateRepoTemp"
	temp_file := "dat1"
	defer os.RemoveAll(path)
	err := os.Mkdir(path, 0755)
	if err != nil {
		t.Fatal(err)
	}

	d1 := []byte("hello\ngo\n")
	if err := os.WriteFile(filepath.Join(path, temp_file), d1, 0644); err != nil {
		t.Fatal(err)
	}

	repo := new(core.Repository)
	if err := repo.New(path, true); err != nil {
		t.Fatal(err)
	}

	if err := repo.Add(temp_file); err != nil {
		t.Fatal(err)
	}

}

func TestAddDir(t *testing.T) {
	path := "TestCreateRepoTemp"
	temp_dir := "temp_dir"
	temp_file := filepath.Join(temp_dir, "dat1")
	defer os.RemoveAll(path)
	err := os.MkdirAll(filepath.Join(path, temp_dir), 0755)
	if err != nil {
		t.Fatal(err)
	}

	d1 := []byte("hello\ngo\n")
	if err := os.WriteFile(filepath.Join(path, temp_file), d1, 0644); err != nil {
		t.Fatal(err)
	}

	repo := new(core.Repository)
	if err := repo.New(path, true); err != nil {
		t.Fatal(err)
	}

	if err := repo.Add(temp_dir); err != nil {
		t.Fatal(err)
	}

}
