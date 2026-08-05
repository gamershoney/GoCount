package main

import (
	"errors"
	"io/fs"
	"path/filepath"
)

func FindSheet(path string) (*Input, error) {
	if !fs.ValidPath(path) {
		return nil, errors.New("error: file not found")
	}
	return &Input{path: path, filename: filepath.Base(path)}, nil
}
