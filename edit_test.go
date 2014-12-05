package ebuf

import "testing"

func TestInsert(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte("foobarbaz"))
	if b.States.Len() != 2 {
		t.Fatal("event length")
	}
	if b.currentIndex() != 1 {
		t.Fatal("current")
	}
	if string(b.Current.Prev().Value.(*State).Rope.Bytes()) != "" {
		t.Fatal("string of index 0")
	}
	op := b.Current.Value.(*State).LastOp
	if op.Type != Insert || op.Pos != 0 || op.Len != 9 {
		t.Fatal("insert operation")
	}
	if string(b.Current.Value.(*State).Rope.Bytes()) != "foobarbaz" {
		t.Fatal("string of index 2")
	}
}

func TestDelete(t *testing.T) {
	b := New([]byte("foobarbaz"))
	b.Delete(3, 3)
	if b.States.Len() != 2 {
		t.Fatal("event length")
	}
	if b.currentIndex() != 1 {
		t.Fatal("current")
	}
	if string(b.Current.Prev().Value.(*State).Rope.Bytes()) != "foobarbaz" {
		t.Fatal("string of index 0")
	}
	op := b.Current.Value.(*State).LastOp
	if op.Type != Delete || op.Pos != 3 || op.Len != 3 {
		t.Fatal("delete operation")
	}
	if string(b.Current.Value.(*State).Rope.Bytes()) != "foobaz" {
		t.Fatal("string of index 2")
	}
}

func TestInsertWithWatcher(t *testing.T) {
	b := New([]byte("foobar"))
	c1 := 0
	c2 := 3
	c3 := 6
	cursors := NewCursorSet(&c1, &c2, &c3)
	b.InsertWithWatcher(3, []byte("baz"), cursors)
	if c1 != 0 || c2 != 6 || c3 != 9 {
		t.Fatal("cursor pos")
	}
}

func TestDeleteWithWatcher(t *testing.T) {
	b := New([]byte("foobarbaz"))
	c1 := 0
	c2 := 3
	c3 := 6
	cursors := NewCursorSet(&c1, &c2, &c3)
	b.DeleteWithWatcher(3, 6, cursors)
	if c1 != 0 || c2 != 3 || c3 != 3 {
		t.Fatal("cursor pos")
	}
}
