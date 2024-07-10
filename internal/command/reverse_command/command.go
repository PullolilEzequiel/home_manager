package reversecommand

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func Execute(cmd *cobra.Command, args []string) {
	fmt.Println("Reverting system settings")
	t := time.Now()

	defer func() { fmt.Printf("Reverting ended in %f \n", time.Since(t).Seconds()) }()
	rvm := ReverseManager()

	cobra.CheckErr(rvm.ReverseConfigState())
}
