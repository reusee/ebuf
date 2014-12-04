package ebuf

import "fmt"

var (
	pt = fmt.Printf
)

type Range struct {
	Begin int
	End   int
}

type RangesSorter []Range

func (r RangesSorter) Len() int           { return len(r) }
func (r RangesSorter) Less(i, j int) bool { return r[i].Begin < r[j].Begin }
func (r RangesSorter) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
