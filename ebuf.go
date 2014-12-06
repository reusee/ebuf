package ebuf

import (
	"container/list"

	"github.com/reusee/rope"
)

// Buffer represents a editing buffer
type Buffer struct {
	States     *list.List
	skipping   bool
	Current    *list.Element
	Cursors    *Cursors
	LineBreaks *LineBreaks
	Watchers   map[string]Watcher
}

// State represents a editing state
type State struct {
	Rope   *rope.Rope
	LastOp Op
	Skip   bool
}

// Op represents editing operation type
type Op struct {
	Type bool
	Pos  int
	Len  int
}

// Editing Operations
const (
	Insert = true
	Delete = false
)

// New creates a new buffer with initial bytes
func New(bs []byte) *Buffer {
	cursors := NewCursors()
	lineBreaks := NewLineBreaks()
	buf := &Buffer{
		States:     list.New(),
		Cursors:    cursors,
		LineBreaks: lineBreaks,
		Watchers: map[string]Watcher{
			"cursors":     cursors,
			"line-breaks": lineBreaks,
		},
	}
	buf.Current = buf.States.PushBack(&State{
		Rope: rope.NewFromBytes(bs),
	})
	for _, watcher := range buf.Watchers {
		watcher.Operate(Op{
			Type: Insert,
			Pos:  0,
			Len:  len(bs),
		}, buf.Current)
	}
	return buf
}

// CurrentBytes get current bytes of buffer
func (b *Buffer) CurrentBytes() []byte {
	return b.Current.Value.(*State).Rope.Bytes()
}

func (b *Buffer) currentIndex() int {
	ret := 0
	cur := b.States.Front()
	for cur != b.Current {
		ret++
		cur = cur.Next()
	}
	return ret
}
