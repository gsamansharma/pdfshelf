package main

import (
	"fmt"
	"os"
	"pdfshelf/internal/commands" 
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pdfshelf <command> [options]")
		fmt.Println("Commands: add, list, open")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: pdfshelf add <path/to/file.pdf> [name]")
			return
		}
		filePath := os.Args[2]
		name := ""
		if len(os.Args) > 3 {
			name = os.Args[3]
		}
		commands.Add(filePath, name)

	case "list":
		commands.List()

	case "open":
		if len(os.Args) < 3 {
			fmt.Println("Usage: pdfshelf open <id>")
			return
		}
		commands.Open(os.Args[2])

	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}