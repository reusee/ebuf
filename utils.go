package ebuf

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
)

var (
	pt = fmt.Printf
)

func init() {
	var seed int64
	binary.Read(crand.Reader, binary.LittleEndian, &seed)
	rand.Seed(seed)
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
