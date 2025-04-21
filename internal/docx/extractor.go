package docx

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Extractor struct{
	InputPath string
	OutputPath string
}

func NewExtractor(inputPath, outputPath string) *Extractor {
	return &Extractor{
		InputPath: inputPath,
		OutputPath: outputPath,
	}
}

// Extract unzip DOCX file to an out directory

func (e *Extractor) Extract() error{
	// Open DOCX file

	reader, err := zip.OpenReader(e.InputPath)
	if err != nil {
		return fmt.Errorf("Error opening DOCX file: %w", err)
	}
	defer reader.Close()

	// Create extraction directory if it doesn't exist
	if err := os.MkdirAll(e.OutputPath, 0755);
	err != nil {
		return fmt.Errorf("Error creatin extraction directory %w", err)
	}

	// Extrac each file
	for _, file := range reader.File {
		// Get the file path
		filePath := filepath.Join(e.OutputPath,file.Name)

		// Check if File is a directory then create the directory
		if file.FileInfo().IsDir(){
			os.MkdirAll(filePath, 0755)
		}

		if err := os.MkdirAll(filepath.Dir(filePath), 0755);
		err != nil {
			return fmt.Errorf("Error creating the directory: %w", err)
		}

		// Create file
		outFile, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("Error creating file: %w", err)
		}

		inFile, err := file.Open()
		if err !=nil {
			outFile.Close()
			return fmt.Errorf("Error opening file in archive: %w", err)
		}

		if _, err := io.Copy(outFile, inFile);
		err !=nil{
			outFile.Close()
			inFile.Close()
			return fmt.Errorf("Error copying file: %w", err)
		}

		// Close Files
		outFile.Close()
		inFile.Close()
	}

	return nil
}