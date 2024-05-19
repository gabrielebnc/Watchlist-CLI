package cmd

import (
	"net/url"
	"os"
	"strconv"

	"github.com/gabrielebnc/Watchlist-CLI/persistency"
	"github.com/gabrielebnc/Watchlist-CLI/utils"
	"github.com/gabrielebnc/Watchlist-CLI/youtube"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	infoCmd = &cobra.Command{
		Use:   "info {'all'|index}",
		Short: "Info Command for Watchlist CLI",
		Long:  "Get infos for one or all listed items, possible args: 'all', index",

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

			if arg == "all" { //TODO implement
				utils.PrintfSTDOUT("Infos for all elements in the list")
				if linesCount == 0 {
					utils.PrintfSTDOUT("No items in the watchlist.\nAdd one using 'watch add <url>'")
				}
				for i, item := range allLines {
					utils.PrintfSTDOUT("%v.\n %v", i+1, item)
				}
				return
			}

			index, err := strconv.Atoi(arg)
			if isIndexExistent := index <= len(allLines) && index > 0; err == nil && isIndexExistent { //TODO refactor
				u, err := url.Parse(allLines[index-1])

				if err != nil {
					utils.PrintfSTDERR("Error parsing URL: %v", allLines[index-1])
					os.Exit(1)
				}

				q := u.Query()
				vQueryParams := q["v"]
				if len(vQueryParams) < 1 {
					utils.PrintfSTDERR("Couldn't find a videoId in the given URL")
					os.Exit(1)
				}
				vInfo := youtube.SearchVideoById(vQueryParams[0], youtubeApiKey)

				utils.PrintfSTDOUT("Getting info for element #%v in the watchlist\n%v.\n%v", index, index, vInfo.Fstring())
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
	youtubeApiKey = viper.GetString("watchcli.configs.youtubeApiKey")
}
