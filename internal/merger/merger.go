package merger

import (
	"fmt"

	"github.com/tasnimzotder/go-junitxml/internal/parser"
	"github.com/tasnimzotder/go-junitxml/internal/utils"
	"github.com/tasnimzotder/go-junitxml/internal/validator"
	"github.com/tasnimzotder/go-junitxml/models"
)

type Merger interface {
	Merge(suites ...*models.TestSuites) (*models.TestSuites, error)
	MergeFiles(paths ...string) (*models.TestSuites, error)
}

type XMLTestSuitesMerger struct{}

func (m *XMLTestSuitesMerger) Merge(suites ...*models.TestSuites) (*models.TestSuites, error) {
	if len(suites) == 0 {
		return nil, fmt.Errorf("no test suites to merge")
	}

	validator := validator.NewValidator()

	merged := &models.TestSuites{}
	for _, suite := range suites {
		merged.TestSuites = append(merged.TestSuites, suite.TestSuites...)
	}

	_, err := utils.CalculateTotalsRootSuite(merged)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate totals: %w", err)
	}

	updateRootSuiteStats(merged)

	// validate the merged test suites
	err = validator.Validate([]byte(merged.String()))
	if err != nil {
		return nil, fmt.Errorf("failed to validate merged test suites: %w", err)
	}

	return merged, nil
}

func (m *XMLTestSuitesMerger) MergeFiles(paths ...string) (*models.TestSuites, error) {
	parser := parser.NewParser()
	validator := validator.NewValidator()
	suites := make([]*models.TestSuites, 0)

	for _, path := range paths {
		// validate the file
		err := validator.ValidateFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to validate file: %w", err)
		}

		ts, err := parser.ParseFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to parse file: %w", err)
		}

		suites = append(suites, ts)
	}

	return m.Merge(suites...)
}

func NewMerger() *XMLTestSuitesMerger {
	return &XMLTestSuitesMerger{}
}

func updateRootSuiteStats(root *models.TestSuites) {
	root.Tests = root.Totals.CountTestCases
	root.Failures = root.Totals.FailedTestCases
	root.Errors = root.Totals.ErroredTestCases
	root.Time = utils.Float64ToSting(root.Totals.Duration.Seconds())
}
