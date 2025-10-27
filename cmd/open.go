package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pdfshelf/internal/commands"
)

var openCmd = &cobra.Command{
	Use:   "open <id>",
	Short: "Open a PDF from your library using zathura",
	Long:  `Opens the PDF with the specified ID in zathura, starting at the last bookmarked page.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entry, index, err := commands.FindByID(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := commands.LaunchZathura(*entry, index); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}