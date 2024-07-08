package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {

}

var initC = &cobra.Command{
	Use:   "init",
	Short: "Create initial config for wizard_home",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("init")
	}}
var saveC = &cobra.Command{
	Use:   "save",
	Short: "Persist our actual config to our remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("save")
	}}
var reverseC = &cobra.Command{
	Use:   "reverse",
	Short: "Reverse the actual system config to the last commit in our remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("reverse")
	}}
var setupC = &cobra.Command{
	Use:   "setup",
	Short: "Change the system configuration to another Wizard Home repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("setup")
	}}

func main() {
	rootCmd := &cobra.Command{
		Use:   "wizard_home",
		Short: "Wizard home is a CLI tool that seeks to help preserve the personalized configuration of our system.",
	}

	rootCmd.AddCommand(initC)
	rootCmd.AddCommand(saveC)
	rootCmd.AddCommand(reverseC)
	rootCmd.AddCommand(setupC)
	rootCmd.Execute()
}
