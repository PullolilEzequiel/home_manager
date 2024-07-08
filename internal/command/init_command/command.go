package initcommand

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

/*
Generate the file config.json in the folder wizard_home if
not exist, in the case config.json exist in the folder raise an exception
*/
func Execute(cmd *cobra.Command, args []string) {
	var user, _ = os.UserHomeDir()

	configFolder := fmt.Sprintf("%s/.config/wizard_home", user)
	fmt.Printf("Init Wizard-home config in the directory: %s \n", configFolder)
	if err := createConfig(configFolder); err != nil {
		cobra.CheckErr(err)
	}

}
