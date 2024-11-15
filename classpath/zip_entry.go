package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func (z *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(z.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()
	for _, file := range reader.File {
		if file.Name == className {
			rc, err := file.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := io.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, z, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (z *ZipEntry) String() string {
	return z.absPath
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}
