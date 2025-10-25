package cmd

import (
	"github.com/spf13/cobra"
	"pdfshelf/internal/commands"
)

var pdfName string

var addCmd = &cobra.Command{
	Use:   "add <filepath>", 
	Short: "Add a new PDF to your library",
	Long:  `Add a PDF to your library by providing its file path.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		commands.Add(filePath, pdfName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&pdfName, "name", "n", "", "Set a custom name for the PDF")
}