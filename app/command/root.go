package command

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "conflr",
		Short: "interact between solr and confluence",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func init() {
	rootCmd.AddCommand(importCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
