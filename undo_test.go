package ebuf

import "testing"

func TestUndo(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte("foobarbaz"))
	b.Undo()
	if b.currentIndex() != 0 {
		t.Fatal("current")
	}
	b.Undo()
	if b.currentIndex() != 0 {
		t.Fatal("current")
	}
}

func TestRedo(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte("foobarbaz"))
	b.Undo()
	if b.currentIndex() != 0 {
		t.Fatal("current")
	}
	b.Redo()
	if b.currentIndex() != 1 {
		t.Fatal("current")
	}
	b.Redo()
	if b.currentIndex() != 1 {
		t.Fatal("current")
	}
}

func TestUndoRedoWithEdit(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte("foo"))

	b.Undo()
	b.Insert(0, []byte("bar")) // redo is cleared
	if b.States.Len() != 2 {
		t.Fatal("events length")
	}
	if b.currentIndex() != 1 {
		t.Fatal("current")
	}
	if string(b.Current.Value.(*State).Rope.Bytes()) != "bar" {
		t.Fatal("string of index 2")
	}

	b.Undo()
	if b.States.Len() != 2 {
		t.Fatal("events length")
	}
	b.Delete(0, 0)
	if b.States.Len() != 2 {
		t.Fatal("events length")
	}
}
