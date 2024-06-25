package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tasnimzotder/go-junitxml/internal/models"
	"github.com/tasnimzotder/go-junitxml/internal/utils"
)

type Parser interface {
	Parse(data []byte) (*models.TestSuites, error)
	ParseFile(path string) (*models.TestSuites, error)
	ParseURL(url string) (*models.TestSuites, error)
}

type XMLParser struct{}

func (p *XMLParser) Parse(data []byte) (*models.TestSuites, error) {
	var testSuites models.TestSuites

	err := xml.Unmarshal(data, &testSuites)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal xml: %w", err)
	}

	_, err = utils.CalculateTotalsRootSuite(&testSuites)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate totals: %w", err)
	}

	return &testSuites, nil
}

func (p *XMLParser) ParseFile(path string) (*models.TestSuites, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return p.Parse(file)
}

func (p *XMLParser) ParseURL(url string) (*models.TestSuites, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get url: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get url: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return p.Parse(data)
}

func NewParser() Parser {
	return &XMLParser{}
}
