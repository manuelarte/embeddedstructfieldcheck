package internal

import (
	"go/ast"
)

func IsFieldEmbedded(field *ast.Field) bool {
	return len(field.Names) == 0
}

func IsFieldSyncMutex(field *ast.Field) bool {
	if p, ok := field.Type.(*ast.StarExpr); ok {
		if se, isSelectorExpr := p.X.(*ast.SelectorExpr); isSelectorExpr {
			return selectorExprIsSyncMutex(se)
		}

		return false
	}

	if se, ok := field.Type.(*ast.SelectorExpr); ok {
		return selectorExprIsSyncMutex(se)
	}

	return false
}

func selectorExprIsSyncMutex(se *ast.SelectorExpr) bool {
	if packageName, isIdent := se.X.(*ast.Ident); isIdent && packageName.Name == "sync" {
		return se.Sel.Name == "Mutex"
	}

	return false
}
