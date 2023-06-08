package cmd

import (
	"fmt"
	"github.com/adlio/trello"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	archiveCardCmd.Flags().BoolVarP(&Head, "head", "", false, "Use card at the head (top) of list")
	archiveCardCmd.Flags().BoolVarP(&Tail, "tail", "", false, "Use card at the tail (bottom) of list")
	archiveCmd.AddCommand(archiveCardCmd)
}

var archiveCardCmd = &cobra.Command{
	Use:   "card",
	Short: "Archive card",
	Long:  "Archive card",
	Run: func(cmd *cobra.Command, args []string) {
		var card *trello.Card
		var err error
		if Head {
			card, err = getHeadCard()
		} else if Tail {
			card, err = getTailCard()
		} else {
			card, err = getCard()
		}
		if err != nil {
			er(err)
		}
		viper.Set("card", card.ID)
		err = card.Archive()
		if err != nil {
			er(err)
		}
		fmt.Printf("Archived card %v: %v: %v\n", card.ID, card.Name, card.Desc)
	},
}
