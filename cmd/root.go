package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "echosium",
	Short: "Your audio companion for coding",
	Long: `Echosium is an open-source CLI tool designed for developers seeking the perfect audio experience while coding. 
With seamless integration of the Jamendo API, Echobard offers a dynamic, mood-based music selection to enhance your coding flow.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
