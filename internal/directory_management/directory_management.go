package directorymanagement

import (
	"fmt"
	"os"
)

/*
Replace the the file o folder passed for parameter to the location

@sourcePath string : The source of the folder o file
@locationToReplace string : The new location for the folder o file
*/
func ReplaceFileOrFolderFor(locationToReplace, sourcePath string) error {
	fmt.Println(sourcePath, locationToReplace)
	mode, err := os.Stat(sourcePath)
	if err != nil {
		return err
	}
	if mode.IsDir() {
		return ReplaceFolder(locationToReplace, sourcePath)
	} else {
		return ReplaceFile(locationToReplace, sourcePath)
	}
}

/*
Replace the oldDirectory for the new directory passed for parameter
@oldDirectory string: Path of the directory to replace
@sourceDirectory string: Source of the new directory
*/
func ReplaceFolder(oldDirectory, sourceDirectory string) error {
	if err := os.RemoveAll(oldDirectory); err != nil {
		return err
	}
	return os.Rename(sourceDirectory, oldDirectory)
}

/*
Replace the oldFile for the new file passed for parameter
@oldFile string: Path of the file to replace
@sourceFile string: Source of the new file
*/
func ReplaceFile(oldFile, sourceFile string) error {
	if err := os.Remove(oldFile); err != nil {
		return err
	}
	return os.Rename(sourceFile, oldFile)
}
