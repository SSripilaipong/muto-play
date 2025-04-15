package common

import "strings"

type BufferedPrinter struct {
	buffer []string
}

func NewBufferedPrinter() *BufferedPrinter {
	return &BufferedPrinter{}
}

func (b *BufferedPrinter) Pop() (string, bool) {
	if len(b.buffer) == 0 {
		return "", false
	}
	x := b.buffer[0]
	b.buffer = b.buffer[1:]
	return x, true
}

func (b *BufferedPrinter) Clear() {
	b.buffer = nil
}

func (b *BufferedPrinter) Print(x string) {
	b.buffer = append(b.buffer, x)
}

func (b *BufferedPrinter) ReadPrintBuffer() string {
	result := ""
	for {
		s, ok := b.Pop()
		if !ok {
			break
		}
		result += s + "\n"
	}
	return strings.TrimSpace(result)
}
