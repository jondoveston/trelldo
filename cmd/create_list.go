package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

func init() {
	createCmd.AddCommand(createListCmd)
}

var createListCmd = &cobra.Command{
	Use:   "list [name]",
	Short: "Create list",
	Long:  "Create list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		list, err := createList(strings.Join(args, " "))
		if err != nil {
			er(err)
		}
		viper.Set("list", list.Name)
		fmt.Printf("Created list %v\n", list.Name)
	},
}
