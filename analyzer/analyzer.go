package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/manuelarte/embeddedstructfieldcheck/internal"
)

func NewAnalyzer() *analysis.Analyzer {
	l := embeddedstructfieldcheck{}

	a := &analysis.Analyzer{
		Name: "embeddedstructfieldcheck",
		Doc: "Embedded types should be at the top of the field list of a struct, " +
			"and there must be an empty line separating embedded fields from regular fields. methods, and constructors",
		Run:      l.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	return a
}

type embeddedstructfieldcheck struct{}

func (l embeddedstructfieldcheck) run(pass *analysis.Pass) (any, error) {
	fset := pass.Fset

	insp, found := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !found {
		//nolint:nilnil // impossible case.
		return nil, nil
	}

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
		(*ast.CommentGroup)(nil),
	}

	var a internal.StructAnalyzer

	insp.Preorder(nodeFilter, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.StructType:
			if diag, ok := a.Analyze(); ok {
				pass.Report(diag)
			}

			a = internal.NewStructAnalyzer(fset, node)
		case *ast.CommentGroup:
			if a.IsAnalyzingStruct() && node.End() <= a.GetEndPos() {
				a.CheckCommentGroup(node)
			}
		}
	},
	)

	if diag, ok := a.Analyze(); ok {
		pass.Report(diag)
	}

	//nolint:nilnil //any, error
	return nil, nil
}
