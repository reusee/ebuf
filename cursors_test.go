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
		if cur.Pos() != i {
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
		if cur.Pos() != i*2+1 {
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
		if cur.Pos() != i*2+1 {
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
			pt("%d ", cur.Pos())
			cur = cur.Next[level]
		}
		pt("\n")
	}
}

func TestCursorWithInsert(t *testing.T) {
	b := New(nil)

	cursor := b.AddCursor(0)
	if b.Cursors.Len() != 1 {
		t.Fatal("cursors length")
	}
	b.Insert(0, []byte("foobar"))
	if *cursor != 6 {
		t.Fatal("cursor pos")
	}

	cursor = b.AddCursor(3)
	if b.Cursors.Len() != 2 {
		t.Fatal("cursors length")
	}
	b.Insert(6, []byte("baz"))
	if *cursor != 3 {
		t.Fatal("cursor pos")
	}
}

func TestCursorWithDelete(t *testing.T) {
	b := New([]byte("foobarbaz"))

	cursor := b.AddCursor(-1) // for coverage
	if *cursor != 0 {
		t.Fatal("cursor pos")
	}
	if b.Cursors.Len() != 1 {
		t.Fatal("cursors length")
	}
	b.Delete(0, 3)
	if *cursor != 0 {
		t.Fatal("cursor pos")
	}

	cursor = b.AddCursor(b.Current.Value.(*State).Rope.Len() + 42) // for coverage
	if *cursor != 6 {
		t.Fatal("cursor pos")
	}
	b.Delete(0, 3)
	if *cursor != 3 {
		t.Fatal("cursor pos")
	}

	cursor = b.AddCursor(2)
	b.Delete(0, 3)
	if *cursor != 0 {
		t.Fatal("cursor pos")
	}
}

func TestCursorWithUndoRedo(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte("foobarbaz"))
	cursor := b.AddCursor(9)
	b.Undo()
	if *cursor != 0 {
		t.Fatal("cursor pos")
	}
	b.Insert(0, []byte("foobar"))
	if *cursor != 6 {
		t.Fatal("cursor pos")
	}

	b = New([]byte("foobar"))
	b.Delete(3, 3)
	cursor = b.AddCursor(3)
	b.Undo()
	if *cursor != 6 {
		t.Fatal("cursor pos")
	}

	b = New([]byte("foobar"))
	b.Insert(3, []byte("baz"))
	cursor = b.AddCursor(5)
	b.Undo()
	if *cursor != 3 {
		t.Fatal("cursor pos")
	}

	b = New([]byte("foobarbaz"))
	b.Delete(3, 3)
	b.Undo()
	cursor = b.AddCursor(5)
	b.Redo()
	if *cursor != 3 {
		t.Fatal("cursor pos")
	}

	b = New([]byte("foobar"))
	b.Insert(3, []byte("baz"))
	b.Undo()
	cursor = b.AddCursor(3)
	b.Redo()
	if *cursor != 6 {
		t.Fatal("cursor pos")
	}

	b = New([]byte("foobarbaz"))
	b.Delete(3, 3)
	b.Undo()
	cursor = b.AddCursor(9)
	b.Redo()
	if *cursor != 6 {
		t.Fatal("cursor pos")
	}
}

func TestDelCursor(t *testing.T) {
	for i := 0; i < 1024; i++ {
		b := New(nil)
		cur := b.AddCursor(0)
		b.DelCursor(cur)
		if b.Cursors.Len() != 0 {
			t.Fatal("cursor amount")
		}
	}
}

func TestCursorsIterate(t *testing.T) {
	for count := 0; count < 1024; count++ {
		cursors := NewCursors()
		for i := 0; i < 16; i++ {
			i := i
			cursors.Add(&i)
		}
		n := 0
		cursors.Iterate(func(cursor *Cursor) bool {
			if cursor.Pos() > 10 {
				return false
			}
			n++
			return true
		})
		if n != 11 {
			t.Fatal("n")
		}
	}
}

func TestDeleteCursorAtSamePos(t *testing.T) {
	for i := 0; i < 1024; i++ {
		cursors := NewCursors()
		c1 := 0
		c2 := 0
		cursors.Add(&c1)
		cursors.Add(&c2)
		cursors.DelCursor(&c1)
		if cursors.Len() != 1 {
			t.Fatal("cursors len")
		}
	}
}

func TestDeleteSamePos(t *testing.T) {
	for i := 0; i < 1024; i++ {
		cursors := NewCursors()
		for n := 0; n < 128; n++ {
			c := 0
			cursors.Add(&c)
		}
		cursors.DelPos(0)
		if cursors.Len() != 0 {
			t.Fatal("cursors len")
		}
	}
}
