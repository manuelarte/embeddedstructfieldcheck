package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/manuelarte/embeddedcheck/internal"
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
	insp, found := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !found {
		//nolint:nilnil // impossible case.
		return nil, nil
	}

	nodeFilter := []ast.Node{
		(*ast.TypeSpec)(nil),
	}

	a := internal.NewStructAnalyzer()

	insp.Preorder(nodeFilter, func(n ast.Node) {
		if node, isStruct := n.(*ast.TypeSpec); isStruct {
			if diag, isNotValid := a.Analyze(node); isNotValid {
				pass.Report(diag)
			}
		}
	})

	//nolint:nilnil //any, error
	return nil, nil
}
