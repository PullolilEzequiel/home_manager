package configmanager

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type config struct {
	Repository_url     string   `json:repository_url`
	Configs_to_persist []string `json:configs_to_persist`
}

type Config struct {
	wizard_home_path string
	repository_url   string
	config_paths     []string
}

func GetConfig() Config {
	obj := config{}
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
