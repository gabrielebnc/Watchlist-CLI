package cmd

import (
	"strconv"

	"github.com/gabrielebnc/Watchlist-CLI/persistency"
	"github.com/gabrielebnc/Watchlist-CLI/utils"
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

			allLines, err := persistency.ReadAllLines(persistencePath)
			linesCount := len(allLines)

			if err != nil {
				utils.PrintfSTDERR("Error reading all lines: %v", err)
				return
			}

			if arg == "all" {
				utils.PrintfSTDOUT("Listing all items in the watchlist")
				if linesCount == 0 {
					utils.PrintfSTDOUT("No items in the watchlist.\nAdd one using 'watch add <url>'")
				}
				for i, item := range allLines {
					utils.PrintfSTDOUT("%v. %v", i+1, item)
				}
				return
			}

			index, err := strconv.Atoi(arg)
			if isIndexExistent := index <= len(allLines) && index > 0; err == nil && isIndexExistent {
				utils.PrintfSTDOUT("Getting element #%v in the watchlist\n%v. %v", index, index, allLines[index-1])
				return
			} else if err != nil {
				utils.PrintfSTDERR("Invalid argument: %v", arg)
				return
			} else if !isIndexExistent {
				utils.PrintfSTDERR("Index does not exist: %v", index)
				return
			}

		},
	}
)

func init() {

	persistencePath = viper.GetString("watchcli.configs.persistencePath")
}
