package ebuf

func (b *Buffer) Undo() {
	if b.Current == 0 {
		return
	}
	b.Current--
	//TODO update cursors
}

func (b *Buffer) Redo() {
	if b.Current+1 == len(b.States) {
		return
	}
	b.Current++
	//TODO update cursors
}
