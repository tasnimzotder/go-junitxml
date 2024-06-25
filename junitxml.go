package junitxml

import (
	"github.com/tasnimzotder/go-junitxml/internal/merger"
	"github.com/tasnimzotder/go-junitxml/internal/models"
	"github.com/tasnimzotder/go-junitxml/internal/parser"
	"github.com/tasnimzotder/go-junitxml/internal/validator"
	"github.com/tasnimzotder/go-junitxml/internal/writer"
)

type TestSuites = models.TestSuites
type TestSuite = models.TestSuite
type TestCase = models.TestCase
type MergeOptions = merger.MergeOptions

type JUnitXML struct {
	parser    parser.Parser
	validator validator.Validator
	writer    writer.Writer
	merger    merger.Merger
}

// NewJUnitXML creates a new instance of JUnitXML.
// It initializes the parser, validator, writer, and merger.
func NewJUnitXML() *JUnitXML {
	return &JUnitXML{
		parser:    parser.NewParser(),
		validator: validator.NewValidator(),
		writer:    writer.NewWriter(),
		merger:    merger.NewMerger(),
	}
}

// parser methods

// Parse parses the given XML data and returns a TestSuites object or an error.
// It checks the root element of the XML and delegates the parsing to the appropriate method.
func (j *JUnitXML) Parse(data []byte) (*models.TestSuites, error) {
	return j.parser.Parse(data)
}

// ParseFile reads the file at the given path and parses its contents as XML to generate a TestSuites object.
// It returns the parsed TestSuites object and any error encountered during the process.
func (j *JUnitXML) ParseFile(path string) (*models.TestSuites, error) {
	return j.parser.ParseFile(path)
}

// ParseURL parses the XML content from the specified URL and returns the parsed TestSuites.
// It performs an HTTP GET request to retrieve the XML content and then parses it using the Parse method.
// If any error occurs during the process, it returns nil and the corresponding error.
func (j *JUnitXML) ParseURL(url string) (*models.TestSuites, error) {
	return j.parser.ParseURL(url)
}

// validator methods

// Validate validates the given XML data.
// It unmarshals the XML data and checks the root element to determine the validation strategy.
// The XML data should be provided as a byte slice.
// Returns an error if the XML data is invalid or if the root element is unexpected.
func (j *JUnitXML) Validate(data []byte) error {
	return j.validator.Validate(data)
}

// ValidateFile reads the file at the specified path and validates its contents using the XMLValidator.
// It returns an error if the file cannot be read or if the validation fails.
func (j *JUnitXML) ValidateFile(path string) error {
	return j.validator.ValidateFile(path)
}

// ValidateURL validates the XML content from the specified URL.
// It sends an HTTP GET request to the URL, reads the response body,
// and then validates the XML data using the XMLValidator's Validate method.
// If any error occurs during the process, it returns an error.
func (j *JUnitXML) ValidateURL(url string) error {
	return j.validator.ValidateURL(url)
}

// writer methods

// Write writes the TestSuites object to an XML file at the specified path.
// It returns an error if the writing process fails.
func (j *JUnitXML) WriteToFile(data *models.TestSuites, path string) error {
	return j.writer.WriteToFile(data, path)
}

// merger methods

// Merge merges the provided TestSuites objects into a single TestSuites object.
// It returns the merged TestSuites object and any error encountered during the process.
func (j *JUnitXML) Merge(suites ...*models.TestSuites) (*models.TestSuites, error) {
	return j.merger.Merge(suites...)
}

// MergeWithOpts merges the provided TestSuites objects into a single TestSuites object with the given options.
// It returns the merged TestSuites object and any error encountered during the process.
func (j *JUnitXML) MergeWithOpts(opts *merger.MergeOptions, suites ...*models.TestSuites) (*models.TestSuites, error) {
	return j.merger.MergeWithOpts(opts, suites...)
}

// MergeFiles merges the XML files at the specified paths into a single TestSuites object.
// It returns the merged TestSuites object and any error encountered during the process.
func (j *JUnitXML) MergeFiles(paths ...string) (*models.TestSuites, error) {
	return j.merger.MergeFiles(paths...)
}

// MergeFilesWithOpts merges the XML files at the specified paths into a single TestSuites object with the given options.
// It returns the merged TestSuites object and any error encountered during the process.
func (j *JUnitXML) MergeFilesWithOpts(opts *merger.MergeOptions, paths ...string) (*models.TestSuites, error) {
	return j.merger.MergeFilesWithOpts(opts, paths...)
}
