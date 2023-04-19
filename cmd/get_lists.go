package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getListsCmd)
}

var getListsCmd = &cobra.Command{
	Use:   "lists",
	Short: "Get lists",
	Long:  "Get lists",
	Run: func(cmd *cobra.Command, args []string) {
		lists, err := getLists()
		if err != nil {
			er(err)
		}
		for _, list := range lists {
			fmt.Printf("%v: %v\n", list.ID, list.Name)
		}
	},
}
