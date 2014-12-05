package ebuf

import "github.com/reusee/rope"

// Buffer represents a editing buffer
type Buffer struct {
	States   []State
	skipping bool
	Current  int
	Cursors  CursorSet
	Watchers []Watcher
}

// State represents a editing state
type State struct {
	Rope   *rope.Rope
	LastOp Op
	Skip   bool
}

// Op represents editing operation type
type Op struct {
	Type  bool
	Pos   int
	Bytes []byte
}

// Editing Operations
const (
	Insert = true
	Delete = false
)

// New creates a new buffer with initial bytes
func New(bs []byte) *Buffer {
	cursors := CursorSet(make(map[*int]struct{}))
	buf := &Buffer{
		States: []State{
			State{
				Rope: rope.NewFromBytes(bs),
			},
		},
		Cursors: cursors,
		Watchers: []Watcher{
			cursors,
		},
	}
	for _, watcher := range buf.Watchers {
		watcher.Operate(Op{
			Type:  Insert,
			Pos:   0,
			Bytes: bs,
		})
	}
	return buf
}

// CurrentBytes get current bytes of buffer
func (b *Buffer) CurrentBytes() []byte {
	return b.States[b.Current].Rope.Bytes()
}
