package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pdfshelf/internal/commands"
)

var removeCmd = &cobra.Command{
	Use:   "remove <id>",
	Short: "Remove a PDF from your library",
	Long:  `Removes a PDF from your library by its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := commands.Remove(args[0]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("PDF removed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}