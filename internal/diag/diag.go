package diag

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func NewEmbeddedTypeAfterNotEmbeddedTypeDiag(embeddedField *ast.Field) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos:     embeddedField.Pos(),
		Message: "embedded types should be listed before non embedded types",
	}
}
