package astutils

import (
	"go/ast"
)

func IsFieldEmbedded(field *ast.Field) bool {
	return len(field.Names) == 0
}

func HasEmbeddedTypes(n *ast.StructType) bool {
	return true
	/*for _, field := range n.Fields.Fields.List {
		if len(field.Names) == 0
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
	}*/
}
