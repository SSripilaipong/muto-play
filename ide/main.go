package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("loadCode", js.FuncOf(LoadCode))
	js.Global().Set("execute", js.FuncOf(Execute))
	<-make(chan struct{})
}
