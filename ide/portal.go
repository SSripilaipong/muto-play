package main

import (
	"github.com/SSripilaipong/muto/common/optional"
	"github.com/SSripilaipong/muto/common/rods"
	"github.com/SSripilaipong/muto/core/base"
	"github.com/SSripilaipong/muto/core/portal"

	"github.com/SSripilaipong/muto-play/common"
)

func NewPortal(printer *common.BufferedPrinter) *portal.Portal {
	return portal.New(rods.NewMap(map[string]portal.Port{
		"stdout": NewStdOut(printer),
		"stdin":  NewStdIn(),
	}))
}

type StdOut struct {
	printer *common.BufferedPrinter
}

func NewStdOut(printer *common.BufferedPrinter) StdOut {
	return StdOut{printer: printer}
}

func (s StdOut) Call(x base.Node) optional.Of[base.Node] {
	if !base.IsStringNode(x) {
		return optional.Empty[base.Node]()
	}
	s.printer.Print(base.UnsafeNodeToString(x).Value())
	return optional.Value[base.Node](base.Null())
}

type StdIn struct{}

func NewStdIn() StdIn {
	return StdIn{}
}

func (s StdIn) Call(x base.Node) optional.Of[base.Node] {
	if base.IsClassNode(x) && base.UnsafeNodeToClass(x).Name() == "$" {
		return optional.Value[base.Node](base.NewErrorWithMessage("console input is not supported"))
	}
	return optional.Empty[base.Node]()
}
