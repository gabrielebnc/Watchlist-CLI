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
	getCmd = &cobra.Command{
		Use:   "get {'all'|index}",
		Short: "Get Command for Watchlist CLI",
		Long:  "Get one or all listed items, possible args: 'all', index",

		Args: cobra.ExactArgs(1),
		//TODO add actual args validation

		Run: func(cmd *cobra.Command, args []string) {

			arg := args[0]

			allLines, err := persistency.ReadAllLines(persistence_path)
			linesCount := len(allLines)

			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				return
			}

			if arg == "all" {
				fmt.Fprintln(os.Stdout, "Listing all items in the watchlist")
				if linesCount == 0 {
					fmt.Fprintln(os.Stdout, "No items in the watchlist.\nAdd one using 'watch add <url>'")
				}
				for i, item := range allLines {
					output := fmt.Sprintf("%v. %v", i+1, item)
					fmt.Fprintln(os.Stdout, output)
				}
				return
			}

			index, err := strconv.Atoi(arg)
			if isIndexExistent := index <= len(allLines) && index > 0; err == nil && isIndexExistent {
				output := fmt.Sprintf("Getting element #%v in the watchlist\n%v. %v", index, index, allLines[index-1])
				fmt.Fprintln(os.Stdout, output)
				return
			} else if err != nil {
				output := fmt.Sprintf("Invalid argument: %v", arg)
				fmt.Fprintln(os.Stderr, output)
				return
			} else if !isIndexExistent {
				output := fmt.Sprintf("Index does not exist: %v", index)
				fmt.Fprintln(os.Stderr, output)
				return
			}

		},
	}
)

func init() {

	persistence_path = viper.GetString("watchcli.configs.persistencePath")
}
