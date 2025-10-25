package cmd

import (
	"github.com/spf13/cobra"
	"pdfshelf/internal/commands"
)

var openCmd = &cobra.Command{
	Use:   "open <id>",
	Short: "Open a PDF from your library using zathura",
	Long:  `Opens the PDF with the specified ID in zathura, starting at the last bookmarked page.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commands.Open(args[0])
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}