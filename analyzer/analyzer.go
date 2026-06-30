package analyzer

import (
	"go/ast"
	"slices"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/manuelarte/embeddedstructfieldcheck/internal"
)

const (
	EmptyLineCheck   = "empty-line"
	ForbidMutexCheck = "forbid-mutex"
)

func NewAnalyzer() *analysis.Analyzer {
	var (
		emptyLine   bool
		forbidMutex bool
	)

	a := &analysis.Analyzer{
		Name: "embeddedstructfieldcheck",
		Doc: "Embedded types should be at the top of the field list of a struct, " +
			"and there must be an empty line separating embedded fields from regular fields.",
		URL: "https://github.com/manuelarte/embeddedstructfieldcheck",
		Run: func(pass *analysis.Pass) (any, error) {
			run(pass, emptyLine, forbidMutex)

			//nolint:nilnil // impossible case.
			return nil, nil
		},
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}

	a.Flags.BoolVar(&emptyLine, EmptyLineCheck, true,
		"Checks that there is an empty space between the embedded fields and regular fields.")
	a.Flags.BoolVar(&forbidMutex, ForbidMutexCheck, false,
		"Checks that sync.Mutex and sync.RWMutex are not used as an embedded fields.")

	return a
}

func run(pass *analysis.Pass, emptyLine, forbidMutex bool) {
	insp, found := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	if !found {
		return
	}

	nodeFilter := []ast.Node{
		new(ast.File),
		new(ast.StructType),
	}

	var fileComments []*ast.CommentGroup

	insp.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:
			fileComments = n.Comments
		case *ast.StructType:
			fieldComments := make([]*ast.CommentGroup, 0, len(getFields(n.Fields)))
			for _, field := range getFields(n.Fields) {
				if field.Doc == nil {
					continue
				}

				fieldComments = append(fieldComments, field.Doc)
			}

			insideStructNotAttachedComments := make([]*ast.CommentGroup, 0)

			for _, cg := range fileComments {
				if cg.Pos() > n.Pos() && cg.End() < n.End() && !slices.Contains(fieldComments, cg) {
					insideStructNotAttachedComments = append(insideStructNotAttachedComments, cg)
				}
			}

			internal.Analyze(pass, n, emptyLine, forbidMutex, insideStructNotAttachedComments)
		}
	})
}

func getFields(fl *ast.FieldList) []*ast.Field {
	if fl == nil {
		return nil
	}

	return fl.List
}
