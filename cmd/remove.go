package cmd

import (
	"strconv"

	"github.com/gabrielebnc/Watchlist-CLI/persistency"
	"github.com/gabrielebnc/Watchlist-CLI/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	removeCmd = &cobra.Command{
		Use:   "remove index",
		Short: "Remove Command for Watchlist CLI",
		Long:  "Removes item at specified index from list",

		Aliases: []string{"delete"},

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			index, err := strconv.Atoi(args[0])
			if err != nil {
				utils.PrintfSTDERR("Index must be an integer")
				return
			}

			err = persistency.RemoveLineAtIndex(persistencePath, index)
			if err != nil {
				utils.PrintfSTDERR("Error removing line: %v", err)
				return
			}
			utils.PrintfSTDOUT("Deleted successfully")
		},
	}
)

func init() {

	persistencePath = viper.GetString("watchcli.configs.persistencePath")
}
