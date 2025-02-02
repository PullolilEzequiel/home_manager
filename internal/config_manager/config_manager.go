package configmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	directorymanagement "github.com/PullolilEzequiel/wizard-home/internal/directory_management"
)

type __jsonConfig struct {
	Repository_url     string   `json:repository_url`
	Configs_to_persist []string `json:configs_to_persist`
}

type Config struct {
	user_home        string
	wizard_home_path string
	repository_url   string
	repository_name  string
	config_paths     []string
}

func GetConfig() Config {
	obj := __jsonConfig{}
	user, err := os.UserHomeDir()
	panicCheck(err)
	p := fmt.Sprintf("%s/.config/wizard_home", user)

	data := readUserConfig(p)

	json.Unmarshal(data, &obj)
	return Config{
		user_home:        user,
		wizard_home_path: p,
		repository_url:   obj.Repository_url,
		repository_name:  path.Base(obj.Repository_url),
		config_paths:     obj.Configs_to_persist}
}

func readUserConfig(path string) []byte {
	content, err := os.Open(path + "/config.json")
	panicCheck(err)
	data, err := io.ReadAll(content)
	panicCheck(err)
	return data

}

func panicCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func (c Config) Path() string           { return c.wizard_home_path }
func (c Config) ConfigPaths() []string  { return c.config_paths }
func (c Config) RepoUrl() string        { return c.repository_url }
func (c Config) RepoName() string       { return c.repository_name }
func (c Config) HomeDir() string        { return c.user_home }
func (c Config) ConfigFilePath() string { return path.Join(c.wizard_home_path, "config.json") }

/*
Creates a temporal folder in the wizard_home folder with the patter passed for paramter and make action inside it and remove the folder
@folderPattern string :The pattern name

@insideTemporalFolderAction func(folderName string) error : Action performed within the temporary folder.
*/
func (c Config) CreateTemporalFolder(folderPattern string, insideTemporalFolderAction func(folderName string) error) error {
	s, err := os.MkdirTemp(c.Path(), folderPattern)
	if err != nil {
		return err
	}
	defer os.RemoveAll(s)

	return insideTemporalFolderAction(s)
}

func (c *Config) AddConfigPathAndSave(fileOrDirectoryPath string) error {
	fmt.Printf("Added \"%s\" to the config.json \n", fileOrDirectoryPath)

	if err := c.addConfig(fileOrDirectoryPath); err != nil {
		return err
	}

	return c.saveConfig()
}

func (c *Config) addConfig(fileOrDirectoryPath string) error {
	pathToAdd, isValid := directorymanagement.TransformPath(fileOrDirectoryPath)

	if !isValid {
		return errors.New("The path " + fileOrDirectoryPath + "is invalid")
	}

	c.config_paths = append(c.config_paths, pathToAdd)

	return nil
}

func (c Config) saveConfig() error {

	publicConfig := __jsonConfig{
		Repository_url:     c.repository_url,
		Configs_to_persist: c.config_paths,
	}
	if content, err := json.MarshalIndent(publicConfig, "", "\t"); err != nil {
		return err
	} else {
		return os.WriteFile(c.ConfigFilePath(), content, 0644)
	}

}
