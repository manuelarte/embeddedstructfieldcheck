package astutils

import (
	"github.com/manuelarte/embeddedcheck/internal/diag"
	"go/ast"
)

func HasEmbeddedTypes(n *ast.TypeSpec) bool {
	for _, field := range n.Fields.List {
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
