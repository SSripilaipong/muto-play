package main

import "syscall/js"

func main() {
	js.Global().Set("interpret", js.FuncOf(Interpret))
	<-make(chan struct{})
}
