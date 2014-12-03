package ebuf

const (
	Insert = iota
	Delete
)

type Op struct {
	Type, Pos, Len int
}

func (b *Buffer) Insert(pos int, bs []byte) {
	r := b.States[b.Current].Rope.Insert(pos, bs)
	b.dropStates()
	b.States = append(b.States, State{
		Rope: r,
		LastOp: Op{
			Type: Insert,
			Pos:  pos,
			Len:  len(bs),
		},
	})
	b.Current += 1
	// update cursors
	for _, cursor := range b.Cursors {
		cursorPos := cursor.Pos[b.Current-1]
		if cursorPos >= pos {
			cursorPos += len(bs)
		}
		cursor.Pos[b.Current] = cursorPos
	}
}

func (b *Buffer) Delete(pos, length int) {
	r := b.States[b.Current].Rope.Delete(pos, length)
	b.dropStates()
	b.States = append(b.States, State{
		Rope: r,
		LastOp: Op{
			Type: Delete,
			Pos:  pos,
			Len:  length,
		},
	})
	b.Current += 1
	// update cursors
	for _, cursor := range b.Cursors {
		cursorPos := cursor.Pos[b.Current-1]
		if cursorPos > pos && cursorPos < pos+length {
			cursorPos = pos
		} else if cursorPos >= pos+length {
			cursorPos -= length
		}
		cursor.Pos[b.Current] = cursorPos
	}
}

func (b *Buffer) dropStates() {
	b.States = b.States[:b.Current+1]
	for _, cursor := range b.Cursors {
		for i, _ := range cursor.Pos {
			if i > b.Current {
				delete(cursor.Pos, i)
			}
		}
	}
}
