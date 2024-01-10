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
	Short: "Print the version number of test-cobra",
	Long:  `All software has versions. This is test-cobra's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test-cobra v0.0 -- HEAD")
	},
}
