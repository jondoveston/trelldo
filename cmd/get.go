package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get things",
	Long:  "Get things",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get")
	},
}
