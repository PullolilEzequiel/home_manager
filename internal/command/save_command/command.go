package savecommand

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func Execute(cmd *cobra.Command, args []string) {
	fmt.Println("Saving config files...")
	t := time.Now()
	defer func() { fmt.Printf("Save ended in %.2f \n", time.Since(t).Seconds()) }()
	svm := SaveManager()
	cobra.CheckErr(svm.SaveConfigState())
}
