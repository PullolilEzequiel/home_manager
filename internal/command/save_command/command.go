package savecommand

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

type JsonConfig struct {
	Repository_url     string   `json:repository_url`
	Configs_to_persist []string `json:configs_to_persist`
}

func Execute(cmd *cobra.Command, args []string) {
	var user, _ = os.UserHomeDir()
	fmt.Println("Saving config files...")
	configFolder := fmt.Sprintf("%s/.config/wizard_home", user)
	dir, _ := os.MkdirTemp(configFolder, "saving_state")
	config := getConfigFields(configFolder)
	PersistFiles(dir, configFolder, config.Configs_to_persist)
	PushChanges(dir, config.Repository_url)

	defer os.RemoveAll(dir)
}
func getConfigFields(configPath string) JsonConfig {
	obj := JsonConfig{}
	content, _ := os.Open(fmt.Sprintf("%s/config.json", configPath))

	data, _ := io.ReadAll(content)
	json.Unmarshal(data, &obj)
	return obj
}
