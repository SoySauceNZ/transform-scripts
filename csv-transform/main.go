package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "csv-transform [input-crashes] [input-coords] [output]",
	Short: "csv-transform input1.csv input2.csv output.csv",
	Long: `Migrate LatLng .csv and map .csv into count and severity. 
		Count is the number of crashes in the LatLng viewport using two 
		coordinates. Severity is the average of crash severity.`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if err := read(args); err != nil {
			panic(err)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
