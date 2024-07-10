package configmanager

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type __jsonConfig struct {
	Repository_url     string   `json:repository_url`
	Configs_to_persist []string `json:configs_to_persist`
}

type Config struct {
	wizard_home_path string
	repository_url   string
	config_paths     []string
}

func GetConfig() Config {
	obj := __jsonConfig{}
	user, err := os.UserHomeDir()
	panicCheck(err)
	path := fmt.Sprintf("%s/.config/wizard_home", user)

	data := readUserConfig(path)

	json.Unmarshal(data, &obj)
	return Config{
		wizard_home_path: path,
		repository_url:   obj.Repository_url,
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

func (c Config) Path() string          { return c.wizard_home_path }
func (c Config) RepoUrl() string       { return c.repository_url }
func (c Config) ConfigPaths() []string { return c.config_paths }

/*
Creates a temporal folder in the wizard_home folder with the patter passed for paramter and make action inside it and remove the folder
@folderPattern string
  - The pattern name

@insideTemporalFolderAction func(folderName string) error
  - Action performed within the temporary folder.
*/
func (c Config) CreateTemporalFolder(folderPattern string, insideTemporalFolderAction func(folderName string) error) error {
	s, err := os.MkdirTemp(c.Path(), folderPattern)
	if err != nil {
		return err
	}
	defer os.RemoveAll(s)

	return insideTemporalFolderAction(s)
}
