package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// TODO read fils based on the path give and return the file that is read
// handle error greadefully
func isFilesExist(pathtoRead string) bool {
	info, err := os.Stat(pathtoRead)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Errorf("files does not exist")
		return false
	}

	return info.IsDir()
}

func ReadFile(Filepath string) string {
	homePath, err := os.UserHomeDir()
	check(err)
	fullFilePath := filepath.Join(homePath, ".", Filepath)
	isDir := isFilesExist(fullFilePath)
	if isDir {
		fmt.Errorf("not files path")
		panic(err)
	}

	file, err := os.ReadFile(fullFilePath)
	check(err)

	// fmt.Println(string(file))
	return string(file)
}
