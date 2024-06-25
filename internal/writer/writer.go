package writer

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Writer interface {
	WriteToFile(data interface{}, path string) error
	WriteToBytes(data interface{}) ([]byte, error)
}

type XMLWriter struct{}

func (w *XMLWriter) WriteToFile(data interface{}, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "    ")

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode data: %w", err)
	}

	return nil
}

func (w *XMLWriter) WriteToBytes(data interface{}) ([]byte, error) {
	bytes, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	return bytes, nil
}

func NewWriter() Writer {
	return &XMLWriter{}
}
