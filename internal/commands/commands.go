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
		ID:             lib.NextID,
		Name:           name,
		FilePath:       absPath,
		AddedOn:        time.Now(),
		TotalTimeSpent: 0,
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
		fmt.Printf("[%d] %s (Time: %s)\n    %s\n",
			entry.ID,
			entry.Name,
			durationStr,
			entry.FilePath)
	}
}

func FindByID(idStr string) (*model.PDFEntry, int, error) {
	lib, err := storage.LoadLibrary()
	if err != nil {
		return nil, -1, fmt.Errorf("Error loading library: %v", err)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, -1, fmt.Errorf("Invalid ID. Must be a number.")
	}

	var entryIndex int = -1
	for i := range lib.PDFs {
		if lib.PDFs[i].ID == id {
			entryIndex = i
			break
		}
	}

	if entryIndex == -1 {
		return nil, -1, fmt.Errorf("Error: No PDF found with ID %d", id)
	}

	return &lib.PDFs[entryIndex], entryIndex, nil
}

func LaunchZathura(entry model.PDFEntry, entryIndex int) error {
	if _, err := exec.LookPath("zathura"); err != nil {
		return fmt.Errorf("Error: zathura executable not found in your PATH.")
	}

	cmd := exec.Command("zathura", entry.FilePath)

	fmt.Printf("Opening '%s'. Timer is running...\n", entry.Name)
	startTime := time.Now()

	if err := cmd.Run(); err != nil {
		fmt.Println("Finished session with an error:", err)
	}

	duration := time.Since(startTime)
	fmt.Printf("Session finished. Time spent: %s\n", duration.Round(time.Second).String())

	lib, err := storage.LoadLibrary()
	if err != nil {
		return fmt.Errorf("Error loading library to save time: %v", err)
	}

	var found bool
	for i := range lib.PDFs {
		if lib.PDFs[i].ID == entry.ID {
			lib.PDFs[i].TotalTimeSpent += duration
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Error saving time: could not find entry after reload.")
	}

	if err := storage.SaveLibrary(lib); err != nil {
		return fmt.Errorf("Error saving session time: %v", err)
	}

	fmt.Println("Time saved.")
	return nil
}

func Remove(idStr string) error {
	lib, err := storage.LoadLibrary()
	if err != nil {
		return fmt.Errorf("Error loading library: %v", err)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("Invalid ID. Must be a number.")
	}

	foundIndex := -1
	for i, entry := range lib.PDFs {
		if entry.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return fmt.Errorf("Error: No PDF found with ID %d", id)
	}

	lib.PDFs = append(lib.PDFs[:foundIndex], lib.PDFs[foundIndex+1:]...)

	if err := storage.SaveLibrary(lib); err != nil {
		return fmt.Errorf("Error saving library after removal: %v", err)
	}
	return nil
}