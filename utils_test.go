package ebuf

import (
	"sort"
	"testing"
)

func TestRangesSorter(t *testing.T) {
	ranges := []Range{
		Range{1, 1},
		Range{3, 3},
		Range{2, 2},
	}
	sort.Sort(RangesSorter(ranges))
	if ranges[0].Begin != 1 || ranges[1].Begin != 2 || ranges[2].Begin != 3 {
		t.Fatal("sort")
	}
}
