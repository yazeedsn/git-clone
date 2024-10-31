package compress

import (
	"compress/zlib"
	"io"
	"os"
)

func CompressContent(content string, targetPath string) error {
	file, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := zlib.NewWriter(file)
	defer w.Close()

	_, err = w.Write([]byte(content))
	if err != nil {
		return err
	}
	return nil
}

func CompressFile(sourcePath, targetPath string) error {
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()
	target, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer target.Close()

	w := zlib.NewWriter(target)
	defer w.Close()

	_, err = io.Copy(w, source)
	if err != nil {
		return err
	}

	return nil
}
