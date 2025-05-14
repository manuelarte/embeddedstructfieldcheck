package internal

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func NewMisplacedEmbeddedFieldDiag(embeddedField *ast.Field) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos:     embeddedField.Pos(),
		Message: "embedded fields should be listed before regular fields",
	}
}

func NewMissingSpaceDiag(
	lastEmbeddedField *ast.Field,
	firstRegularField *ast.Field,
) analysis.Diagnostic {
	suggestedPos := firstRegularField.Pos()
	if firstRegularField.Doc != nil {
		suggestedPos = firstRegularField.Doc.Pos()
	}

	return analysis.Diagnostic{
		Pos:     lastEmbeddedField.Pos(),
		Message: "there must be an empty line separating embedded fields from regular fields",
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: "adding extra line separating embedded fields from regular fields",
				TextEdits: []analysis.TextEdit{
					{
						Pos:     suggestedPos,
						NewText: []byte("\n\n"),
					},
				},
			},
		},
	}
}

func NewForbidMutexDiag(mutexField *ast.Field) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos:     mutexField.Pos(),
		Message: "sync.Mytex should not be embedded",
	}
}
