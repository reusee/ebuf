package ebuf

import "testing"

func TestCursorWithInsert(t *testing.T) {
	b := New(nil)

	cursor := b.AddCursor(0)
	if len(b.Cursors) != 1 {
		t.Fatal("cursors length")
	}
	b.Insert(0, []byte("foobar"))
	if *cursor != 6 {
		t.Fatal("cursor pos")
	}

	cursor = b.AddCursor(3)
	if len(b.Cursors) != 2 {
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
	if len(b.Cursors) != 1 {
		t.Fatal("cursors length")
	}
	b.Delete(0, 3)
	if *cursor != 0 {
		t.Fatal("cursor pos")
	}

	cursor = b.AddCursor(b.States[b.Current].Rope.Len() + 42) // for coverage
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
