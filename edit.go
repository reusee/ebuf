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
	b.InsertWithTempCursors(pos, bs, nil)
}

// InsertWithTempCursors inserts bytes to specified position, keeping a slice of cursors valid
func (b *Buffer) InsertWithTempCursors(pos int, bs []byte, tempCursors []*int) {
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
	// update cursors
	for _, cursor := range tempCursors {
		if *cursor >= pos {
			*cursor += len(bs)
		}
	}
	// watchers
	for _, watcher := range b.Watchers {
		watcher.Insert(op)
	}
}

// Delete deletes specified lengthed bytes from specified position
func (b *Buffer) Delete(pos, length int) {
	b.DeleteWithTempCursors(pos, length, nil)
}

// DeleteWithTempCursors deletes specified lengthed bytes from specified position, keeping a slice of cursors valid
func (b *Buffer) DeleteWithTempCursors(pos, length int, tempCursors []*int) {
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
	// update cursors
	for _, cursor := range tempCursors {
		if *cursor > pos && *cursor < pos+length {
			*cursor = pos
		} else if *cursor >= pos+length {
			*cursor -= length
		}
	}
	// watchers
	for _, watcher := range b.Watchers {
		watcher.Delete(op)
	}
}

func (b *Buffer) dropStates() {
	b.States = b.States[:b.Current+1]
}
