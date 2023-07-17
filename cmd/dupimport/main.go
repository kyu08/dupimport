package main

import (
	"github.com/kyu08/dupimport"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(dupimport.Analyzer) }
