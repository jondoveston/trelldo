package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	createCardCmd.Flags().BoolVarP(&Head, "head", "", false, "Create card at the head (top) of list")
	createCardCmd.Flags().BoolVarP(&Tail, "tail", "", false, "Create card at the tail (bottom) of list")
	createCmd.AddCommand(createCardCmd)
}

var createCardCmd = &cobra.Command{
	Use:   "card [name] [description?]",
	Short: "Create card",
	Long:  "Create card",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			args = append(args, "")
		}
		card, err := createCard(args[0], args[1])
		if err != nil {
			er(err)
		}
    viper.Set("card", card.ID)
		fmt.Printf("Created card %v: %v: %v\n", card.ID, card.Name, card.Desc)
		if Head {
			err = card.MoveToTopOfList()
			if err != nil {
				er(err)
			}
			fmt.Println("Moved to head of list")
		} else if Tail {
			err = card.MoveToBottomOfList()
			if err != nil {
				er(err)
			}
			fmt.Println("Moved to tail of list")
		}
	},
}
