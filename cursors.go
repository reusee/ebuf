package ebuf

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

func (l *Cursors) Add(c *int) {
	// get prevs
	prevs := make([]**_Cursor, maxLevel)
	cur := l.Head
	for level := maxLevel - 1; level >= 0; level-- {
		for cur.Next[level] != nil && *cur.Next[level].Value < *c {
			cur = cur.Next[level]
		}
		prevs[level] = &cur.Next[level]
	}
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
