package structanalyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"

	"github.com/manuelarte/embeddedcheck/internal/astutils"
	"github.com/manuelarte/embeddedcheck/internal/diag"
)

type StructAnalyzer struct {
	fset *token.FileSet
	st   *ast.StructType

	cg map[int]*ast.CommentGroup
}

func New(fset *token.FileSet, st *ast.StructType) StructAnalyzer {
	return StructAnalyzer{
		fset: fset,
		st:   st,

		cg: make(map[int]*ast.CommentGroup),
	}
}

//nolint:nestif // refactor later
func (s *StructAnalyzer) Analyze() (analysis.Diagnostic, bool) {
	if !s.IsAnalyzingStruct() {
		return analysis.Diagnostic{}, false
	}
	var firstEmbeddedField *ast.Field
	var lastEmbeddedField *ast.Field
	var firstNotEmbeddedField *ast.Field

	for _, field := range s.st.Fields.List {
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

	// check for missing space (TODO: isn't it easy to remove as many lines as comments between last embededed type and first not embedded)
	if lastEmbeddedField != nil && firstNotEmbeddedField != nil {
		line := s.fset.Position(lastEmbeddedField.End()).Line

		nextLine := s.fset.Position(firstNotEmbeddedField.Pos()).Line
		if cg, ok := s.cg[nextLine-1]; ok {
			nextLine = s.fset.Position(cg.Pos()).Line
		}
		if nextLine != line+2 {
			return diag.NewMissingSpaceBetweenLastEmbeddedTypeAndFirstNotEmbeddedTypeDiag(lastEmbeddedField, firstNotEmbeddedField), true
		}
	}
	return analysis.Diagnostic{}, false
}

func (s *StructAnalyzer) CheckCommentGroup(node *ast.CommentGroup) {
	lineEnd := s.fset.Position(node.End()).Line
	s.cg[lineEnd] = node
}

func (s *StructAnalyzer) IsAnalyzingStruct() bool {
	return s.st != nil
}

func (s *StructAnalyzer) GetEndPos() token.Pos {
	return s.st.End()
}
