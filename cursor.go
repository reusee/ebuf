package ebuf

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
