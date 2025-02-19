// Pole v programovacím jazyku Go
//
// - zjištění kapacity pole (počtu prvků)

package main

import "fmt"

func main() {
	var a1 [10]byte
	var a2 [20]int32
	a3 := [30]int32{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}

	// výpis kapacity všech polí
	fmt.Printf("Array 1 capacity: %d\n", cap(a1))
	fmt.Printf("Array 2 capacity: %d\n", cap(a2))
	fmt.Printf("Array 3 capacity: %d\n", cap(a3))
}
