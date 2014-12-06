package ebuf

import "container/list"

type LineBreaks struct {
	*Cursors
}

func NewLineBreaks() *LineBreaks {
	return &LineBreaks{
		Cursors: NewCursors(),
	}
}

func (l *LineBreaks) Operate(op Op, cur *list.Element) {
	bs := cur.Value.(*State).Rope.Sub(op.Pos, op.Len)
	switch op.Type {
	case Insert:
		l.Cursors.Operate(op, cur)
		for i, b := range bs {
			if b != '\n' {
				continue
			}
			pos := op.Pos + i
			l.Cursors.Add(&pos)
		}
	case Delete:
		for i, b := range bs {
			if b != '\n' {
				continue
			}
			pos := op.Pos + i
			l.Cursors.DelPos(pos)
		}
		l.Cursors.Operate(op, cur)
	}
}
