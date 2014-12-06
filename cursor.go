package ebuf

import "github.com/reusee/rope"

// AddCursor adds one cursor to buffer
func (b *Buffer) AddCursor(pos int) *int {
	if pos < 0 {
		pos = 0
	}
	if l := b.Current.Value.(*State).Rope.Len(); pos > l {
		pos = l
	}
	b.Cursors[&pos] = struct{}{}
	return &pos
}

// DelCursor deletes cursor from buffer
func (b *Buffer) DelCursor(cursor *int) {
	delete(b.Cursors, cursor)
}

// CursorSet represents buffer cursors
type CursorSet map[*int]struct{}

// NewCursorSet create a CursorSet with cursor arguments
func NewCursorSet(cursors ...*int) CursorSet {
	set := CursorSet(make(map[*int]struct{}))
	for _, cursor := range cursors {
		set[cursor] = struct{}{}
	}
	return set
}

// Operate deals with an buffer operation
func (c CursorSet) Operate(op Op, _ *rope.Rope) {
	switch op.Type {
	case Insert:
		for cursor := range c {
			if *cursor >= op.Pos {
				*cursor += op.Len
			}
		}
	case Delete:
		for cursor := range c {
			if *cursor > op.Pos && *cursor < op.Pos+op.Len {
				*cursor = op.Pos
			} else if *cursor >= op.Pos+op.Len {
				*cursor -= op.Len
			}
		}
	}
}
