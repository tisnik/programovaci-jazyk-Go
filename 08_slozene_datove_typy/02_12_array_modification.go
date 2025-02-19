// Pole v programovacím jazyku Go
//
// - modifikace pole prvek po prvku

package main

import "fmt"

func main() {
	// deklarace pole a naplnění jeho prvků nulami
	var a [10]int

	fmt.Printf("Original array: %v\n", a)

	// modifikace pole
	for i := range a {
		a[i] = i * 2
	}

	fmt.Printf("Modified array: %v\n", a)
}
