package main

import (
	replProgram "github.com/SSripilaipong/muto/builder/repl"

	"github.com/SSripilaipong/muto-play/common"
)

var reader = common.NewBufferedReader()
var printer = common.NewBufferedPrinter()

var prog = replProgram.New(reader, printer)
