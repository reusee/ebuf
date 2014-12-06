package ebuf

import (
	"container/list"
	"math/rand"
	"strconv"
	"strings"
)

// AddCursor adds one cursor to buffer
func (b *Buffer) AddCursor(pos int) *int {
	if pos < 0 {
		pos = 0
	}
	if l := b.Current.Value.(*State).Rope.Len(); pos > l {
		pos = l
	}
	b.Cursors.Add(&pos)
	return &pos
}

// DelCursor deletes cursor from buffer
func (b *Buffer) DelCursor(cursor *int) {
	b.Cursors.DelCursor(cursor)
}

type Cursors struct {
	Head   *Cursor
	length int
}

type Cursor struct {
	Value *int
	Next  []*Cursor
}

func (c *Cursor) Pos() int {
	return *c.Value
}

const maxLevel = 22

func NewCursors() *Cursors {
	return &Cursors{
		Head: &Cursor{
			Next: make([]*Cursor, maxLevel),
		},
	}
}

func (l *Cursors) getPrevs(n int) []**Cursor {
	prevs := make([]**Cursor, maxLevel)
	cur := l.Head
	for level := maxLevel - 1; level >= 0; level-- {
		for cur.Next[level] != nil && cur.Next[level].Pos() < n {
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
	cursor := &Cursor{
		Value: c,
		Next:  make([]*Cursor, level),
	}
	// update pointers
	for i := 0; i < level; i++ {
		var next *Cursor
		if *prevs[i] != nil {
			next = *prevs[i]
		}
		cursor.Next[i] = next
		*prevs[i] = cursor
	}
	l.length++
}

func (l *Cursors) DelCursor(c *int) {
	// get prevs
	prevs := l.getPrevs(*c)
	// update pointers
	deleted := false
	for level := 0; level < maxLevel; level++ {
		if *prevs[level] == nil {
			continue
		}
		if (*prevs[level]).Value != c {
			continue
		}
		*prevs[level] = (*prevs[level]).Next[level]
		deleted = true
	}
	if deleted {
		l.length--
	}
}

func (l *Cursors) DelPos(pos int) {
	// get prevs
	prevs := l.getPrevs(pos)
	// update pointers
	deleted := false
	for level := 0; level < maxLevel; level++ {
		if *prevs[level] == nil {
			continue
		}
		if (*prevs[level]).Pos() != pos {
			continue
		}
		*prevs[level] = (*prevs[level]).Next[level]
		deleted = true
	}
	if deleted {
		l.length--
	}
}

func (l *Cursors) Iterate(fn func(*Cursor) bool) {
	cur := l.Head.Next[0]
	for cur != nil {
		if !fn(cur) {
			break
		}
		cur = cur.Next[0]
	}
}

func (l *Cursors) Len() int {
	return l.length
}

func (l *Cursors) String() string {
	var strs []string
	l.Iterate(func(cursor *Cursor) bool {
		strs = append(strs, strconv.Itoa(cursor.Pos()))
		return true
	})
	return strings.Join(strs, " ")
}

// Operate deals with an buffer operation
func (l *Cursors) Operate(op Op, _ *list.Element) {
	switch op.Type {
	case Insert:
		l.Iterate(func(cursor *Cursor) bool {
			if cursor.Pos() >= op.Pos {
				*cursor.Value += op.Len
			}
			return true
		})
	case Delete:
		l.Iterate(func(cursor *Cursor) bool {
			if cursor.Pos() > op.Pos && cursor.Pos() < op.Pos+op.Len {
				*cursor.Value = op.Pos
			} else if cursor.Pos() >= op.Pos+op.Len {
				*cursor.Value -= op.Len
			}
			return true
		})
	}
}
