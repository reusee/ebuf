package ebuf

import (
	"math/rand"
	"testing"
)

func TestCursors(t *testing.T) {
	cursors := NewCursors()
	n := 51200
	var ints []int
	for i := 0; i < n; i++ {
		ints = append(ints, i)
	}
	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		ints[i], ints[j] = ints[j], ints[i]
	}
	for _, i := range ints {
		i := i
		cursors.Add(&i)
	}
	cur := cursors.Head.Next[0]
	for i := 0; i < n; i++ {
		if *cur.Value != i {
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
