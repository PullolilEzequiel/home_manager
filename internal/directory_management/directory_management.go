package directorymanagement

import (
	"fmt"
	"os"
)

func ReplaceFileOrFolder(pathToMove, path string) error {
	fmt.Println(pathToMove, path)
	mode, err := os.Stat(pathToMove)
	if err != nil {
		return err
	}
	if mode.IsDir() {
		return ReplaceFolder(pathToMove, path)
	} else {
		return ReplaceFile(pathToMove, path)
	}
}

func ReplaceFolder(folderPath, path string) error {
	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return os.Rename(folderPath, path)
}

func ReplaceFile(filePath, path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return os.Rename(filePath, path)
}
