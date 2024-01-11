package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of hello",
	Long:  `All software has versions. This is hello's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello v0.0 -- by boywei")
	},
}
