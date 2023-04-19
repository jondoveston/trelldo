package cmd

import (
	"fmt"
	"github.com/adlio/trello"
	"github.com/spf13/cobra"
)

func init() {
	actionCmd.AddCommand(actionBoardCmd)
}

var actionBoardCmd = &cobra.Command{
	Use:   "board",
	Short: "Action board",
	Long:  "Action board",
	Run: func(cmd *cobra.Command, args []string) {
		board, err := getBoard()
		if err != nil {
			er(err)
		}
		actions, err := board.GetActions(trello.Defaults())
		if err != nil {
			er(err)
		}
		for _, action := range actions {
			fmt.Println(action.Type)
		}
	},
}
