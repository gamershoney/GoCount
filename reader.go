package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindSheet(path string) (*Input, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("error: finding file: %w", err)
	}
	return &Input{path: path, filename: filepath.Base(path)}, nil
}
