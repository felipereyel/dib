package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tui called")
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
