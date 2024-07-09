package savecommand

import (
	"fmt"
	"os"
	"strings"
)

func PersistFiles(folderPath, configPath string, paths []string) {
	for _, path := range paths {
		fmt.Printf("- Saving state of %s \n", path)
		copyPathIn(path, folderPath)
	}
}

func copyPathIn(path, pathReceptor string) {
	mode, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	if mode.IsDir() {
		copyDirectory(path, pathReceptor)
	} else {
		copyFile(path, pathReceptor)
	}
}

func copyDirectory(directoryPath, path string) {
	dir := createFolder(path, nameOf(directoryPath))

	files, _ := os.ReadDir(directoryPath)
	for _, f := range files {
		pathOf := fmt.Sprintf("%s/%s", directoryPath, f.Name())
		copyPathIn(pathOf, dir)
	}
}

func createFolder(path, name string) string {
	dir := fmt.Sprintf("%s/%s", path, name)
	if err := os.Mkdir(dir, 0777); err != nil {
		panic(err)
	}

	return dir
}
func copyFile(filePath, path string) {

	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	os.WriteFile(createFilePath(filePath, path), content, 0666)
}

func createFilePath(filePath, path string) string {
	name := nameOf(filePath)

	return fmt.Sprintf("%s/%s", path, name)
}

func nameOf(filePath string) string {
	s := strings.Fields(strings.ReplaceAll(filePath, "/", " "))
	return s[len(s)-1]
}
