package main

import (
	replProgram "github.com/SSripilaipong/muto/builder/repl/core/program"
	"github.com/SSripilaipong/muto/core/mutation"
	"github.com/SSripilaipong/muto/core/mutation/normal/builtin"
	mutoProgram "github.com/SSripilaipong/muto/program"
	"github.com/SSripilaipong/muto/syntaxtree/base"

	"github.com/SSripilaipong/muto-play/common"
)

var reader = newCliReaderWrapper(common.NewBufferedReader())
var printer = common.NewBufferedPrinter()
var builtins = builtin.NewMutators(reader, printer)
var prog = replProgram.New(mutoProgram.New(mutation.NewFromStatements(nil, builtins)), printer)

func newProgram(st []base.Statement) replProgram.Wrapper {
	return replProgram.New(mutoProgram.New(mutation.NewFromStatements(st, builtins)), printer)
}
