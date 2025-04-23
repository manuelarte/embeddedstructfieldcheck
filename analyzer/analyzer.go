package analyzer

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func NewAnalyzer() *analysis.Analyzer {
	l := embeddedcheck{}

	a := &analysis.Analyzer{
		Name: "embeddedcheck",
		Doc: "Embedded types should be at the top of the field list of a struct, " +
			"and there must be an empty line separating embedded fields from regular fields. methods, and constructors",
		Run:      l.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	return a
}

type embeddedcheck struct{}

func (l embeddedcheck) run(pass *analysis.Pass) (any, error) {
	return nil, nil
}
