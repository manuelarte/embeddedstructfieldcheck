package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"

	"github.com/manuelarte/embeddedcheck/internal/structanalyzer"
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
	fset := pass.Fset
	for _, file := range pass.Files {
		var a structanalyzer.StructAnalyzer
		ast.Inspect(file, func(n ast.Node) bool {
			switch node := n.(type) {
			case *ast.StructType:
				if diag, ok := a.Analyze(fset); ok {
					pass.Report(diag)
				}
				a = structanalyzer.New(node)
			default:
				if a.IsAnalyzingStruct() && node != nil && node.End() <= a.GetEndPos() {
					a.CheckNode(n)
				}
			}
			return true
		})
		if diag, ok := a.Analyze(fset); ok {
			pass.Report(diag)
		}
	}

	//nolint:nilnil //any, error
	return nil, nil
}
