package ebuf

import "github.com/reusee/rope"

type OpInsert struct {
	Pos, Len int
}

type OpDelete struct {
	Pos, Len int
}

func (b *Buffer) Insert(pos int, bs []byte) {
	r := b.Events[b.Current].(*rope.Rope).Insert(pos, bs)
	b.Events = b.Events[:b.Current+1]
	b.Events = append(b.Events, OpInsert{pos, len(bs)})
	b.Events = append(b.Events, r)
	b.Current += 2
}

func (b *Buffer) Delete(pos, length int) {
	r := b.Events[b.Current].(*rope.Rope).Delete(pos, length)
	b.Events = b.Events[:b.Current+1]
	b.Events = append(b.Events, OpDelete{pos, length})
	b.Events = append(b.Events, r)
	b.Current += 2
}
