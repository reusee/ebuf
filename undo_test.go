package ebuf

import "testing"

func TestUndo(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte("foobarbaz"))
	b.Undo()
	if b.Current != 0 {
		t.Fatal("current")
	}
	b.Undo()
	if b.Current != 0 {
		t.Fatal("current")
	}
}

func TestRedo(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte("foobarbaz"))
	b.Undo()
	if b.Current != 0 {
		t.Fatal("current")
	}
	b.Redo()
	if b.Current != 2 {
		t.Fatal("current")
	}
	b.Redo()
	if b.Current != 2 {
		t.Fatal("current")
	}
}
