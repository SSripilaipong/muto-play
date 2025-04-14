package main

type bufferedReader struct {
	buffer string
}

func newBufferedReader() *bufferedReader {
	return &bufferedReader{}
}

func (b *bufferedReader) Set(s string) {
	b.buffer = s
}

func (b *bufferedReader) ReadLine() (string, error) {
	return b.buffer, nil
}
