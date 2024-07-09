package savecommand

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type JsonConfig struct {
	Repository_url     string   `json:repository_url`
	Configs_to_persist []string `json:configs_to_persist`
}

func PersistFiles(folderPath, configPath string) {
	config := getConfigFields(configPath)
	fmt.Println(config.Repository_url)
	copyFiles(folderPath, config.Configs_to_persist)
}

/*
Copy all the files in config path to wizard_home folder
*/
func copyFiles(tmpFolder string, paths []string) {

	for _, path := range paths {
		copyPathIn(path, tmpFolder)
	}
}

func copyPathIn(path, pathReceptor string) {
	fmt.Printf("Copy file %s in directory %s \n", path, pathReceptor)
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
	fmt.Println("--------------------")
	fmt.Println(filePath)
	fmt.Println(path)
	fmt.Println(createFilePath(filePath, path))
	fmt.Println("--------------------")
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

func getConfigFields(configPath string) JsonConfig {
	obj := JsonConfig{}
	content, _ := os.Open(fmt.Sprintf("%s/config.json", configPath))

	data, _ := io.ReadAll(content)
	json.Unmarshal(data, &obj)
	return obj
}
