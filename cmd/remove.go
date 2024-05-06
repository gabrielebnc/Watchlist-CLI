package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gabrielebnc/Watchlist-CLI/persistency"
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
				fmt.Fprintln(os.Stderr, "Index must be an integer")
				return
			}

			err = persistency.RemoveLineAtIndex(persistence_path, index)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				return
			}
			fmt.Fprintln(os.Stdout, "Deleted successfully")
		},
	}
)

func init() {

	persistence_path = viper.GetString("watchcli.configs.persistencePath")
}
