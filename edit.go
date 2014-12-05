package ebuf

// Insert inserts bytes to specified position
func (b *Buffer) Insert(pos int, bs []byte) {
	b.InsertWithWatcher(pos, bs, nil)
}

// InsertWithWatcher inserts bytes to specified position with extra watcher
func (b *Buffer) InsertWithWatcher(pos int, bs []byte, watcher Watcher) {
	r := b.States[b.Current].Rope.Insert(pos, bs)
	b.dropStates()
	op := Op{
		Type:  Insert,
		Pos:   pos,
		Bytes: bs,
	}
	b.States = append(b.States, State{
		Rope:   r,
		LastOp: op,
		Skip:   b.skipping,
	})
	b.Current++
	// watchers
	for _, watcher := range b.Watchers {
		watcher.Operate(op)
	}
	if watcher != nil {
		watcher.Operate(op)
	}
}

// Delete deletes specified lengthed bytes from specified position
func (b *Buffer) Delete(pos, length int) {
	b.DeleteWithWatcher(pos, length, nil)
}

// DeleteWithWatcher deletes specified lengthed bytes from specified position with extra watcher
func (b *Buffer) DeleteWithWatcher(pos, length int, watcher Watcher) {
	bs := b.States[b.Current].Rope.Sub(pos, length)
	r := b.States[b.Current].Rope.Delete(pos, length)
	b.dropStates()
	op := Op{
		Type:  Delete,
		Pos:   pos,
		Bytes: bs,
	}
	b.States = append(b.States, State{
		Rope:   r,
		LastOp: op,
		Skip:   b.skipping,
	})
	b.Current++
	// watchers
	for _, watcher := range b.Watchers {
		watcher.Operate(op)
	}
	if watcher != nil {
		watcher.Operate(op)
	}
}

func (b *Buffer) dropStates() {
	b.States = b.States[:b.Current+1]
}
