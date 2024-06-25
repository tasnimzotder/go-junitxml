package validator

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Validator interface {
	Validate(data []byte) error
	ValidateFile(path string) error
	ValidateURL(url string) error
}

type XMLValidator struct{}

func (v *XMLValidator) Validate(data []byte) error {
	var rootElement struct {
		XMLName xml.Name `xml:"testsuites"`
	}

	if err := xml.Unmarshal(data, &rootElement); err != nil {
		return fmt.Errorf("failed to unmarshal xml: %w", err)
	}

	if rootElement.XMLName.Local != "testsuites" {
		return fmt.Errorf("invalid root element: %s", rootElement.XMLName.Local)
	}

	return nil
}

func (v *XMLValidator) ValidateFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return v.Validate(data)
}

func (v *XMLValidator) ValidateURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to get url: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get url: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	return v.Validate(data)
}

func NewValidator() Validator {
	return &XMLValidator{}
}
