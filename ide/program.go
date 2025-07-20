package main

import (
	replProgram "github.com/SSripilaipong/muto/builder/repl/core/program"
	"github.com/SSripilaipong/muto/builtin"
	"github.com/SSripilaipong/muto/builtin/global"
	"github.com/SSripilaipong/muto/common/slc"
	"github.com/SSripilaipong/muto/core/module"
	mutoProgram "github.com/SSripilaipong/muto/program"
	st "github.com/SSripilaipong/muto/syntaxtree"
	stBase "github.com/SSripilaipong/muto/syntaxtree/base"

	"github.com/SSripilaipong/muto-play/common"
)

var printer = common.NewBufferedPrinter()
var builtins = global.NewModule()
var prog = newProgram(nil)

func newProgram(s []stBase.Statement) replProgram.Wrapper {
	portal := NewPortal(printer)
	imports := slc.Map(st.ImportToJoinedPath)(st.FilterImportFromStatement(s))
	importMapping := builtin.NewBuiltinImportMapping(imports).
		Attach(builtins, portal)

	mod := module.BuildUserDefinedModule(st.NewModule(slc.Pure(st.NewFile(s)))).
		Attach(module.NewDependency(builtins, portal, importMapping))
	return replProgram.New(mutoProgram.New(mod), printer)
}
