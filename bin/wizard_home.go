package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

//var urlInfo :=

var initC = &cobra.Command{Use: "init", Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("init")
	}}
var saveC = &cobra.Command{Use: "save", Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("save")
	}}
var reverseC = &cobra.Command{Use: "reverse", Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("reverse")
	}}
var setupC = &cobra.Command{
	Use: "setup", Short: "",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("setup")
	}}

func init() {

}

func main() {
	rootCmd := &cobra.Command{Use: "wizard_home", Short: ""}

	rootCmd.AddCommand(initC)
	rootCmd.AddCommand(saveC)
	rootCmd.AddCommand(reverseC)
	rootCmd.AddCommand(setupC)
	rootCmd.Execute()
}
