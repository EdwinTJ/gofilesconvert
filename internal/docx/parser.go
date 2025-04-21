package docx

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Document struct {
	XMLName xml.Name `xml:"document"`
	Body Body `xml:"body"`
}

type Body struct {
	XMLName xml.Name `xml:"body"`
	Paragraphs []Paragraph `xml:"p"`
}

type Paragraph struct {
	XMLName xml.Name `xml:"p"`
	Runs []Run `xml:"r"`
	Style Style `xml:"pPr>pStyle"`
}

type Run struct {
	XMLName xml.Name `xml:"r"`
	Text    []Text   `xml:"t"`
	Bold    *Bold    `xml:"rPr>b"`
	Italic  *Italic  `xml:"rPr>i"`
}

type Text struct {
	XMLName xml.Name `xml:"t"`
	Content string `xml:",chardata"`
}

type Style struct {
	Val string `xml:"val,attr"`
}

type Bold struct{}
type Italic struct{}

type Parser struct {
	ExtractPath string
}

func NewParser(extractPath string) *Parser {
	return &Parser{
		ExtractPath: extractPath,
	}
} 

func (p *Parser) ParseDocument() (string, error){
	documentPath := filepath.Join(p.ExtractPath, "word", "document.xml")

	// Check if file exits
	if _, err := os.Stat(documentPath);
	os.IsNotExist(err){
		return "", fmt.Errorf("document.xml not found at %s", documentPath)
	}

	xmlData, err := os.ReadFile(documentPath)
	if err != nil {
		return "", fmt.Errorf("Error reading document.xml: %w", err)
	}

	// Parse XML
	var doc Document
	err = xml.Unmarshal(xmlData, &doc)
	if err != nil{
		return "", fmt.Errorf("Error parsing XML: %w", err)
	}

	// Extract text with styling

	var result strings.Builder

	for _, paragraph := range doc.Body.Paragraphs {
		style := "normal"
		if paragraph.Style.Val != "" {
			style = paragraph.Style.Val
		}

		for _,run := range paragraph.Runs{
			for _, text := range run.Text{
				// Apply styling information
				formatting := ""
				if run.Bold != nil {
					formatting += "[Bold] "
				}
				if run.Italic != nil {
					formatting += "[Italic] "
				}
				if formatting != "" {
					result.WriteString(fmt.Sprintf("%s%s",formatting, text.Content))
				}else {
					result.WriteString(text.Content)
				}
			}
		}

		// Add paragraph style info
		result.WriteString(fmt.Sprintf(" [Style: %s]\n", style))
	}

	return result.String(), nil
}