package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	count int
	name  string

	showCmd = &cobra.Command{
		Use:   "show",
		Short: "Show information about robots",
		Long:  "Show information about robots",
		// Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("robot's name:", name)
			fmt.Println("print's count:", count)
			for _, arg := range args {
				fmt.Println("show: " + arg)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().IntVarP(&count, "count", "c", 0, "count to print")
	showCmd.Flags().StringVarP(&name, "name", "n", "", "the robot's name")
}
