package utils

import (
	"fmt"

	"github.com/tasnimzotder/go-junitxml/internal/models"
)

// CalculateTotalsRootSuite calculates the totals for the root test suite and updates the provided TestSuites object.
// It iterates over each test suite, calculates the totals for test cases, and updates the overall totals.
// The function returns the updated TestSuites object with the calculated totals or an error if the calculation fails.
func CalculateTotalsRootSuite(ts *models.TestSuites) (*models.TestSuites, error) {
	totals := models.TotalsTestSuites{}

	for i := range ts.TestSuites {
		_, err := calculateTotalsTestSuite(&ts.TestSuites[i])
		if err != nil {
			return nil, fmt.Errorf("failed to calculate totals: %w", err)
		}

		if ts.TestSuites[i].Status == models.StatusPassed {
			totals.PassedTestSuites += 1
		} else {
			totals.FailedTestSuites += 1
		}

		totals.CountTestSuites += 1

		// calculate totals for test cases
		totals.CountTestCases += ts.TestSuites[i].Totals.Count
		totals.PassedTestCases += ts.TestSuites[i].Totals.Passed
		totals.FailedTestCases += ts.TestSuites[i].Totals.Failed
		totals.SkippedTestCases += ts.TestSuites[i].Totals.Skipped
		totals.ErroredTestCases += ts.TestSuites[i].Totals.Errored

		totals.Duration += ts.TestSuites[i].Totals.Duration
	}

	// match the total number of test suites
	if totals.CountTestSuites != len(ts.TestSuites) {
		return nil, fmt.Errorf("failed to calculate totals: mismatched test suites count")
	}

	ts.Totals = totals

	return ts, nil
}

// calculateTotalsTestSuite calculates the totals for a given test suite.
// It takes a pointer to a TestSuite struct as input and returns a pointer to the updated TestSuite struct and an error, if any.
// The function calculates the total count of test cases, the number of passed, failed, errored, and skipped test cases,
// and the total duration of all test cases in the test suite.
// It also updates the status of the test suite based on the test case statuses.
// If the totals are mismatched, an error is returned.
func calculateTotalsTestSuite(ts *models.TestSuite) (*models.TestSuite, error) {
	statusPassedFlag := true
	totals := models.TotalsTestCases{}

	// check if totals are matching
	if ts.Tests != 0 && ts.Tests != len(ts.TestCases) {
		return nil, fmt.Errorf("failed to calculate totals: mismatched test cases count")
	}

	totals.Count = len(ts.TestCases)

	for i := range ts.TestCases {
		setTestCaseStatus(&ts.TestCases[i])

		switch ts.TestCases[i].Status {
		case models.StatusPassed:
			totals.Passed += 1
		case models.StatusFailed:
			totals.Failed += 1
			statusPassedFlag = false
		case models.StatusErrored:
			totals.Errored += 1
			statusPassedFlag = false
		case models.StatusSkipped:
			totals.Skipped += 1
			// statusPassedFlag = true
		}

		ts.TestCases[i].Duration = StringToDuration(ts.TestCases[i].Time)
		totals.Duration += ts.TestCases[i].Duration
	}

	ts.Totals = totals

	if statusPassedFlag {
		ts.Status = models.StatusPassed
	} else {
		ts.Status = models.StatusFailed
	}

	return ts, nil
}

// setTestCaseStatus sets the status of a test case based on its failure, skipped, or passed status.
// It takes a pointer to a TestCase struct as input and modifies its Status field accordingly.
// If the test case has a failure with the message "Failed", the status is set to StatusFailed.
// If the test case has a failure with a message other than "Failed", the status is set to StatusErrored.
// If the test case is skipped, the status is set to StatusSkipped.
// If none of the above conditions are met, the status is set to StatusPassed.
func setTestCaseStatus(tc *models.TestCase) {
	if tc.Failure != nil {
		if tc.Failure.Message == "Failed" {
			tc.Status = models.StatusFailed
		} else {
			tc.Status = models.StatusErrored
		}
	} else if tc.Skipped != nil {
		tc.Status = models.StatusSkipped
	} else {
		tc.Status = models.StatusPassed
	}
}
