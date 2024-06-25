package merger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tasnimzotder/go-junitxml/internal/models"
	"github.com/tasnimzotder/go-junitxml/internal/writer"
)

var (
	suite1 = &models.TestSuites{
		TestSuites: []models.TestSuite{
			{
				Name:     "Suite1",
				Tests:    2,
				Failures: 1,
				Time:     "0.123",
				TestCases: []models.TestCase{
					{
						Name: "Test1",
						Time: "0.045",
						Failure: &models.Failure{
							Message: "Failed",
						},
					},
					{
						Name: "Test2",
						Time: "0.078",
					},
				},
			},
		},
	}

	suite2 = &models.TestSuites{
		TestSuites: []models.TestSuite{
			{
				Name:     "Suite2",
				Tests:    1,
				Failures: 0,
				Time:     "0.056",
				TestCases: []models.TestCase{
					{
						Name: "Test3",
						Time: "0.056",
					},
				},
			},
		},
	}
)

func TestXMLTestSuitesMerger_Merge(t *testing.T) {
	merger := NewMerger()

	// Merge the test suites
	merged, err := merger.Merge(suite1, suite2)
	assert.Nil(t, err)
	assert.NotNil(t, merged)

	// Assert the merged test suites
	assert.Equal(t, 2, len(merged.TestSuites))
	assert.Equal(t, 3, merged.Tests)
	assert.Equal(t, 1, merged.Failures)
	assert.Equal(t, 0, merged.Errors)
	assert.Equal(t, "0.179", merged.Time)
}
func TestXMLTestSuitesMerger_MergeFiles(t *testing.T) {
	merger := NewMerger()
	writer := writer.NewWriter()

	// create temp files from the test suites
	file1, err := os.CreateTemp("", "file1*.xml")
	assert.Nil(t, err)
	defer os.Remove(file1.Name())

	file2, err := os.CreateTemp("", "file2*.xml")
	assert.Nil(t, err)
	defer os.Remove(file2.Name())

	// write the test suites to the temp files
	err = writer.WriteToFile(suite1, file1.Name())
	assert.Nil(t, err)

	err = writer.WriteToFile(suite2, file2.Name())
	assert.Nil(t, err)

	// Merge the test suites
	merged, err := merger.MergeFiles(file1.Name(), file2.Name())
	assert.Nil(t, err)
	assert.NotNil(t, merged)

	// Assert the merged test suites
	assert.Equal(t, 2, len(merged.TestSuites))
	assert.Equal(t, 3, merged.Tests)
	assert.Equal(t, 1, merged.Failures)
	assert.Equal(t, 0, merged.Errors)
	assert.Equal(t, "0.179", merged.Time)
}
