package main

import (
	"fmt"
	"syscall/js"

	mutationRuleBuilder "github.com/SSripilaipong/muto/core/mutation/rule/builder"
	"github.com/SSripilaipong/muto/core/pattern/parameter"
	resultParser "github.com/SSripilaipong/muto/parser/result"
)

func Execute(this js.Value, args []js.Value) any {
	defer printer.Clear()
	input := args[0].String()

	err := execute(input)
	if err != nil {
		return js.ValueOf(err.Error())
	}

	return js.ValueOf(printer.ReadPrintBuffer())
}

func execute(input string) error {
	syntaxNode, err := resultParser.ParseNakedObjectMultilines(input).Return()
	if err != nil {
		return err
	}
	builder := mutationRuleBuilder.NewLiteral(syntaxNode)
	node := builder.Build(parameter.New())
	if node.IsEmpty() {
		return fmt.Errorf("cannot build node")
	}
	if _, exit := prog.MutateNode(node.Value()).Return(); exit {
		return fmt.Errorf("exit doesn't support")
	}
	return nil
}
