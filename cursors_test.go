package ebuf

import (
	"math/rand"
	"testing"
)

func TestCursorsAdd(t *testing.T) {
	cursors := NewCursors()
	n := 10240
	// gen random sequence
	var ints []int
	for i := 0; i < n; i++ {
		ints = append(ints, i)
	}
	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		ints[i], ints[j] = ints[j], ints[i]
	}
	// add
	for _, i := range ints {
		i := i
		cursors.Add(&i)
	}
	// validate
	cur := cursors.Head.Next[0]
	for i := 0; i < n; i++ {
		if *cur.Value != i {
			t.Fatal("cursor pos")
		}
		cur = cur.Next[0]
	}
}

func TestCursorsDelCursor(t *testing.T) {
	cursors := NewCursors()
	n := 10240
	// add
	var cs []*int
	for i := 0; i < n; i++ {
		i := i
		cursors.Add(&i)
		cs = append(cs, &i)
	}
	// delete even numbers
	for _, c := range cs {
		if *c%2 == 0 {
			cursors.DelCursor(c)
		}
	}
	// validate
	cur := cursors.Head.Next[0]
	for i := 0; i < n/2; i++ {
		if *cur.Value != i*2+1 {
			t.Fatal("cursor pos")
		}
		cur = cur.Next[0]
	}
}

func TestCursorDelPos(t *testing.T) {
	cursors := NewCursors()
	n := 10240
	// add
	for i := 0; i < n; i++ {
		i := i
		cursors.Add(&i)
	}
	// delete even numbers
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			cursors.DelPos(i)
		}
	}
	// validate
	cur := cursors.Head.Next[0]
	for i := 0; i < n/2; i++ {
		if *cur.Value != i*2+1 {
			t.Fatal("cursor pos")
		}
		cur = cur.Next[0]
	}
}

func (l *Cursors) dump() {
	for level := maxLevel - 1; level >= 0; level-- {
		cur := l.Head.Next[level]
		pt("level %d: ", level)
		for cur != nil {
			pt("%d ", *cur.Value)
			cur = cur.Next[level]
		}
		pt("\n")
	}
}
