package circular

import "errors"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
	head  int
	items int
	size  int
	data  []byte
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		head:  0,
		items: 0,
		size:  size,
		data:  make([]byte, size, size),
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.items < 1 {
		return 0, errors.New("No more data available")
	}

	c := b.data[b.head]
	b.items--
	if b.items > 0 {
		b.head = (b.head + 1) % b.size
	}

	return c, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.items == b.size {
		return errors.New("No room available")
	}

	b.data[(b.head+b.items)%b.size] = c
	b.items++

	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.WriteByte(c) != nil {
		b.data[b.head] = c
		b.head = (b.head + 1) % b.size
	}
}

func (b *Buffer) Reset() {
	b.head = 0
	b.items = 0
}
