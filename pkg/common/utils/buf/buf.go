package buf

import (
	"bytes"
)

type Buffer struct {
	*bytes.Buffer
}

func (b *Buffer) Close() error {
	b.Buffer.Reset()
	return nil
}

func New(b *bytes.Buffer) *Buffer {
	if b == nil {
		b = bytes.NewBuffer(nil)
	}
	return &Buffer{b}
}
