package main

import (
	"github.com/gcpug/zagane/zagane"
	"github.com/gostaticanalysis/ctxfield"
	"github.com/gostaticanalysis/exclude"
	"github.com/gostaticanalysis/forcetypeassert"
	"github.com/gostaticanalysis/loopdefer"
	"github.com/gostaticanalysis/nilerr"
	"github.com/gostaticanalysis/sqlrows/passes/sqlrows"
	"github.com/gostaticanalysis/testhelper"
	"github.com/gostaticanalysis/typeswitch"
	"github.com/gostaticanalysis/unitconst"
	"github.com/gostaticanalysis/wraperrfmt"
	"github.com/kyoh86/looppointer"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(exclude.All(append(
		zagane.Analyzers(),
		ctxfield.Analyzer,
		forcetypeassert.Analyzer,
		loopdefer.Analyzer,
		nilerr.Analyzer,
		sqlrows.Analyzer,
		testhelper.Analyzer,
		typeswitch.Analyzer,
		unitconst.Analyzer,
		wraperrfmt.Analyzer,
		looppointer.Analyzer,
	), exclude.GeneratedFile, exclude.FileWithPattern)...)
}
