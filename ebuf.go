package ebuf

import "github.com/reusee/rope"

type Buffer struct {
	Events  []Event
	Current int
}

type Event interface {
}

func New(bs []byte) *Buffer {
	return &Buffer{
		Events: []Event{
			rope.NewFromBytes(bs),
		},
	}
}
