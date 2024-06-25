package merger

import (
	"fmt"

	"github.com/tasnimzotder/go-junitxml/internal/models"
	"github.com/tasnimzotder/go-junitxml/internal/parser"
	"github.com/tasnimzotder/go-junitxml/internal/utils"
	"github.com/tasnimzotder/go-junitxml/internal/validator"
)

type Merger interface {
	Merge(suites ...*models.TestSuites) (*models.TestSuites, error)
	MergeWithOpts(opts *MergeOptions, suites ...*models.TestSuites) (*models.TestSuites, error)
	MergeFiles(paths ...string) (*models.TestSuites, error)
	MergeFilesWithOpts(opts *MergeOptions, paths ...string) (*models.TestSuites, error)
}

type XMLTestSuitesMerger struct{}

func (m *XMLTestSuitesMerger) Merge(suites ...*models.TestSuites) (*models.TestSuites, error) {
	return m.MergeWithOpts(parseOptions(), suites...)
}

func (m *XMLTestSuitesMerger) MergeWithOpts(opts *MergeOptions, suites ...*models.TestSuites) (*models.TestSuites, error) {
	if len(suites) == 0 {
		return nil, fmt.Errorf("no test suites to merge")
	}

	validator := validator.NewValidator()
	var merged *models.TestSuites
	var err error

	if opts.SingleTestSuite {
		merged, err = m.mergeIntoSingleSuite(suites)
		if err != nil {
			return nil, fmt.Errorf("failed to merge into single suite: %w", err)
		}
	} else if !opts.SingleTestSuite {
		merged, err = m.mergeIntoMultipleSuites(suites)
		if err != nil {
			return nil, fmt.Errorf("failed to merge into multiple suites: %w", err)
		}
	}

	_, err = utils.CalculateTotalsRootSuite(merged)
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
	return m.MergeFilesWithOpts(parseOptions(), paths...)
}

func (m *XMLTestSuitesMerger) MergeFilesWithOpts(opts *MergeOptions, paths ...string) (*models.TestSuites, error) {
	parser := parser.NewParser()
	validator := validator.NewValidator()

	var allSuites []*models.TestSuites

	for _, path := range paths {
		// validate the file
		err := validator.ValidateFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to validate file: %w", err)
		}

		suites, err := parser.ParseFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to parse file: %w", err)
		}

		allSuites = append(allSuites, suites)
	}

	// if opts.SingleTestSuite {
	// 	return m.Merge(allSuites...)
	// }

	return m.MergeWithOpts(opts, allSuites...)
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
