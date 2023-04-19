package cmd

import (
	"fmt"
	"github.com/adlio/trello"
	"github.com/spf13/cobra"
)

func init() {
	actionCmd.AddCommand(actionListCmd)
}

var actionListCmd = &cobra.Command{
	Use:   "list",
	Short: "Action list",
	Long:  "Action list",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := getList()
		if err != nil {
			er(err)
		}
		actions, err := list.GetActions(trello.Defaults())
		if err != nil {
			er(err)
		}
		for _, action := range actions {
			fmt.Println(action.Type)
		}
	},
}
