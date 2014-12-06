package ebuf

import (
	"container/list"

	"github.com/reusee/rope"
)

type LineBreaks struct {
	*list.List
}

func NewLineBreaks() *LineBreaks {
	return &LineBreaks{
		List: list.New(),
	}
}

func (l *LineBreaks) Operate(op Op, rp *rope.Rope) {
	switch op.Type {
	case Insert:
	case Delete:
	}
}
