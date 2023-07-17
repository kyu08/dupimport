package gopher_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/kyu08/gopher"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, gopher.Analyzer, "a")
}
