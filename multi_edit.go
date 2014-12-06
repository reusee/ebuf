package ebuf

import "sort"

// InsertAtCursors inserts bytes at all cursors
func (b *Buffer) InsertAtCursors(bs []byte) {
	// deduplicate cursors
	cursors := make(map[int]*int)
	b.Cursors.Iterate(func(cursor *Cursor) bool {
		cursors[cursor.Pos()] = cursor.Value
		return true
	})
	// insert
	b.Action(func() {
		for _, cursor := range cursors {
			b.Insert(*cursor, bs)
		}
	})
}

// DeleteAtCursors deletes specified lengthed bytes at all cursors
func (b *Buffer) DeleteAtCursors(length int) {
	// calculate ranges
	var ranges []Range
	b.Cursors.Iterate(func(cursor *Cursor) bool {
		end := cursor.Pos() + length
		if end > cursor.Pos() {
			ranges = append(ranges, Range{cursor.Pos(), end})
		} else {
			ranges = append(ranges, Range{end, cursor.Pos()})
		}
		return true
	})
	// merge overlapped ranges
	sort.Sort(RangesSorter(ranges))
	delRanges := []Range{}
	for i, r := range ranges {
		if i == 0 {
			delRanges = append(delRanges, r)
		} else {
			last := delRanges[len(delRanges)-1]
			if r.Begin >= last.Begin && r.Begin <= last.End { // overlapped
				if r.End > last.End {
					last.End = r.End
					delRanges[len(delRanges)-1] = last
				}
			} else {
				delRanges = append(delRanges, r)
			}
		}
	}
	// delete
	b.Action(func() {
		cursors := NewCursors()
		lengths := make(map[*int]int)
		for _, r := range delRanges {
			pos := r.Begin
			cursors.Add(&pos)
			lengths[&pos] = r.End - pos
		}
		cursors.Iterate(func(cursor *Cursor) bool {
			b.DeleteWithWatcher(cursor.Pos(), lengths[cursor.Value], cursors)
			return true
		})
	})
}
