package ebuf

import "testing"

func TestLineBreaks(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte{'\n'})
	if b.LineBreaks.CursorSet.String() != "[0]" {
		t.Fatal("line breaks")
	}
	b.Insert(0, []byte{'\n'})
	if b.LineBreaks.CursorSet.String() != "[0 1]" {
		t.Fatal("line breaks")
	}
	b.Insert(0, []byte{'\n'})
	if b.LineBreaks.CursorSet.String() != "[0 1 2]" {
		t.Fatal("line breaks")
	}

	b.Undo()
	if b.LineBreaks.CursorSet.String() != "[0 1]" {
		t.Fatal("line breaks")
	}
	b.Undo()
	if b.LineBreaks.CursorSet.String() != "[0]" {
		t.Fatal("line breaks")
	}
	b.Undo()
	if b.LineBreaks.CursorSet.String() != "[]" {
		t.Fatal("line breaks")
	}
}
