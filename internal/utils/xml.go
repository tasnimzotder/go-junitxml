package utils

import (
	"fmt"

	"github.com/tasnimzotder/go-junitxml/models"
)

func CalculateTotalsRootSuite(ts *models.TestSuites) (*models.TestSuites, error) {
	totals := models.TotalsTestSuites{}

	for i := range ts.TestSuites {
		_, err := calculateTotalsTestSuite(&ts.TestSuites[i])
		if err != nil {
			return nil, fmt.Errorf("failed to calculate totals: %w", err)
		}

		// calculate totals for test suites
		// for j := range ts.TestSuites[i].TestCases {
		// 	if ts.TestSuites[i].TestCases[j].Status != "passed" {
		// 		isPassedFlag = false
		// 		break
		// 	}
		// }

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
