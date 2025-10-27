package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"pdfshelf/internal/commands"
	"pdfshelf/internal/storage"
	"pdfshelf/internal/tui" 
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all PDFs in your library (interactive)",
	Long:  `Launches an interactive TUI to browse, search, and open PDFs.`,
	Run: func(cmd *cobra.Command, args []string) {
		lib, err := storage.LoadLibrary()
		if err != nil {
			fmt.Println("Error loading library:", err)
			os.Exit(1)
		}

		if len(lib.PDFs) == 0 {
			fmt.Println("Your shelf is empty. Add a PDF with 'pdfshelf add ...'")
			return
		}

		selectedEntry, selectedIndex, err := tui.StartTUI(lib.PDFs)
		if err != nil {
			fmt.Println("Error running TUI:", err)
			os.Exit(1)
		}

		if selectedEntry != nil {
			if err := commands.LaunchZathura(*selectedEntry, selectedIndex); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}