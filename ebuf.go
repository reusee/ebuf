package ebuf

import "github.com/reusee/rope"

// Buffer represents a editing buffer
type Buffer struct {
	States   []State
	skipping bool
	Current  int
	Cursors  map[*int]struct{}
}

// State represents a editing state
type State struct {
	Rope   *rope.Rope
	LastOp Op
	Skip   bool
}

// New creates a new buffer with initial bytes
func New(bs []byte) *Buffer {
	return &Buffer{
		States: []State{
			State{
				Rope: rope.NewFromBytes(bs),
			},
		},
		Cursors: make(map[*int]struct{}),
	}
}

// CurrentBytes get current bytes of buffer
func (b *Buffer) CurrentBytes() []byte {
	return b.States[b.Current].Rope.Bytes()
}
