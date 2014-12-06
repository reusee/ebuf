package ebuf

import "math/rand"

type Cursors struct {
	Head *_Cursor
}

type _Cursor struct {
	Value *int
	Next  []*_Cursor
}

const maxLevel = 22

func NewCursors() *Cursors {
	return &Cursors{
		Head: &_Cursor{
			Next: make([]*_Cursor, maxLevel),
		},
	}
}

func (l *Cursors) getPrevs(n int) []**_Cursor {
	prevs := make([]**_Cursor, maxLevel)
	cur := l.Head
	for level := maxLevel - 1; level >= 0; level-- {
		for cur.Next[level] != nil && *cur.Next[level].Value < n {
			cur = cur.Next[level]
		}
		prevs[level] = &cur.Next[level]
	}
	return prevs
}

func (l *Cursors) Add(c *int) {
	// get prevs
	prevs := l.getPrevs(*c)
	// get new cursor level
	level := 1
	for i := 0; i < maxLevel-1; i++ {
		if rand.Intn(2) == 0 {
			level++
		} else {
			break
		}
	}
	// create new cursor
	cursor := &_Cursor{
		Value: c,
		Next:  make([]*_Cursor, level),
	}
	// update pointers
	for i := 0; i < level; i++ {
		var next *_Cursor
		if *prevs[i] != nil {
			next = *prevs[i]
		}
		cursor.Next[i] = next
		*prevs[i] = cursor
	}
}

func (l *Cursors) DelCursor(c *int) {
	// get prevs
	prevs := l.getPrevs(*c)
	// update pointers
	for level := 0; level < maxLevel; level++ {
		if *prevs[level] == nil {
			continue
		}
		if (*prevs[level]).Value != c {
			continue
		}
		*prevs[level] = (*prevs[level]).Next[level]
	}
}

func (l *Cursors) DelPos(pos int) {
	// get prevs
	prevs := l.getPrevs(pos)
	// update pointers
	for level := 0; level < maxLevel; level++ {
		if *prevs[level] == nil {
			continue
		}
		if *((*prevs[level]).Value) != pos {
			continue
		}
		*prevs[level] = (*prevs[level]).Next[level]
	}
}
