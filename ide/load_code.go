package main

import (
	"syscall/js"

	fileParser "github.com/SSripilaipong/muto/parser/file"
)

func LoadCode(this js.Value, args []js.Value) any {
	defer printer.Clear()
	input := args[0].String()

	err := loadCode(input)
	if err != nil {
		printer.Clear()
		return js.ValueOf(err.Error())
	}
	return js.ValueOf(printer.ReadPrintBuffer())
}

func loadCode(code string) error {
	pkg, err := fileParser.ParsePackageFromString(code).Return()
	if err != nil {
		return err
	}
	files := pkg.Files()
	if len(files) == 0 {
		prog = newProgram(nil)
		return nil
	}
	prog = newProgram(files[0].Statements())
	return nil
}
