package cmd

import (
	"github.com/gabrielebnc/Watchlist-CLI/persistency"
	"github.com/gabrielebnc/Watchlist-CLI/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	addCmd = &cobra.Command{
		Use:   "add <item>",
		Short: "Add Command for Watchlist CLI",
		Long:  "Adds an item to the Watchlist",

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {

			err := persistency.PersistItem(args[0], persistencePath)
			if err != nil {
				utils.PrintfSTDERR("Error saving item: %v", err)
				return
			}
			utils.PrintfSTDOUT("Successfully added item.")
		},
	}
)

func init() {

	persistencePath = viper.GetString("watchcli.configs.persistencePath")
}
