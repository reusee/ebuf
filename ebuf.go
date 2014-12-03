package ebuf

import "github.com/reusee/rope"

type Buffer struct {
	States  []State
	Current int
	Cursors []*int
}

type State struct {
	Rope   *rope.Rope
	LastOp Op
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
