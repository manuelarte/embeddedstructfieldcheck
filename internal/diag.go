package internal

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func NewEmbeddedTypeAfterRegularTypeDiag(embeddedField *ast.Field) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos:     embeddedField.Pos(),
		Message: "embedded fields should be listed before regular fields",
	}
}

func NewMissingSpaceBetweenLastEmbeddedTypeAndFirstRegularTypeDiag(
	lastEmbeddedField *ast.Field,
) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos:     lastEmbeddedField.Pos(),
		Message: "there must be an empty line separating embedded fields from regular fields",
		SuggestedFixes: []analysis.SuggestedFix{
			{
				Message: "adding extra line separating embedded fields from regular fields",
				TextEdits: []analysis.TextEdit{
					{
						Pos:     lastEmbeddedField.End(),
						NewText: []byte("\n\n"),
					},
				},
			},
		},
	}
}
