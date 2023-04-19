package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(actionCmd)
}

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Action things",
	Long:  "Action things",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("action")
	},
}
