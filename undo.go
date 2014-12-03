package ebuf

func (b *Buffer) Undo() {
	if b.Current == 0 {
		return
	}
	op := b.States[b.Current].LastOp
	switch op.Type {
	case Insert: //TODO
	case Delete: //TODO
	}
	b.Current--
}

func (b *Buffer) Redo() {
	if b.Current+1 == len(b.States) {
		return
	}
	op := b.States[b.Current].LastOp
	switch op.Type {
	case Insert: //TODO
	case Delete: //TODO
	}
	b.Current++
}
