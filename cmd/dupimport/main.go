package main

import (
	"github.com/dupimport"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(dupimport.Analyzer) }
