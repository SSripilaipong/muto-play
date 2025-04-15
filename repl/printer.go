package main

type bufferedPrinter struct {
	buffer []string
}

func newBufferedPrinter() *bufferedPrinter {
	return &bufferedPrinter{}
}

func (b *bufferedPrinter) Pop() (string, bool) {
	if len(b.buffer) == 0 {
		return "", false
	}
	x := b.buffer[0]
	b.buffer = b.buffer[:len(b.buffer)-1]
	return x, true
}

func (b *bufferedPrinter) Print(x string) {
	b.buffer = append(b.buffer, x)
}
