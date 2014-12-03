package ebuf

import "github.com/reusee/rope"

func (b *Buffer) Undo() {
	if b.Current == 0 {
		return
	}
	_, ok := b.Events[b.Current-1].(*rope.Rope)
	for !ok {
		b.Current--
		_, ok = b.Events[b.Current-1].(*rope.Rope)
	}
	b.Current--
}

func (b *Buffer) Redo() {
	if b.Current+1 == len(b.Events) {
		return
	}
	_, ok := b.Events[b.Current+1].(*rope.Rope)
	for !ok {
		b.Current++
		_, ok = b.Events[b.Current+1].(*rope.Rope)
	}
	b.Current++
}
