package ebuf

// Undo undos one edit operation
func (b *Buffer) Undo() {
	if b.Current == 0 {
		return
	}
undo:
	op := b.States[b.Current].LastOp
	if op.Type == Insert {
		op.Type = Delete
	} else {
		op.Type = Insert
	}
	for _, watcher := range b.Watchers {
		watcher.Operate(op)
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
	for _, watcher := range b.Watchers {
		watcher.Operate(op)
	}
	if b.States[b.Current].Skip {
		goto redo
	}
}
