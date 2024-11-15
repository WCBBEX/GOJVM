package classpath

import (
	"os"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func (d *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(d.absDir, className)
	data, err := os.ReadFile(fileName)
	return data, d, err
}

func (d *DirEntry) String() string {
	return d.absDir
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
