package cmd

import (
	"dib/internal/cmd/migrate"
	"dib/internal/cmd/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dib",
	Short: "dib app CLI",
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the application",
	Run:   server.Serve,
}

var migrateUpCmd = &cobra.Command{
	Use:   "migrate:up",
	Short: "Migrates Up the database",
	Run:   migrate.Up,
}

var migrateDownCmd = &cobra.Command{
	Use:   "migrate:down",
	Short: "Migrates Down the database",
	Run:   migrate.Down,
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
