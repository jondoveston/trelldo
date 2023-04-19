package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(unarchiveCmd)
}

var unarchiveCmd = &cobra.Command{
	Use:   "unarchive",
	Short: "Unarchive things",
	Long:  "Unarchive things",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unarchive")
	},
}
