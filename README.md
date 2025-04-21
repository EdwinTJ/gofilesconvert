# DOCX to TXT Converter

A Go application for converting DOCX files to formatted text files.

## Features

- Extracts DOCX file contents
- Parses document.xml to extract text and basic styling
- Formats text with basic styling (bold, italic)
- Preserves paragraph styles
- Supports text flow, pagination, and line management

## Folder Structure

```
goconvert/
├── cmd/
│   └── docx2txt/
│       └── main.go         # Application entry point
├── internal/
│   ├── docx/
│   │   ├── extractor.go    # DOCX unzipping functionality
│   │   └── parser.go       # XML parsing functionality (your parsexml.go)
│   └── text/
│       ├── engine.go       # Text layout engine (your textengine.go)
│       ├── formatting.go   # Text formatting utilities
│       └── models.go       # Shared data structures
├── pkg/
│   └── fileutils/
│       └── fileutils.go    # General file utility functions
├── example_docs/           # Sample DOCX files for testing
├── go.mod
└── README.md
```

Build the project using the following command:

```bash
go mod init goconvert
go mod tidy
go build ./cmd/docx2txt
./docx2txt -input example_docs/CoverLetter.docx -output document.txt
```
