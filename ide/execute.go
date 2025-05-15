package main

import (
	"fmt"
	"syscall/js"

	resultParser "github.com/SSripilaipong/muto/parser/result"
	"github.com/SSripilaipong/muto/syntaxtree/result"
)

func Execute(this js.Value, args []js.Value) any {
	defer printer.Clear()
	input := args[0].String()

	err := execute(input)
	if err != nil {
		return js.ValueOf(map[string]any{
			"err": err.Error(),
		})
	}

	return js.ValueOf(map[string]any{
		"result": printer.ReadPrintBuffer(),
	})
}

func execute(input string) error {
	syntaxNode, err := resultParser.ParseNakedObjectMultilines(input).Return()
	if err != nil {
		return err
	}
	node := prog.BuildNode(result.UnsafeNodeToObject(syntaxNode))
	if node.IsEmpty() {
		return fmt.Errorf("cannot build node")
	}
	if _, exit := prog.MutateNode(node.Value()).Return(); exit {
		return fmt.Errorf("exit doesn't support")
	}
	return nil
}
