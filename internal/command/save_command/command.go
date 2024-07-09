package savecommand

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute(cmd *cobra.Command, args []string) {
	var user, _ = os.UserHomeDir()
	fmt.Println("Saving config files...")
	configFolder := fmt.Sprintf("%s/.config/wizard_home", user)
	dir, _ := os.MkdirTemp(configFolder, "saving_state")

	PersistFiles(dir, configFolder)

	//defer os.RemoveAll(dir)
}
