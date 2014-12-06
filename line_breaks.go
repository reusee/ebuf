package ebuf

import "container/list"

type LineBreaks struct {
	CursorSet
}

func NewLineBreaks() *LineBreaks {
	return &LineBreaks{
		CursorSet: NewCursorSet(),
	}
}

func (l *LineBreaks) Operate(op Op, cur *list.Element) {
	bs := cur.Value.(*State).Rope.Sub(op.Pos, op.Len)
	switch op.Type {
	case Insert:
		l.CursorSet.Operate(op, cur)
		for i, b := range bs {
			if b != '\n' {
				continue
			}
			pos := op.Pos + i
			l.CursorSet[&pos] = struct{}{}
		}
	case Delete:
		for i, b := range bs {
			if b != '\n' {
				continue
			}
			pos := op.Pos + i
			for cursor := range l.CursorSet { //TODO slow
				if *cursor == pos {
					delete(l.CursorSet, cursor)
				}
			}
		}
		l.CursorSet.Operate(op, cur)
	}
}
