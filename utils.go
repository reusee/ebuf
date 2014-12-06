package ebuf

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	pt = fmt.Printf
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Range represent an integer range
type Range struct {
	Begin int
	End   int
}

// RangesSorter sort Ranges by Begin
type RangesSorter []Range

func (r RangesSorter) Len() int           { return len(r) }
func (r RangesSorter) Less(i, j int) bool { return r[i].Begin < r[j].Begin }
func (r RangesSorter) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
