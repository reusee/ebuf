package ebuf

import "testing"

func TestAction(t *testing.T) {
	b := New(nil)
	cursor := b.AddCursor(0)
	b.Action(func() {
		b.Insert(0, []byte("foo"))
		b.Insert(3, []byte("bar"))
		b.Insert(6, []byte("baz"))
	})
	if b.Current != 3 {
		t.Fatal("current")
	}
	if *cursor != 9 {
		t.Fatal("cursor pos")
	}

	b.Undo()
	if b.Current != 0 {
		t.Fatal("current")
	}
	if *cursor != 0 {
		t.Fatal("cursor pos")
	}

	b.Redo()
	if b.Current != 3 {
		t.Fatal("current")
	}
	if *cursor != 9 {
		t.Fatal("cursor pos")
	}

	b.Action(func() {
		b.Delete(0, 3)
		b.Delete(0, 3)
	})
	if b.Current != 5 {
		t.Fatal("current")
	}
	if *cursor != 3 {
		t.Fatal("cursor pos")
	}

	b.Undo()
	if b.Current != 3 {
		t.Fatal("current")
	}
	if *cursor != 9 {
		t.Fatal("cursor pos")
	}
}
