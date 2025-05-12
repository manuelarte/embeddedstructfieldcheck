package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		desc     string
		patterns string
	}{
		{
			desc:     "simple",
			patterns: "simple",
		},
		{
			desc:     "comments",
			patterns: "comments",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			a := NewAnalyzer()

			analysistest.Run(t, analysistest.TestData(), a, test.patterns)
		})
	}
}

func TestAnalyzerWithFix(t *testing.T) {
	testCases := []struct {
		desc     string
		patterns string
	}{
		{
			desc:     "simple fix",
			patterns: "simple-fix",
		},
		{
			desc:     "comments fix",
			patterns: "comments-fix",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			a := NewAnalyzer()

			analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), a, test.patterns)
		})
	}
}
