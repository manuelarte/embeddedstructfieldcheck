package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/manuelarte/embeddedcheck/analyzer"
)

func main() {
	singlechecker.Main(analyzer.NewAnalyzer())
}
