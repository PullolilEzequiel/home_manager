package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	initCommand "github.com/PullolilEzequiel/wizard-home/src/init_command"
	reverseCommand "github.com/PullolilEzequiel/wizard-home/src/reverse_command"
	saveCommand "github.com/PullolilEzequiel/wizard-home/src/save_command"
	setupCommand "github.com/PullolilEzequiel/wizard-home/src/setup_command"
)

var initC = &cobra.Command{
	Use: "init", Short: "Create initial config for wizard_home",
	Run: initCommand.Execute,
}
var saveC = &cobra.Command{
	Use: "save", Short: "Persist our actual config to our remote repository",
	Run: saveCommand.Execute}
var reverseC = &cobra.Command{
	Use: "reverse", Short: "Reverse the actual system config to the last commit in our remote repository",
	Run: reverseCommand.Execute}
var setupC = &cobra.Command{
	Use: "setup", Short: "Change the system configuration to another Wizard Home repository",
	Args: cobra.ExactArgs(1),
	Run:  setupCommand.Execute,
}

func main() {
	rootCmd := &cobra.Command{Use: "wizard_home",
		Short: "Wizard home is a CLI tool that seeks to help preserve the personalized configuration of our system.",
	}

	rootCmd.AddCommand(initC)
	rootCmd.AddCommand(saveC)
	rootCmd.AddCommand(reverseC)
	rootCmd.AddCommand(setupC)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
