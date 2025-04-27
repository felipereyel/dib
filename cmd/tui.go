package cmd

import (
	"dib/internal/tui"
	"log"

	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Run DIB in tui format for CSV file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Bad Args")
		}

		if err := tui.Tui(args[0]); err != nil {
			log.Fatal("Tui Error", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
