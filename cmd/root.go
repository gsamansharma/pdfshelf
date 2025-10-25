package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdfshelf",
	Short: "A CLI tool to manage your PDF library",
	Long: `pdfshelf helps you add, list, and open PDFs from your terminal,
keeping track of your reading time and bookmarks.`,

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

