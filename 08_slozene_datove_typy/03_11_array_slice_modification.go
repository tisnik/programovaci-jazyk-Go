// Řezy v programovacím jazyku Go
//
// - modifikace řezu
// - zjištění, jak se změní pole, na které řez ukazuje

package main

import "fmt"

func main() {
	// pole
	var a [10]int

	// vytvoření řezu z pole
	slice := a[:]

	fmt.Printf("Pole před modifikací: %v\n", a)
	fmt.Printf("Řez před modifikací:  %v\n", slice)
	fmt.Println()

	// modifikace pole
	for i := 0; i < len(a); i++ {
		a[i] = i * 2
	}

	fmt.Printf("Pole po modifikací:   %v\n", a)
	fmt.Printf("Řez po modifikaci:    %v\n", slice)
	fmt.Println()

	// modifikace řezu
	for i := 0; i < len(slice); i++ {
		slice[i] = 42
	}

	fmt.Printf("Pole po modifikací:   %v\n", a)
	fmt.Printf("Řez po modifikaci:    %v\n", slice)
}
