package main

import (
	replProgram "github.com/SSripilaipong/muto/builder/repl/core/program"
	"github.com/SSripilaipong/muto/builtin/global"
	"github.com/SSripilaipong/muto/core/module"
	mutoProgram "github.com/SSripilaipong/muto/program"
	"github.com/SSripilaipong/muto/syntaxtree/base"

	"github.com/SSripilaipong/muto-play/common"
)

var reader = newCliReaderWrapper(common.NewBufferedReader())
var printer = common.NewBufferedPrinter()
var builtins = global.NewModule(reader, printer)
var prog = newProgram(nil)

func newProgram(st []base.Statement) replProgram.Wrapper {
	mod := module.BuildModuleFromStatements(st)
	mod.LoadGlobal(builtins)
	return replProgram.New(mutoProgram.New(mod), printer)
}
