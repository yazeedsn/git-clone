package hash

import (
	"bytes"
	"os"
	"testing"
)

func TestMatchingHash(t *testing.T) {
	path := "MatchingHashTesttemp"
	var testData = []byte{1, 2, 3, 4, 5, 6}
	if err := os.WriteFile(path, testData, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)
	hash1, err1 := HashFile(path)
	hash2, err2 := HashFile(path)
	if err1 != nil || err2 != nil {
		t.Fatal("could not hash file.")
	}
	if !bytes.Equal(hash1, hash2) {
		t.Fatal("falid to generate a consistant hash for the same file")
	}
	if err := os.Remove(path); err != nil {
		t.Fatal()
	}
}

func TestMatchingHashTwoFiles(t *testing.T) {
	path := "MatchingHashTestTwoFilestemp"
	var testData = []byte{1, 2, 3, 4, 5, 6}
	if err := os.WriteFile(path, testData, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)
	hash1, err := HashFile(path)
	if err != nil {
		t.Fatal(err)
	}

	path2 := "MatchingHashTestTwoFilestemp2"
	if err := os.WriteFile(path2, testData, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path2)
	hash2, err := HashFile(path2)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(hash1, hash2) {
		t.Fatal("falid to generate a consistant hash for two files containing the same content")
	}
}

func TestUnmatchingHash(t *testing.T) {
	path := "UnMatchingHashTestemp"
	var testData = []byte{1, 2, 3, 4, 5, 6}
	if err := os.WriteFile(path, testData, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path)

	hash1, err := HashFile(path)
	if err != nil {
		t.Fatal(err)
	}

	path2 := "UnMatchingHashTesttemp2"
	var testData2 = []byte{1, 2, 3, 4, 5, 1}
	if err := os.WriteFile(path2, testData2, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.Remove(path2)

	hash2, err := HashFile(path2)
	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(hash1, hash2) {
		t.Fatal("falid to distinct hashs for two files containing different content")
	}
}
