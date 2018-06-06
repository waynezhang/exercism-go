package circular

import "errors"

// Buffer doc here
type Buffer struct {
	buffer   []byte
	start    int
	occupied int
}

// ReadByte doc here
func (b *Buffer) ReadByte() (byte, error) {
	if b.occupied == 0 {
		return 0, errors.New("")
	}
	defer func() {
		b.start = (b.start + 1) % len(b.buffer)
		b.occupied--
	}()
	return b.buffer[b.start], nil
}

// WriteByte doc here
func (b *Buffer) WriteByte(c byte) error {
	if b.occupied == len(b.buffer) {
		return errors.New("")
	}

	defer func() {
		b.occupied++
	}()
	idx := (b.start + b.occupied) % len(b.buffer)
	b.buffer[idx] = c

	return nil
}

// Reset doc here
func (b *Buffer) Reset() {
	b.start = 0
	b.occupied = 0
}

// Overwrite doc here
func (b *Buffer) Overwrite(c byte) {
	if b.occupied == len(b.buffer) {
		b.buffer[b.start] = c
		b.start = (b.start + 1) % len(b.buffer)
	} else {
		b.buffer[b.start+1] = c
		b.occupied++
	}
}

// NewBuffer doc here
func NewBuffer(size int) *Buffer {
	return &Buffer{buffer: make([]byte, size)}
}
