package writer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tasnimzotder/go-junitxml/internal/models"
)

var (
	testData = models.TestSuite{
		Name:     "Suite1",
		Tests:    2,
		Failures: 1,
		Errors:   0,
		Skipped:  0,
		Time:     "0.123",
		TestCases: []models.TestCase{
			{
				Name:      "Test1",
				Classname: "TestClass1",
				Time:      "0.045",
				Failure: &models.Failure{
					Message: "Assertion failed",
					Type:    "AssertionError",
					Content: "Details of the failure",
				},
			},
			{
				Name:      "Test2",
				Classname: "TestClass2",
				Time:      "0.078",
			},
		},
	}
)

func TestXMLWriter_WriteToFile(t *testing.T) {
	writer := NewWriter()
	tempFile, err := os.CreateTemp("", "test-*.xml")
	assert.Nil(t, err)
	assert.NotNil(t, tempFile)
	defer os.Remove(tempFile.Name())

	err = writer.WriteToFile(testData, tempFile.Name())
	assert.Nil(t, err)

	// read the file and check the content
	writtenData, err := os.ReadFile(tempFile.Name())
	assert.Nil(t, err)
	assert.NotNil(t, writtenData)

	expectedData, err := writer.WriteToBytes(testData)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedData), string(writtenData))

	// if !reflect.DeepEqual(string(expectedData), string(writtenData)) {
	// 	t.Errorf("expected: %s, got: %s", string(expectedData), string(writtenData))
	// }
}

func TestXMLWriter_WriteToBytes(t *testing.T) {
	writer := NewWriter()
	xmlBytes, err := writer.WriteToBytes(testData)
	assert.Nil(t, err)
	assert.NotNil(t, xmlBytes)

	// todo: implement
}
