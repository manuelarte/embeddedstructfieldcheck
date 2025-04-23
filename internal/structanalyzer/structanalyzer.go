package structanalyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"

	"github.com/manuelarte/embeddedcheck/internal/diag"
)

type StructAnalyzer struct {
}

func New() *StructAnalyzer {
	return &StructAnalyzer{}
}

//nolint:nestif // refactor later
func (s *StructAnalyzer) Analyze(tp *ast.TypeSpec) (analysis.Diagnostic, bool) {
	var firstEmbeddedField *ast.Field
	var firstNotEmbeddedField *ast.Field
	if structTypeExpr, ok := tp.Type.(*ast.StructType); ok {
		for _, field := range structTypeExpr.Fields.List {
			if len(field.Names) == 0 {
				if firstEmbeddedField == nil {
					firstEmbeddedField = field
				}
				if firstNotEmbeddedField != nil && field.Pos() > firstNotEmbeddedField.Pos() {
					return diag.NewEmbeddedTypeAfterNotEmbeddedTypeDiag(field), true
				}
			} else {
				if firstNotEmbeddedField == nil {
					firstNotEmbeddedField = field
				}
			}
		}
	}
	return analysis.Diagnostic{}, false
}
