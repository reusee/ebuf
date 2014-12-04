package ebuf

// Undo undos one edit operation
func (b *Buffer) Undo() {
	if b.Current == 0 {
		return
	}
undo:
	op := b.States[b.Current].LastOp
	switch op.Type {
	case Insert:
		for cursor := range b.Cursors {
			if *cursor > op.Pos && *cursor < op.Pos+op.Len {
				*cursor = op.Pos
			} else if *cursor >= op.Pos+op.Len {
				*cursor -= op.Len
			}
		}
	case Delete:
		for cursor := range b.Cursors {
			if *cursor >= op.Pos {
				*cursor += op.Len
			}
		}
	}
	b.Current--
	if b.States[b.Current].Skip {
		goto undo
	}
}

// Redo redos one edit operation
func (b *Buffer) Redo() {
	if b.Current+1 == len(b.States) {
		return
	}
redo:
	b.Current++
	op := b.States[b.Current].LastOp
	switch op.Type {
	case Insert:
		for cursor := range b.Cursors {
			if *cursor >= op.Pos {
				*cursor += op.Len
			}
		}
	case Delete:
		for cursor := range b.Cursors {
			if *cursor > op.Pos && *cursor < op.Pos+op.Len {
				*cursor = op.Pos
			} else if *cursor >= op.Pos+op.Len {
				*cursor -= op.Len
			}
		}
	}
	if b.States[b.Current].Skip {
		goto redo
	}
}
