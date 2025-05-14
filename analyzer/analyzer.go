package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/manuelarte/embeddedstructfieldcheck/internal"
)

const ForbidMutexName = "forbid-mutex"

func NewAnalyzer() *analysis.Analyzer {
	l := embeddedstructfieldcheck{}

	a := &analysis.Analyzer{
		Name: "embeddedstructfieldcheck",
		Doc: "Embedded types should be at the top of the field list of a struct, " +
			"and there must be an empty line separating embedded fields from regular fields. methods, and constructors",
		URL:      "https://github.com/manuelarte/embeddedstructfieldcheck",
		Run:      l.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	a.Flags.BoolVar(&l.forbidMutex, ForbidMutexName, false, "Checks that sync.Mutex is not used as an embedded field.")

	return a
}

type embeddedstructfieldcheck struct {
	forbidMutex bool
}

func (l *embeddedstructfieldcheck) run(pass *analysis.Pass) (any, error) {
	insp, found := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !found {
		//nolint:nilnil // impossible case.
		return nil, nil
	}

	nodeFilter := []ast.Node{
		(*ast.StructType)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		node, ok := n.(*ast.StructType)
		if !ok {
			return
		}

		a := internal.NewStructAnalyzer(l.forbidMutex)
		a.Analyze(pass, node)
	})

	//nolint:nilnil //any, error
	return nil, nil
}
