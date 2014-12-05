package main

import (
	"fmt"
	"runtime"

	"github.com/reusee/memstat"

	ebuf ".."
)

func main() {
	b := ebuf.New(nil)
	for i := 0; i < 10*10000; i++ {
		b.Insert(0, []byte("o"))
	}

	runtime.GC()
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	fmt.Printf("%f\n", float64(stat.Alloc)/1000000)

	b.Insert(0, []byte("o"))

	memstat.Print()
}
