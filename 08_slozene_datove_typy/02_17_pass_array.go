// Pole v programovacím jazyku Go
//
// - předání pole do funkce přes ukazatel (odkaz)
// - modifikace pole v této funkci
// - zjištění, zda se pole předává hodnotou nebo odkazem

package main

import "fmt"

func modifyArray(array *[10]int) {
	for i := 0; i < len(array); i++ {
		array[i] = i * 2
	}
	fmt.Printf("Modifikované pole:  %v\n", array)
}

func main() {
	var a1 [10]int

	fmt.Printf("Původní pole:       %v\n", a1)

	modifyArray(&a1)
	fmt.Printf("Po modifikaci:      %v\n", a1)
}
