package setupcommand

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func Execute(cmd *cobra.Command, args []string) {
	fmt.Println("Setup config from: ", args[0])
	t := time.Now()

	defer func() { fmt.Printf("Setup ended in %.2f \n", time.Since(t).Seconds()) }()
	sm := SetupManager(args[0])
	cobra.CheckErr(sm.SetupConfigState())
}
