package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tasnimzotder/go-junitxml/internal/models"
)

func TestSetTestCaseStatus(t *testing.T) {
	// test case with failure
	tc := &models.TestCase{
		Failure: &models.Failure{
			Message: "Failed",
		},
	}
	setTestCaseStatus(tc)
	assert.Equal(t, "failed", tc.Status)

	// test case with error
	tc = &models.TestCase{
		Failure: &models.Failure{
			Message: "Error",
		},
	}
	setTestCaseStatus(tc)
	assert.Equal(t, "errored", tc.Status)

	// test case with skipped
	tc = &models.TestCase{
		Skipped: &models.Skipped{},
	}
	setTestCaseStatus(tc)
	assert.Equal(t, "skipped", tc.Status)

	// test case with no failure or skipped
	tc = &models.TestCase{}
	setTestCaseStatus(tc)
	assert.Equal(t, "passed", tc.Status)
}

func TestCalculateTotalsTestSuite(t *testing.T) {
	ts := &models.TestSuite{
		Tests: 3,
		TestCases: []models.TestCase{
			{
				Status: "passed",
				Time:   "1s",
			},
			{
				Status: "failed",
				Time:   "2s",
				Failure: &models.Failure{
					Message: "Failed",
				},
			},
			{
				Status:  "skipped",
				Time:    "3s",
				Skipped: &models.Skipped{},
			},
		},
	}

	expectedTotals := models.TotalsTestCases{
		Count:    3,
		Passed:   1,
		Failed:   1,
		Errored:  0,
		Skipped:  1,
		Duration: 6 * time.Second,
	}

	_, err := calculateTotalsTestSuite(ts)

	assert.NoError(t, err)
	assert.Equal(t, expectedTotals, ts.Totals)
}

func TestCalculateTotalsRootSuite(t *testing.T) {
	// test case with all passed test suites
	ts := &models.TestSuites{
		TestSuites: []models.TestSuite{
			{
				TestCases: []models.TestCase{
					{
						Status: "passed",
						Time:   "0.1",
					},
					{
						Status: "passed",
						Time:   "0.2",
					},
				},
			},
			{
				TestCases: []models.TestCase{
					{
						Status: "passed",
						Time:   "0.3",
					},
				},
			},
		},
	}

	expectedTotals := models.TotalsTestSuites{
		CountTestSuites:  2,
		PassedTestSuites: 2,
		FailedTestSuites: 0,
		CountTestCases:   3,
		PassedTestCases:  3,
		FailedTestCases:  0,
		SkippedTestCases: 0,
		ErroredTestCases: 0,
		Duration:         time.Duration(600) * time.Millisecond,
	}

	_, err := CalculateTotalsRootSuite(ts)
	assert.NoError(t, err)
	assert.Equal(t, expectedTotals, ts.Totals)

	// test case with failed test suites
	ts = &models.TestSuites{
		TestSuites: []models.TestSuite{
			{
				TestCases: []models.TestCase{
					{
						Status: "passed",
					},
					{
						Status: "failed",
						Failure: &models.Failure{
							Message: "Failed",
						},
					},
				},
			},
			{
				TestCases: []models.TestCase{
					{
						Status: "passed",
					},
				},
			},
		},
	}

	expectedTotals = models.TotalsTestSuites{
		CountTestSuites:  2,
		PassedTestSuites: 1,
		FailedTestSuites: 1,
		CountTestCases:   3,
		PassedTestCases:  2,
		FailedTestCases:  1,
		SkippedTestCases: 0,
		ErroredTestCases: 0,
		Duration:         time.Duration(0),
	}
	_, err = CalculateTotalsRootSuite(ts)
	assert.NoError(t, err)
	assert.Equal(t, expectedTotals, ts.Totals)

	// test case with mismatch test cases count
	ts = &models.TestSuites{
		TestSuites: []models.TestSuite{
			{
				Tests: 2,
				TestCases: []models.TestCase{
					{
						Status: "passed",
					},
				},
			},
		},
	}

	_, err = CalculateTotalsRootSuite(ts)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "mismatched test cases count")
}
