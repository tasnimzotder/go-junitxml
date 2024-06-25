package parser

// func TestSetTestCaseStatus(t *testing.T) {
// 	tc := &models.TestCase{}

// 	// Test case with failure
// 	tc.Failure = &models.Failure{
// 		Message: "Failed",
// 	}
// 	setTestCaseStatus(tc)
// 	assert.Equal(t, "failed", tc.Status)

// 	// Test case with error
// 	tc.Failure = &models.Failure{
// 		Message: "Error",
// 	}
// 	setTestCaseStatus(tc)
// 	assert.Equal(t, "errored", tc.Status)

// 	// Test case with skipped
// 	tc.Failure = nil
// 	tc.Skipped = &models.Skipped{}
// 	setTestCaseStatus(tc)
// 	assert.Equal(t, "skipped", tc.Status)

// 	// Test case with no failure or skipped
// 	tc.Failure = nil
// 	tc.Skipped = nil
// 	setTestCaseStatus(tc)
// 	assert.Equal(t, "passed", tc.Status)
// }
