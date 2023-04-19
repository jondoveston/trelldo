package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getCardsCmd)
}

var getCardsCmd = &cobra.Command{
	Use:   "cards",
	Short: "Get cards",
	Long:  "Get cards",
	Run: func(cmd *cobra.Command, args []string) {
		cards, err := getCards()
		if err != nil {
			er(err)
		}
		for _, card := range cards {
			fmt.Printf("%v: %v: %v\n", card.ID, card.Name, card.Desc)
		}
	},
}
