package addcommand

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Execute(cmd *cobra.Command, args []string) {
	fmt.Println("Adding file or folder", args[0])

}
