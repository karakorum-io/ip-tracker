package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "cobra",
		Short: "IP tracker",
		Long:  "IP tracker Long",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
