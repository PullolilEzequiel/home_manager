package directorymanagement

import (
	"io/fs"
	"os"
	"path"
)

/*
Replace the the file o folder passed for parameter to the location

@sourcePath string : The source of the folder o file
@locationToReplace string : The new location for the folder o file
*/
func ReplaceFileOrFolderFor(locationToReplace, sourcePath string) error {
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

/*
Copy the file or folder from the directory passed by parameter
@fileOrFolderPath -string- Path of  the file or folder to copy
@newPath -string- The path of the folder to copy file or folder
*/
func CopyFolderOrFile(fileOrFolderPath, newPath string) error {
	mode, err := os.Stat(fileOrFolderPath)
	if err != nil {
		return err
	}
	if mode.IsDir() {
		return copyFolder(fileOrFolderPath, newPath)
	} else {
		return copyFile(fileOrFolderPath, newPath)
	}
}

/*
Copy the folder from the directory passed by parameter
@folderPath -string- Path of  the folder to copy
@newPath -string- The path of the folder to copy the folder

! this not copy the .git folder, warning->
*/
func copyFolder(folderPath, newPath string) error {
	name := path.Base(folderPath)
	if name == ".git" {
		return nil
	}
	dir := createFolder(newPath, name)
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}
	for _, f := range files {
		pathOf := path.Join(folderPath, f.Name())
		if err := CopyFolderOrFile(pathOf, dir); err != nil {
			return err
		}
	}
	return nil
}

/*
Copy the folder from the directory passed by parameter
@filePath -string- Path of the file to copy
@newPath -string- The path of the folder to copy the file
*/
func copyFile(filePath, newPath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	pathToCopyFile := path.Join(newPath, path.Base(filePath))
	if err := os.WriteFile(pathToCopyFile, content, 0666); err != nil {
		return err
	}
	return nil
}
func createFolder(nPath, name string) string {
	dir := path.Join(nPath, name)
	if err := os.Mkdir(dir, 0777); err != nil {
		panic(err)
	}

	return dir
}

func TransformPath(pathToValidate string) (string, bool) {
	if !fs.ValidPath(pathToValidate) {
		return "", false
	}

	if pathToValidate == "." {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return dir, true

	}

	if !path.IsAbs(pathToValidate) {

		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return path.Join(dir, pathToValidate), true
	}

	return pathToValidate, true
}
