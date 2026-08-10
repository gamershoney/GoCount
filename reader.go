package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

func FindSheet(path string) (*Input, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("error: finding file: %w", err)
	}
	return &Input{path: path, filename: filepath.Base(path)}, nil
}

type WorkBook struct {
	SheetList   []string
	ActiveSheet string
	Error       error
}

func GetSheets(path string) *WorkBook {
	wb := &WorkBook{}

	file, err := excelize.OpenFile(path)
	if err != nil {
		wb.Error = fmt.Errorf("error opening file: %w", err)
		return wb
	}
	defer file.Close()
	wb.SheetList = file.GetSheetList()
	return wb
}
