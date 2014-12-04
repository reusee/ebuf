package ebuf

import "sort"

func (b *Buffer) InsertAtCursors(bs []byte) {
	// deduplicate cursors
	cursors := make(map[int]*int)
	for _, cursor := range b.Cursors {
		cursors[*cursor] = cursor
	}
	// insert
	b.Action(func() {
		for _, cursor := range cursors {
			b.Insert(*cursor, bs)
		}
	})
}

func (b *Buffer) DeleteAtCursors(length int) {
	// calculate ranges
	var ranges []Range
	for _, cursor := range b.Cursors {
		end := *cursor + length
		if end > *cursor {
			ranges = append(ranges, Range{*cursor, end})
		} else {
			ranges = append(ranges, Range{end, *cursor})
		}
	}
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
		var cursors []*int
		lengths := make(map[*int]int)
		for _, r := range delRanges {
			pos := r.Begin
			cursors = append(cursors, &pos)
			lengths[&pos] = r.End - pos
		}
		for _, cursor := range cursors {
			b.DeleteWithTempCursors(*cursor, lengths[cursor], cursors)
		}
	})
}
