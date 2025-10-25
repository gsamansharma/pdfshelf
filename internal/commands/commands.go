package commands

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"pdfshelf/internal/model"
	"pdfshelf/internal/storage"
)

func Add(filePath string, name string) {
	lib, err := storage.LoadLibrary()
	if err != nil {
		fmt.Println("Error loading library:", err)
		return
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println("Error: File path is invalid:", err)
		return
	}

	if name == "" {
		name = filepath.Base(absPath)
	}

	entry := model.PDFEntry{
		ID:       lib.NextID,
		Name:     name,
		FilePath: absPath,
	}

	lib.PDFs = append(lib.PDFs, entry)
	lib.NextID++

	if err := storage.SaveLibrary(lib); err != nil {
		fmt.Println("Error saving library:", err)
		return
	}
	fmt.Printf("Added '%s' (ID: %d) to your shelf.\n", name, entry.ID)
}

func List() {
	lib, err := storage.LoadLibrary()
	if err != nil {
		fmt.Println("Error loading library:", err)
		return
	}
	if len(lib.PDFs) == 0 {
		fmt.Println("Your shelf is empty. Add a PDF with 'pdfshelf add ...'")
		return
	}
	fmt.Println("Your PDF Shelf:")
	for _, entry := range lib.PDFs {
		durationStr := entry.TotalTimeSpent.Round(time.Second).String()
		fmt.Printf("[%d] %s\n	%s\n    %s\n", entry.ID, entry.Name, durationStr, entry.FilePath)
	}
}

func Open(idStr string) {
	lib, err := storage.LoadLibrary()
	if err != nil {
		fmt.Println("Error loading library:", err)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Must be a number.")
		return
	}

	var foundEntry *model.PDFEntry
	var entryIndex int = -1 
	for i := range lib.PDFs {
		if lib.PDFs[i].ID == id {
			foundEntry = &lib.PDFs[i]
			entryIndex = i
			break
		}
	}

	if foundEntry == nil {
		fmt.Printf("Error: No PDF found with ID %d\n", id)
		return
	}

	if _, err := exec.LookPath("zathura"); err != nil {
		fmt.Println("Error: zathura executable not found in your PATH.")
		return
	}

	cmd := exec.Command("zathura", foundEntry.FilePath)

	fmt.Printf("Opening '%s'. Timer is running. (Terminal will wait until Zathura is closed)...\n", foundEntry.Name)

	startTime := time.Now()

	if err := cmd.Run(); err != nil {
		fmt.Println("Finished session with an error:", err)
	}

	duration := time.Since(startTime)
	
	fmt.Printf("Session finished. Time spent: %s\n", duration.Round(time.Second).String())


	lib.PDFs[entryIndex].TotalTimeSpent += duration

	if err := storage.SaveLibrary(lib); err != nil {
		fmt.Println("Error saving session time:", err)
		return
	}

	fmt.Println("Time saved.")
}