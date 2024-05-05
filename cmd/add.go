package cmd

import (
	"fmt"
	"os"

	"github.com/gabrielebnc/Watchlist-CLI/persistency"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	addCmd = &cobra.Command{
		Use:   "add string",
		Short: "Add Command for Watchlist CLI",
		Long:  "Adds an item to the Watchlist",

		Args: cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {

			err := persistency.PersistItem(args[0], persistence_path)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				return
			}
			fmt.Fprintln(os.Stdout, "Successfully added item.")
		},
	}
)

func init() {

	cobra.OnInitialize(InitFunc)

	persistence_path = viper.GetString("watchcli.configs.persistencePath")
}
