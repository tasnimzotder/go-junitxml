package parser

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tasnimzotder/go-junitxml/models"
)

var (
	xmlData = []byte(`
        <testsuites>
            <testsuite name="Suite1" tests="2" failures="1" time="0.123">
                <testcase name="Test1" classname="TestClass1" time="0.045">
                    <failure message="Assertion failed" type="AssertionError">Detailed failure message</failure>
                </testcase>
                <testcase name="Test2" classname="TestClass2" time="0.078"/>
            </testsuite>
            <testsuite name="Suite2" tests="1" failures="0" time="0.056">
                <testcase name="Test3" classname="TestClass3" time="0.056"/>
            </testsuite>
        </testsuites>
    `)
)

func TestXMLParser_Parse(t *testing.T) {
	parser := NewParser()
	testSuites, err := parser.Parse(xmlData)

	assert.Nil(t, err)
	assert.NotNil(t, testSuites)

	checkSuites(t, testSuites)
}

func TestXMLParser_ParseFailed(t *testing.T) {
	parser := NewParser()
	_, err := parser.Parse([]byte("invalid xml"))

	assert.NotNil(t, err)
}

func TestXMLParser_ParseTotals(t *testing.T) {
	time.Sleep(102 * time.Millisecond)

	parser := NewParser()
	testSuites, err := parser.Parse(xmlData)

	assert.Nil(t, err)
	assert.NotNil(t, testSuites)

	assert.Equal(t, 2, testSuites.Totals.CountTestSuites)
	assert.Equal(t, 3, testSuites.Totals.CountTestCases)
	assert.Equal(t, "123ms", testSuites.TestSuites[0].Totals.Duration.String())

	assert.Equal(t, 1, testSuites.TestSuites[1].Totals.Passed)
	assert.Equal(t, 0, testSuites.TestSuites[1].Totals.Failed)
	assert.Equal(t, 0, testSuites.TestSuites[1].Totals.Errored)
	assert.Equal(t, 0, testSuites.TestSuites[1].Totals.Skipped)
	assert.Equal(t, "56ms", testSuites.TestSuites[1].Totals.Duration.String())
}

func TestXMLParser_ParseFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test-*.xml")
	assert.Nil(t, err)
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(xmlData); err != nil {
		assert.Nil(t, err)
	}

	if err := tempFile.Close(); err != nil {
		assert.Nil(t, err)
	}

	parser := NewParser()
	testSuites, err := parser.ParseFile(tempFile.Name())

	assert.Nil(t, err)
	assert.NotNil(t, testSuites)

	checkSuites(t, testSuites)
}

func TestXMLParser_ParseFileFailed(t *testing.T) {
	parser := NewParser()
	_, err := parser.ParseFile("invalid-file")

	assert.NotNil(t, err)
}

func TestXMLParser_ParseURL(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/xml")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(xmlData))
			},
		),
	)
	defer server.Close()

	parser := NewParser()

	testSuites, err := parser.ParseURL(server.URL)

	assert.Nil(t, err)
	assert.NotNil(t, testSuites)

	checkSuites(t, testSuites)
}

func TestXMLParser_ParseURLFailed(t *testing.T) {
	parser := NewParser()
	_, err := parser.ParseURL("invalid-url")

	assert.NotNil(t, err)
}

func TestXMLParser_ParseURLInvalidContentType(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("invalid content"))
			},
		),
	)
	defer server.Close()

	parser := NewParser()
	_, err := parser.ParseURL(server.URL)

	assert.NotNil(t, err)
}

func checkSuites(t *testing.T, testSuites *models.TestSuites) {
	assert.NotNil(t, testSuites)
	assert.Equal(t, 2, len(testSuites.TestSuites))

	suite1 := testSuites.TestSuites[0]

	assert.Equal(t, "Suite1", suite1.Name)
	assert.Equal(t, 2, suite1.Tests)
	assert.Equal(t, 1, suite1.Failures)
	assert.Equal(t, "0.123", suite1.Time)
	assert.Equal(t, 2, len(suite1.TestCases))

	testCase1 := suite1.TestCases[0]

	assert.Equal(t, "Test1", testCase1.Name)
	assert.Equal(t, "TestClass1", testCase1.Classname)
	assert.Equal(t, "0.045", testCase1.Time)
	assert.NotNil(t, testCase1.Failure)
	assert.Equal(t, "Assertion failed", testCase1.Failure.Message)
	assert.Equal(t, "AssertionError", testCase1.Failure.Type)
	assert.Equal(t, "Detailed failure message", testCase1.Failure.Content)
}
