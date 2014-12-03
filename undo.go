package ebuf

func (b *Buffer) Undo() {
	if b.Current == 0 {
		return
	}
	op := b.States[b.Current].LastOp
	switch op.Type {
	case Insert:
		for _, cursor := range b.Cursors {
			if *cursor > op.Pos && *cursor < op.Pos+op.Len {
				*cursor = op.Pos
			} else if *cursor >= op.Pos+op.Len {
				*cursor -= op.Len
			}
		}
	case Delete:
		for _, cursor := range b.Cursors {
			if *cursor >= op.Pos {
				*cursor += op.Len
			}
		}
	}
	b.Current--
}

func (b *Buffer) Redo() {
	if b.Current+1 == len(b.States) {
		return
	}
	b.Current++
	op := b.States[b.Current].LastOp
	switch op.Type {
	case Insert:
		for _, cursor := range b.Cursors {
			if *cursor >= op.Pos {
				*cursor += op.Len
			}
		}
	case Delete:
		for _, cursor := range b.Cursors {
			if *cursor > op.Pos && *cursor < op.Pos+op.Len {
				*cursor = op.Pos
			} else if *cursor >= op.Pos+op.Len {
				*cursor -= op.Len
			}
		}
	}
}
