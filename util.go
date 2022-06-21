package main

import (
	"errors"
	"fmt"
	"os"
)

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		fmt.Errorf("encountered exception checking if path exists %s: %s", path, err)
		return false
	}
}

func CreateFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Errorf("encountered exception creating file: %s", err)
	}
	err = file.Close()
	if err != nil {
		fmt.Errorf("encountered excpetion closing file")
	}
}
