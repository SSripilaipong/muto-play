package common

type BufferedReader struct {
	buffer string
}

func NewBufferedReader() *BufferedReader {
	return &BufferedReader{}
}

func (b *BufferedReader) Set(s string) {
	b.buffer = s
}

func (b *BufferedReader) ReadLine() (string, error) {
	return b.buffer, nil
}
