package cmd

import (
	"fmt"
	"os"

	"github.com/gabrielebnc/Watchlist-CLI/persistency"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Get Command for Watchlist CLI",
		Long:  "Get commands gets all the listed items",

		Run: func(cmd *cobra.Command, args []string) {

			allLines, err := persistency.ReadAllLines(persistence_path)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			fmt.Fprintln(os.Stdout, "Listing all items in the watchlist")
			for i, item := range allLines {
				output := fmt.Sprintf("%v. %v", i+1, item)
				fmt.Fprintln(os.Stdout, output)
			}
		},
	}
)

func init() {

	cobra.OnInitialize(InitFunc)

	persistence_path = viper.GetString("watchcli.configs.persistencePath")
}
