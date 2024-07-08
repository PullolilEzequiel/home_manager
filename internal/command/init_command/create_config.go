package initcommand

import (
	"fmt"
	"io/fs"
	"os"
)

func createConfig(path string) error {
	if err := os.Mkdir(path, fs.FileMode(0777)); !os.IsExist(err) {
		return createConfigFile(path)
	}

	return fmt.Errorf("config already exist in the folder %s already exist", path)
}

func createConfigFile(path string) error {

	configInitData := "{\n \"repository_url\": \"\", \n \"configs_to_persist\": []\n}"

	if err := os.WriteFile(fmt.Sprintf("%s/config.json", path), []byte(configInitData), 0644); err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}
