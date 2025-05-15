package internal

import (
	"go/ast"
)

func IsFieldEmbedded(field *ast.Field) bool {
	return len(field.Names) == 0
}

func GetFieldSyncMutex(field *ast.Field) (*ast.SelectorExpr, bool) {
	if p, ok := field.Type.(*ast.StarExpr); ok {
		if se, isSelectorExpr := p.X.(*ast.SelectorExpr); isSelectorExpr {
			return se, selectorExprIsSyncMutex(se)
		}

		return nil, false
	}

	if se, ok := field.Type.(*ast.SelectorExpr); ok {
		return se, selectorExprIsSyncMutex(se)
	}

	return nil, false
}

func selectorExprIsSyncMutex(se *ast.SelectorExpr) bool {
	if packageName, isIdent := se.X.(*ast.Ident); isIdent && packageName.Name == "sync" {
		return se.Sel.Name == "Mutex" || se.Sel.Name == "RWMutex"
	}

	return false
}
