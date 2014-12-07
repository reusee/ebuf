package ebuf

import "testing"

func TestLineBreaks(t *testing.T) {
	b := New(nil)
	b.Insert(0, []byte{'\n'})
	if b.LineBreaks.Cursors.String() != "0" {
		t.Fatal("line breaks")
	}
	b.Insert(1, []byte("foo\n"))
	if b.LineBreaks.Cursors.String() != "0 4" {
		t.Fatal("line breaks")
	}
	b.Insert(5, []byte("bar\n"))
	if b.LineBreaks.Cursors.String() != "0 4 8" {
		t.Fatal("line breaks")
	}

	b.Undo()
	if b.LineBreaks.Cursors.String() != "0 4" {
		t.Fatal("line breaks")
	}
	b.Undo()
	if b.LineBreaks.Cursors.String() != "0" {
		t.Fatal("line breaks")
	}
	b.Undo()
	if b.LineBreaks.Cursors.String() != "" {
		t.Fatal("line breaks")
	}
}
