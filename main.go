package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/manuelarte/embeddedstructfieldcheck/analyzer"
)

func main() {
	singlechecker.Main(analyzer.NewAnalyzer())
}
