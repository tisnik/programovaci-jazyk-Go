// Pole v programovacím jazyku Go
//
// - předání pole do funkce
// - modifikace pole v této funkci
// - zjištění, zda se pole předává hodnotou nebo odkazem

package main

import "fmt"

func modifyArray(array [10]int) {
	for i := 0; i < len(array); i++ {
		array[i] = i * 2
	}
	fmt.Printf("Modified array:     %v\n", array)
}

func main() {
	var a1 [10]int

	fmt.Printf("Original array:     %v\n", a1)

	modifyArray(a1)
	fmt.Printf("After modification: %v\n", a1)
}
