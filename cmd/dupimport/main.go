package main

import (
	"github.com/kyu08/gopher"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(gopher.Analyzer) }
