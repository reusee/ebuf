package ebuf

// Insert inserts bytes to specified position
func (b *Buffer) Insert(pos int, bs []byte) {
	b.InsertWithWatcher(pos, bs, nil)
}

// InsertWithWatcher inserts bytes to specified position with extra watcher
func (b *Buffer) InsertWithWatcher(pos int, bs []byte, watcher Watcher) {
	r := b.Current.Value.(*State).Rope.Insert(pos, bs)
	b.dropStates()
	op := Op{
		Type: Insert,
		Pos:  pos,
		Len:  len(bs),
	}
	b.Current = b.States.PushBack(&State{
		Rope:   r,
		LastOp: op,
		Skip:   b.skipping,
	})
	// watchers
	for _, watcher := range b.Watchers {
		watcher.Operate(op, r)
	}
	if watcher != nil {
		watcher.Operate(op, r)
	}
}

// Delete deletes specified lengthed bytes from specified position
func (b *Buffer) Delete(pos, length int) {
	b.DeleteWithWatcher(pos, length, nil)
}

// DeleteWithWatcher deletes specified lengthed bytes from specified position with extra watcher
func (b *Buffer) DeleteWithWatcher(pos, length int, watcher Watcher) {
	rp := b.Current.Value.(*State).Rope
	bs := rp.Sub(pos, length)
	r := rp.Delete(pos, length)
	b.dropStates()
	op := Op{
		Type: Delete,
		Pos:  pos,
		Len:  len(bs),
	}
	b.Current = b.States.PushBack(&State{
		Rope:   r,
		LastOp: op,
		Skip:   b.skipping,
	})
	// watchers
	for _, watcher := range b.Watchers {
		watcher.Operate(op, r)
	}
	if watcher != nil {
		watcher.Operate(op, r)
	}
}

func (b *Buffer) dropStates() {
	for b.States.Back() != b.Current {
		b.States.Remove(b.States.Back())
	}
}
