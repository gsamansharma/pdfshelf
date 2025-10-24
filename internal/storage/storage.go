package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"pdfshelf/internal/model" 
)

func getLibraryPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	appConfigDir := filepath.Join(configDir, "pdfshelf")
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(appConfigDir, "library.json"), nil
}

func LoadLibrary() (*model.Library, error) {
	path, err := getLibraryPath()
	if err != nil {
		return nil, err
	}

	lib := &model.Library{NextID: 1} 

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return lib, nil 
		}
		return nil, err 
	}

	if err = json.Unmarshal(data, lib); err != nil {
		return nil, err
	}
	return lib, nil
}

func SaveLibrary(lib *model.Library) error {
	path, err := getLibraryPath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(lib, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}