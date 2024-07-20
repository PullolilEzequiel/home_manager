package savecommand

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type JsonConfig struct {
	Repository_url     string   `json:repository_url`
	Configs_to_persist []string `json:configs_to_persist`
}

func Execute(cmd *cobra.Command, args []string) {
	fmt.Println("Saving config files...")
	t := time.Now()
	defer func() { fmt.Printf("Save ended in %.2f \n", time.Since(t).Minutes()) }()

	var user, _ = os.UserHomeDir()
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
