package astutils

import (
	"go/ast"
)

func IsFieldEmbedded(field *ast.Field) bool {
	return len(field.Names) == 0
}
