package ebuf

import (
	"testing"

	"github.com/reusee/rope"
)

func TestInsert(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte("foobarbaz"))
	if len(b.Events) != 3 {
		t.Fatal("event length")
	}
	if b.Current != 2 {
		t.Fatal("current")
	}
	if string(b.Events[0].(*rope.Rope).Bytes()) != "" {
		t.Fatal("string of index 0")
	}
	if ins := b.Events[1].(OpInsert); ins.Pos != 0 || ins.Len != 9 {
		t.Fatal("insert operation")
	}
	if string(b.Events[2].(*rope.Rope).Bytes()) != "foobarbaz" {
		t.Fatal("string of index 2")
	}
}

func TestDelete(t *testing.T) {
	b := New([]byte("foobarbaz"))
	b.Delete(3, 3)
	if len(b.Events) != 3 {
		t.Fatal("event length")
	}
	if b.Current != 2 {
		t.Fatal("current")
	}
	if string(b.Events[0].(*rope.Rope).Bytes()) != "foobarbaz" {
		t.Fatal("string of index 0")
	}
	if del := b.Events[1].(OpDelete); del.Pos != 3 || del.Len != 3 {
		t.Fatal("delete operation")
	}
	if string(b.Events[2].(*rope.Rope).Bytes()) != "foobaz" {
		t.Fatal("string of index 2")
	}
}
