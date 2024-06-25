package models

import (
	"encoding/xml"
	"time"
)

// status enum
const (
	StatusPassed  = "passed"
	StatusFailed  = "failed"
	StatusErrored = "errored"
	StatusSkipped = "skipped"
)

type TestSuites struct {
	XMLName    xml.Name         `xml:"testsuites"`
	TestSuites []TestSuite      `xml:"testsuite"`
	Tests      int              `xml:"tests,attr"`
	Failures   int              `xml:"failures,attr"`
	Errors     int              `xml:"errors,attr"`
	Time       string           `xml:"time,attr"`
	Totals     TotalsTestSuites `xml:"-"` // This field is not part of the XML
}

func (ts *TestSuites) String() string {
	// return as xml string
	output, err := xml.MarshalIndent(ts, "", "  ")
	if err != nil {
		return ""
	}

	return string(output)
}

type TestSuite struct {
	XMLName   xml.Name        `xml:"testsuite"`
	Name      string          `xml:"name,attr"`
	Tests     int             `xml:"tests,attr"`
	Failures  int             `xml:"failures,attr"`
	Errors    int             `xml:"errors,attr"`
	Skipped   int             `xml:"skipped,attr"`
	Time      string          `xml:"time,attr"`
	Timestamp string          `xml:"timestamp,attr"`
	Filename  string          `xml:"filename,attr"`
	Status    string          `xml:"-"` // The value can be "passed", "failed", "errored", or "skipped"
	TestCases []TestCase      `xml:"testcase"`
	Totals    TotalsTestCases `xml:"-"` // This field is not part of the XML
}

type TestCase struct {
	XMLName   xml.Name      `xml:"testcase"`
	Name      string        `xml:"name,attr"`
	Classname string        `xml:"classname,attr"`
	Time      string        `xml:"time,attr"`
	Status    string        `xml:"-"` // The value can be "passed", "failed", "errored", or "skipped"
	Failure   *Failure      `xml:"failure,omitempty"`
	Skipped   *Skipped      `xml:"skipped,omitempty"`
	Duration  time.Duration `xml:"-"` // This field is not part of the XML
}

type Failure struct {
	XMLName xml.Name `xml:"failure"`
	Message string   `xml:"message,attr"`
	Type    string   `xml:"type,attr"`
	Content string   `xml:",chardata"`
}

type Skipped struct {
	XMLName xml.Name `xml:"skipped"`
	Message string   `xml:"message,attr"`
}

type TotalsTestCases struct {
	Count    int
	Passed   int
	Failed   int
	Errored  int
	Skipped  int
	Duration time.Duration
}

type TotalsTestSuites struct {
	CountTestSuites  int
	CountTestCases   int
	PassedTestSuites int
	PassedTestCases  int
	FailedTestSuites int
	FailedTestCases  int
	SkippedTestCases int
	ErroredTestCases int
	Duration         time.Duration
}
