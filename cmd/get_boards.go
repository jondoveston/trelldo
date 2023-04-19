package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getBoardsCmd)
}

var getBoardsCmd = &cobra.Command{
	Use:   "boards",
	Short: "Get boards",
	Long:  "Get boards",
	Run: func(cmd *cobra.Command, args []string) {
		boards, err := getBoards()
		if err != nil {
			er(err)
		}
		for _, board := range boards {
			fmt.Printf("%v: %v\n", board.ID, board.Name)
		}
	},
}
