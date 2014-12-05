package ebuf

// Editing Operations
const (
	Insert = true
	Delete = false
)

// Op represents editing operation type
type Op struct {
	Type     bool
	Pos, Len int
}

// Insert inserts bytes to specified position
func (b *Buffer) Insert(pos int, bs []byte) {
	b.InsertWithTempWatcher(pos, bs, nil)
}

// InsertWithTempWatcher inserts bytes to specified position, keeping a slice of cursors valid
func (b *Buffer) InsertWithTempWatcher(pos int, bs []byte, watcher Watcher) {
	r := b.States[b.Current].Rope.Insert(pos, bs)
	b.dropStates()
	op := Op{
		Type: Insert,
		Pos:  pos,
		Len:  len(bs),
	}
	b.States = append(b.States, State{
		Rope:   r,
		LastOp: op,
		Skip:   b.skipping,
	})
	b.Current++
	// watchers
	for _, watcher := range b.Watchers {
		watcher.Insert(op)
	}
	if watcher != nil {
		watcher.Insert(op)
	}
}

// Delete deletes specified lengthed bytes from specified position
func (b *Buffer) Delete(pos, length int) {
	b.DeleteWithTempWatcher(pos, length, nil)
}

// DeleteWithTempWatcher deletes specified lengthed bytes from specified position, keeping a slice of cursors valid
func (b *Buffer) DeleteWithTempWatcher(pos, length int, watcher Watcher) {
	r := b.States[b.Current].Rope.Delete(pos, length)
	b.dropStates()
	op := Op{
		Type: Delete,
		Pos:  pos,
		Len:  length,
	}
	b.States = append(b.States, State{
		Rope:   r,
		LastOp: op,
		Skip:   b.skipping,
	})
	b.Current++
	// watchers
	for _, watcher := range b.Watchers {
		watcher.Delete(op)
	}
	if watcher != nil {
		watcher.Delete(op)
	}
}

func (b *Buffer) dropStates() {
	b.States = b.States[:b.Current+1]
}
