package ebuf

import (
	"testing"

	"github.com/reusee/rope"
)

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

func TestUndoRedoWithEdit(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte("foo"))

	b.Undo()
	b.Insert(0, []byte("bar")) // redo is cleared
	if len(b.Events) != 3 {
		t.Fatal("events length")
	}
	if b.Current != 2 {
		t.Fatal("current")
	}
	if string(b.Events[b.Current].(*rope.Rope).Bytes()) != "bar" {
		t.Fatal("string of index 2")
	}

	b.Undo()
	if len(b.Events) != 3 {
		t.Fatal("events length")
	}
	b.Delete(0, 0)
	if len(b.Events) != 3 {
		t.Fatal("events length")
	}
}
