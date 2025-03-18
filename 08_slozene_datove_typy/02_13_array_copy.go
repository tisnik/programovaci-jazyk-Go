// Pole v programovacím jazyku Go
//
// - kopie pole
// - modifikace pole, ale nikoli jeho kopie

package main

import "fmt"

func main() {
	var a1 [10]int

	// kopie pole
	a2 := a1

	// výpis obsahu obou polí
	fmt.Printf("Pole 1: %v\n", a1)
	fmt.Printf("Pole 2: %v\n", a2)

	// modifikace pole a1
	for i := 0; i < len(a1); i++ {
		a1[i] = i * 2
	}

	fmt.Println("-------------------------------")

	// výpis obsahu obou polí
	fmt.Printf("Pole 1: %v\n", a1)
	fmt.Printf("Pole 2: %v\n", a2)
}
