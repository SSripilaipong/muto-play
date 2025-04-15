package main

import (
	"syscall/js"
)

func Interpret(this js.Value, args []js.Value) any {
	input := args[0].String()

	result, _ := interpret(input)

	return js.ValueOf(result)
}

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
