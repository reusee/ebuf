package ebuf

import "testing"

func TestCursorWithInsert(t *testing.T) {
	b := New(nil)

	cursor := b.AddCursor(0)
	if len(b.Cursors) != 1 {
		t.Fatal("cursors length")
	}
	b.Insert(0, []byte("foobar"))
	if len(cursor.Pos) != 2 {
		t.Fatal("cursor pos length")
	}
	if cursor.Pos[0] != 0 || cursor.Pos[1] != 6 {
		t.Fatal("cursor pos")
	}

	cursor = b.AddCursor(3)
	if len(b.Cursors) != 2 {
		t.Fatal("cursors length")
	}
	b.Insert(6, []byte("baz"))
	if len(cursor.Pos) != 2 {
		t.Fatal("cursor pos length")
	}
	if cursor.Pos[1] != 3 || cursor.Pos[2] != 3 {
		t.Fatal("cursor pos")
	}
}

func TestCursorWithDelete(t *testing.T) {
	b := New([]byte("foobarbaz"))

	cursor := b.AddCursor(-1) // for coverage
	if cursor.Pos[b.Current] != 0 {
		t.Fatal("cursor pos")
	}
	if len(b.Cursors) != 1 {
		t.Fatal("cursors length")
	}
	b.Delete(0, 3)
	if len(cursor.Pos) != 2 {
		t.Fatal("cursor pos length")
	}
	if cursor.Pos[0] != 0 || cursor.Pos[1] != 0 {
		t.Fatal("cursor pos")
	}

	cursor = b.AddCursor(b.States[b.Current].Rope.Len() + 42) // for coverage
	if cursor.Pos[b.Current] != 6 {
		t.Fatal("cursor pos")
	}
	b.Delete(0, 3)
	if len(cursor.Pos) != 2 {
		t.Fatal("cursor pos length")
	}
	if cursor.Pos[b.Current] != 3 {
		t.Fatal("cursor pos")
	}

	cursor = b.AddCursor(2)
	b.Delete(0, 3)
	if cursor.Pos[b.Current] != 0 {
		t.Fatal("cursor pos")
	}
}

func TestCursorWithUndoRedo(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte("foobarbaz"))
	cursor := b.AddCursor(9)
	b.Undo()
	//TODO test cursor pos
	b.Insert(0, []byte("foobar"))
	if cursor.Pos[b.Current] != 6 {
		t.Fatal("cursor pos")
	}
}
