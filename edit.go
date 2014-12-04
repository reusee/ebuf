package ebuf

const (
	Insert = iota
	Delete
)

type Op struct {
	Type, Pos, Len int
}

func (b *Buffer) Insert(pos int, bs []byte) {
	b.InsertWithTempCursors(pos, bs, nil)
}

func (b *Buffer) InsertWithTempCursors(pos int, bs []byte, tempCursors []*int) {
	r := b.States[b.Current].Rope.Insert(pos, bs)
	b.dropStates()
	b.States = append(b.States, State{
		Rope: r,
		LastOp: Op{
			Type: Insert,
			Pos:  pos,
			Len:  len(bs),
		},
		Skip: b.skipping,
	})
	b.Current++
	// update cursors
	for cursor := range b.Cursors {
		if *cursor >= pos {
			*cursor += len(bs)
		}
	}
	for _, cursor := range tempCursors {
		if *cursor >= pos {
			*cursor += len(bs)
		}
	}
}

func (b *Buffer) Delete(pos, length int) {
	b.DeleteWithTempCursors(pos, length, nil)
}

func (b *Buffer) DeleteWithTempCursors(pos, length int, tempCursors []*int) {
	r := b.States[b.Current].Rope.Delete(pos, length)
	b.dropStates()
	b.States = append(b.States, State{
		Rope: r,
		LastOp: Op{
			Type: Delete,
			Pos:  pos,
			Len:  length,
		},
		Skip: b.skipping,
	})
	b.Current++
	// update cursors
	for cursor := range b.Cursors {
		if *cursor > pos && *cursor < pos+length {
			*cursor = pos
		} else if *cursor >= pos+length {
			*cursor -= length
		}
	}
	for _, cursor := range tempCursors {
		if *cursor > pos && *cursor < pos+length {
			*cursor = pos
		} else if *cursor >= pos+length {
			*cursor -= length
		}
	}
}

func (b *Buffer) dropStates() {
	b.States = b.States[:b.Current+1]
}
