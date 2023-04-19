package cmd

import (
	"fmt"
	"github.com/adlio/trello"
	"github.com/spf13/cobra"
)

func init() {
	actionCmd.AddCommand(actionCardCmd)
}

var actionCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Action card",
	Long:  "Action card",
	Run: func(cmd *cobra.Command, args []string) {
		card, err := getCard()
		if err != nil {
			er(err)
		}
		actions, err := card.GetActions(trello.Defaults())
		if err != nil {
			er(err)
		}
		for _, action := range actions {
			fmt.Println(action.Type)
		}
	},
}
