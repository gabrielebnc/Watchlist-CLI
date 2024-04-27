package cmd

import (
	"fmt"
	"os"

	"github.com/gabrielebnc/Watchlist-CLI/persistency"
	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:   "get filepath",
		Short: "Get Command for Watchlist CLI",
		Long:  "Get commands gets all the listed items",

		Run: func(cmd *cobra.Command, args []string) {
			output := fmt.Sprintf("Reading from %v", args[0])
			fmt.Fprintln(os.Stdout, output)
			allLines, err := persistency.ReadAllLines(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
			fmt.Fprintln(os.Stdout, "Listing all items in the watchlist")
			for i, item := range allLines {
				output := fmt.Sprintf("%v. %v", i+1, item)
				fmt.Fprintln(os.Stderr, output)
			}

		},
	}
)
