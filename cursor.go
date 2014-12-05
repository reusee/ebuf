package ebuf

// AddCursor adds one cursor to buffer
func (b *Buffer) AddCursor(pos int) *int {
	if pos < 0 {
		pos = 0
	}
	if l := b.States[b.Current].Rope.Len(); pos > l {
		pos = l
	}
	b.Cursors[&pos] = struct{}{}
	return &pos
}

// DelCursor deletes cursor from buffer
func (b *Buffer) DelCursor(cursor *int) {
	delete(b.Cursors, cursor)
}

type CursorSet map[*int]struct{}

func NewCursorSet(cursors ...*int) CursorSet {
	set := CursorSet(make(map[*int]struct{}))
	for _, cursor := range cursors {
		set[cursor] = struct{}{}
	}
	return set
}

func (c CursorSet) Operate(op Op) {
	opLen := len(op.Bytes)
	switch op.Type {
	case Insert:
		for cursor := range c {
			if *cursor >= op.Pos {
				*cursor += opLen
			}
		}
	case Delete:
		for cursor := range c {
			if *cursor > op.Pos && *cursor < op.Pos+opLen {
				*cursor = op.Pos
			} else if *cursor >= op.Pos+opLen {
				*cursor -= opLen
			}
		}
	}
}
