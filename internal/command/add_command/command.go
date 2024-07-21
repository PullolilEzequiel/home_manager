package addcommand

import (
	configmanager "github.com/PullolilEzequiel/wizard-home/internal/config_manager"
	"github.com/spf13/cobra"
)

func Execute(cmd *cobra.Command, args []string) {

	c := configmanager.GetConfig()
	cobra.CheckErr(c.AddConfigPathAndSave(args[0]))

}
