package ebuf

import "github.com/reusee/rope"

type Buffer struct {
	States   []State
	skipping bool
	Current  int
	Cursors  []*int
}

type State struct {
	Rope   *rope.Rope
	LastOp Op
	Skip   bool
}

func New(bs []byte) *Buffer {
	return &Buffer{
		States: []State{
			State{
				Rope: rope.NewFromBytes(bs),
			},
		},
	}
}
