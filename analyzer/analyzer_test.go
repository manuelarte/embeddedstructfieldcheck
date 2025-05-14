package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		desc     string
		patterns string
		options  map[string]string
	}{
		{
			desc:     "simple",
			patterns: "simple",
			options: map[string]string{
				ForbidMutexName: "false",
			},
		},
		{
			desc:     "comments",
			patterns: "comments",
			options: map[string]string{
				ForbidMutexName: "false",
			},
		},
		{
			desc:     "mutex",
			patterns: "mutex",
			options: map[string]string{
				ForbidMutexName: "true",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			a := NewAnalyzer()

			for k, v := range test.options {
				if err := a.Flags.Set(k, v); err != nil {
					t.Fatal(err)
				}
			}

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
