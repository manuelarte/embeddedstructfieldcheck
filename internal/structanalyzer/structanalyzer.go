package structanalyzer

import (
	"github.com/manuelarte/embeddedcheck/internal/astutils"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/manuelarte/embeddedcheck/internal/diag"
)

type StructAnalyzer struct {
	StructType *ast.StructType
}

func New(st *ast.StructType) StructAnalyzer {
	return StructAnalyzer{
		StructType: st,
	}
}

func (s *StructAnalyzer) CheckNode(node ast.Node) {

}

//nolint:nestif // refactor later
func (s *StructAnalyzer) Analyze(fset *token.FileSet) (analysis.Diagnostic, bool) {
	if !s.IsAnalyzingStruct() {
		return analysis.Diagnostic{}, false
	}
	var firstEmbeddedField *ast.Field
	var lastEmbeddedField *ast.Field
	var firstNotEmbeddedField *ast.Field

	for _, field := range s.StructType.Fields.List {
		if astutils.IsFieldEmbedded(field) {
			if firstEmbeddedField == nil {
				firstEmbeddedField = field
			}
			if lastEmbeddedField == nil || lastEmbeddedField.Pos() < field.Pos() {
				lastEmbeddedField = field
			}
			if firstNotEmbeddedField != nil && firstNotEmbeddedField.Pos() < field.Pos() {
				return diag.NewEmbeddedTypeAfterNotEmbeddedTypeDiag(field), true
			}
		} else {
			if firstNotEmbeddedField == nil {
				firstNotEmbeddedField = field
			}
		}
	}

	// check for missing space
	if lastEmbeddedField != nil && firstNotEmbeddedField != nil {
		line := fset.Position(lastEmbeddedField.End()).Line
		nextLine := fset.Position(firstNotEmbeddedField.Pos()).Line
		if nextLine != line+2 {
			return diag.NewMissingSpaceBetweenLastEmbeddedTypeAndFirstNotEmbeddedTypeDiag(lastEmbeddedField, firstNotEmbeddedField), true
		}
	}
	return analysis.Diagnostic{}, false
}

func (s *StructAnalyzer) IsAnalyzingStruct() bool {
	return s.StructType != nil
}

func (s *StructAnalyzer) GetEndPos() token.Pos {
	return s.StructType.End()
}
