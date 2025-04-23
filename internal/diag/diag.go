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

func NewMissingSpaceBetweenLastEmbeddedTypeAndFirstNotEmbeddedTypeDiag(lastEmbeddedField *ast.Field) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos:     lastEmbeddedField.Pos(),
		Message: "there must be an empty line separating embedded fields from regular fields",
	}
}
