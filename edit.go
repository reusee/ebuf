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
		Skip: b.skipping,
	})
	b.Current += 1
	// update cursors
	for _, cursor := range b.Cursors {
		if *cursor >= pos {
			*cursor += len(bs)
		}
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
		Skip: b.skipping,
	})
	b.Current += 1
	// update cursors
	for _, cursor := range b.Cursors {
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
