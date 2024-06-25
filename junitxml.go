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

func NewJUnitXML() *JUnitXML {
	return &JUnitXML{
		parser:    parser.NewParser(),
		validator: validator.NewValidator(),
		writer:    writer.NewWriter(),
		merger:    merger.NewMerger(),
	}
}

// parser methods
func (j *JUnitXML) Parse(data []byte) (*models.TestSuites, error) {
	return j.parser.Parse(data)
}

func (j *JUnitXML) ParseFile(path string) (*models.TestSuites, error) {
	return j.parser.ParseFile(path)
}

func (j *JUnitXML) ParseURL(url string) (*models.TestSuites, error) {
	return j.parser.ParseURL(url)
}

// validator methods
func (j *JUnitXML) Validate(data []byte) error {
	return j.validator.Validate(data)
}

func (j *JUnitXML) ValidateFile(path string) error {
	return j.validator.ValidateFile(path)
}

func (j *JUnitXML) ValidateURL(url string) error {
	return j.validator.ValidateURL(url)
}

// writer methods
func (j *JUnitXML) WriteToFile(data *models.TestSuites, path string) error {
	return j.writer.WriteToFile(data, path)
}

// merger methods
func (j *JUnitXML) Merge(suites ...*models.TestSuites) (*models.TestSuites, error) {
	return j.merger.Merge(suites...)
}

func (j *JUnitXML) MergeWithOpts(opts *merger.MergeOptions, suites ...*models.TestSuites) (*models.TestSuites, error) {
	return j.merger.MergeWithOpts(opts, suites...)
}

func (j *JUnitXML) MergeFiles(paths ...string) (*models.TestSuites, error) {
	return j.merger.MergeFiles(paths...)
}

func (j *JUnitXML) MergeFilesWithOpts(opts *merger.MergeOptions, paths ...string) (*models.TestSuites, error) {
	return j.merger.MergeFilesWithOpts(opts, paths...)
}
