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
		op.Type = Delete
		for _, watcher := range b.Watchers {
			watcher.Delete(op)
		}
	case Delete:
		op.Type = Insert
		for _, watcher := range b.Watchers {
			watcher.Insert(op)
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
		for _, watcher := range b.Watchers {
			watcher.Insert(op)
		}
	case Delete:
		for _, watcher := range b.Watchers {
			watcher.Delete(op)
		}
	}
	if b.States[b.Current].Skip {
		goto redo
	}
}
