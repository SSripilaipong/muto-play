package main

import (
	replProgram "github.com/SSripilaipong/muto/builder/repl"

	"syscall/js"
)

func Interpret(this js.Value, args []js.Value) interface{} {
	input := args[0].String()

	result, _ := interpret(input)

	return js.ValueOf(result)
}

var reader = newBufferedReader()
var printer = newBufferedPrinter()

var prog = replProgram.New(reader, printer)

func interpret(code string) (string, bool) {
	reader.Set(code)
	cmd := prog.Read()
	if cmd.IsEmpty() {
		return printer.Pop()
	}

	exit := prog.Execute(cmd.Value())
	if exit.IsNotEmpty() {
		return "exit doesn't support here", true
	}
	return printer.Pop()
}
