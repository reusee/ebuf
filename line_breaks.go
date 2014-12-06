package ebuf

import "container/list"

type LineBreaks struct {
	*list.List
}

func NewLineBreaks() *LineBreaks {
	return &LineBreaks{
		List: list.New(),
	}
}

func (l *LineBreaks) Operate(op Op, cur *list.Element) {
	switch op.Type {
	case Insert:
		//pt("insert %s\n", cur.Value.(*State).Rope.Sub(op.Pos, op.Len))
	case Delete:
		//pt("delete %s\n", cur.Value.(*State).Rope.Sub(op.Pos, op.Len))
	}
}
