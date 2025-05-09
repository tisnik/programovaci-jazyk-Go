// Pole v programovacím jazyku Go
//
// - zjištění kapacity pole (počtu prvků)

package main

import "fmt"

func main() {
	// trojice polí s různým počtem a typem prvků
	var a1 [10]byte
	var a2 [20]int32
	a3 := [30]int32{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}

	// výpis kapacity všech polí
	fmt.Printf("Kapacita pole číslo 1: %d\n", cap(a1))
	fmt.Printf("Kapacita pole číslo 2: %d\n", cap(a2))
	fmt.Printf("Kapacita pole číslo 3: %d\n", cap(a3))
}
