package internal

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

type StructAnalyzer struct {
	forbidMutex bool
}

func NewStructAnalyzer(forbidMutex bool) StructAnalyzer {
	return StructAnalyzer{forbidMutex: forbidMutex}
}

//nolint:nestif // not so complex logic
func (a StructAnalyzer) Analyze(pass *analysis.Pass, st *ast.StructType) {
	var firstEmbeddedField *ast.Field

	var lastEmbeddedField *ast.Field

	var firstNotEmbeddedField *ast.Field

	for _, field := range st.Fields.List {
		if IsFieldEmbedded(field) {
			if a.forbidMutex && IsFieldSyncMutex(field) {
				pass.Report(NewForbidMutexDiag(field))
			}

			if firstEmbeddedField == nil {
				firstEmbeddedField = field
			}

			if lastEmbeddedField == nil || lastEmbeddedField.Pos() < field.Pos() {
				lastEmbeddedField = field
			}

			if firstNotEmbeddedField != nil && firstNotEmbeddedField.Pos() < field.Pos() {
				pass.Report(NewMisplacedEmbeddedFieldDiag(field))
				return
			}
		} else if firstNotEmbeddedField == nil {
			firstNotEmbeddedField = field
		}
	}

	checkMissingSpace(pass, lastEmbeddedField, firstNotEmbeddedField)
}

func checkMissingSpace(pass *analysis.Pass, lastEmbeddedField, firstNotEmbeddedField *ast.Field) {
	if lastEmbeddedField != nil && firstNotEmbeddedField != nil {
		// check for missing space
		// TODO: isn't it easy to remove as many lines as comments between last embedded type and first not embedded
		line := pass.Fset.Position(lastEmbeddedField.End()).Line

		nextLine := pass.Fset.Position(firstNotEmbeddedField.Pos()).Line
		if firstNotEmbeddedField.Doc != nil {
			nextLine = pass.Fset.Position(firstNotEmbeddedField.Doc.Pos()).Line
		}

		if nextLine != line+2 {
			pass.Report(NewMissingSpaceDiag(lastEmbeddedField, firstNotEmbeddedField))
		}
	}
}
