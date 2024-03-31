package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Get the version",
		Long:  "Get the application version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("v1.0.1")
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
