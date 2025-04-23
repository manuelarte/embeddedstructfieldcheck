package internal

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

type StructAnalyzer struct {
}

func NewStructAnalyzer() *StructAnalyzer {
	return &StructAnalyzer{}
}

func (s *StructAnalyzer) Analyze(*ast.TypeSpec) (analysis.Diagnostic, bool) {
	return analysis.Diagnostic{}, false
}
