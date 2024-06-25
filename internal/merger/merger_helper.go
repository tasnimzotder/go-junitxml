package merger

import "github.com/tasnimzotder/go-junitxml/internal/models"

func (m *XMLTestSuitesMerger) mergeIntoSingleSuite(suites []*models.TestSuites) (*models.TestSuites, error) {
	merged := &models.TestSuites{}
	var singleSuite models.TestSuite

	for _, suite := range suites {
		for _, testSuite := range suite.TestSuites {
			singleSuite.TestCases = append(singleSuite.TestCases, testSuite.TestCases...)
		}
	}

	merged.TestSuites = append(merged.TestSuites, singleSuite)

	return merged, nil
}

func (m *XMLTestSuitesMerger) mergeIntoMultipleSuites(suites []*models.TestSuites) (*models.TestSuites, error) {
	merged := &models.TestSuites{}

	for _, suite := range suites {
		merged.TestSuites = append(merged.TestSuites, suite.TestSuites...)
	}

	return merged, nil
}
