package ebuf

type Cursor struct {
	Pos map[int]int
}

func (b *Buffer) AddCursor(pos int) *Cursor {
	cursor := &Cursor{
		Pos: make(map[int]int),
	}
	if pos < 0 {
		pos = 0
	}
	if l := b.States[b.Current].Rope.Len(); pos > l {
		pos = l
	}
	cursor.Pos[b.Current] = pos
	b.Cursors = append(b.Cursors, cursor)
	return cursor
}
