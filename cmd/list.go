package cmd

import (
	"github.com/spf13/cobra"
	"pdfshelf/internal/commands" 
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all PDFs in your library",
	Long:  `Prints a formatted table of all PDFs in your library, including their ID, name, page, and time spent.`,
	Run: func(cmd *cobra.Command, args []string) {
		commands.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}