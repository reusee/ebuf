package ebuf

import "testing"

func TestMultiInsert(t *testing.T) {
	b := New([]byte("abc"))
	b.AddCursor(0)
	b.AddCursor(1)
	b.AddCursor(2)
	b.InsertAtCursors([]byte(">"))
	if string(b.CurrentBytes()) != ">a>b>c" {
		t.Fatal("bytes")
	}
	b.Undo()
	if string(b.CurrentBytes()) != "abc" {
		t.Fatal("bytes")
	}
	b.Redo()
	if string(b.CurrentBytes()) != ">a>b>c" {
		t.Fatal("bytes")
	}
	b.InsertAtCursors([]byte("<"))
	if string(b.CurrentBytes()) != "><a><b><c" {
		t.Fatal("bytes")
	}
}

func TestMultiDelete(t *testing.T) {
	b := New([]byte("foobarbaz"))
	b.AddCursor(0)
	b.AddCursor(3)
	b.AddCursor(6)
	b.DeleteAtCursors(6)
	if string(b.CurrentBytes()) != "" {
		t.Fatal("bytes")
	}

	b = New([]byte("foobarbaz"))
	b.AddCursor(3)
	b.AddCursor(6)
	b.AddCursor(9)
	b.DeleteAtCursors(-1)
	if string(b.CurrentBytes()) != "fobaba" {
		t.Fatal("bytes")
	}
}
