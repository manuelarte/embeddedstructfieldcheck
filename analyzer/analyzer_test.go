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
				EmptyLineCheck:   "true",
				ForbidMutexCheck: "false",
			},
		},
		{
			desc:     "comments",
			patterns: "comments",
			options: map[string]string{
				EmptyLineCheck:   "true",
				ForbidMutexCheck: "false",
			},
		},
		{
			desc:     "comments empty-line false",
			patterns: "comments-empty-line",
			options: map[string]string{
				EmptyLineCheck:   "false",
				ForbidMutexCheck: "false",
			},
		},
		{
			desc:     "mutex",
			patterns: "mutex",
			options: map[string]string{
				EmptyLineCheck:   "true",
				ForbidMutexCheck: "true",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			a := NewAnalyzer()

			for k, v := range test.options {
				err := a.Flags.Set(k, v)
				if err != nil {
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
		options  map[string]string
	}{
		{
			desc:     "simple fix",
			patterns: "simple-fix",
			options: map[string]string{
				EmptyLineCheck:   "true",
				ForbidMutexCheck: "false",
			},
		},
		{
			desc:     "simple empty-line fix",
			patterns: "simple-empty-line-fix",
			options: map[string]string{
				EmptyLineCheck:   "false",
				ForbidMutexCheck: "false",
			},
		},
		{
			desc:     "comments fix",
			patterns: "comments-fix",
			options: map[string]string{
				EmptyLineCheck:   "true",
				ForbidMutexCheck: "false",
			},
		},
		{
			desc:     "comments empty-line fix",
			patterns: "comments-empty-line-fix",
			options: map[string]string{
				EmptyLineCheck:   "false",
				ForbidMutexCheck: "false",
			},
		},
		{
			desc:     "block comments (disabling empty-line)",
			patterns: "block-comments",
			options: map[string]string{
				EmptyLineCheck:   "false",
				ForbidMutexCheck: "false",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			a := NewAnalyzer()

			for k, v := range test.options {
				err := a.Flags.Set(k, v)
				if err != nil {
					t.Fatal(err)
				}
			}

			analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), a, test.patterns)
		})
	}
}
