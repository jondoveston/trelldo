package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(archiveCmd)
}

var archiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive things",
	Long:  "Archive things",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("archive")
	},
}
