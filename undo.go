package ebuf

// Undo undos one edit operation
func (b *Buffer) Undo() {
	if b.Current == b.States.Front() {
		return
	}
undo:
	op := b.Current.Value.(*State).LastOp
	if op.Type == Insert {
		op.Type = Delete
	} else {
		op.Type = Insert
	}
	b.Current = b.Current.Prev()
	for _, watcher := range b.Watchers {
		watcher.Operate(op, b.Current.Value.(*State).Rope)
	}
	if b.Current.Value.(*State).Skip {
		goto undo
	}
}

// Redo redos one edit operation
func (b *Buffer) Redo() {
	if b.Current == b.States.Back() {
		return
	}
redo:
	b.Current = b.Current.Next()
	op := b.Current.Value.(*State).LastOp
	for _, watcher := range b.Watchers {
		watcher.Operate(op, b.Current.Value.(*State).Rope)
	}
	if b.Current.Value.(*State).Skip {
		goto redo
	}
}
