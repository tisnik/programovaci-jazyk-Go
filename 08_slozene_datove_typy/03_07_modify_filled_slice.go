// Řezy v programovacím jazyku Go
//
// - konstrukce řezu s jeho inicializací
// - modifikace prvního prvku uloženého v řezu

package main

import "fmt"

func main() {
	// konstrukce řezu s jeho inicializací
	var s []int = []int{1, 2, 3}

	fmt.Printf("Typ:  %T\n", s)

	fmt.Printf("Před modifikací: %v\n", s)

	// modifikace prvku
	s[0] = 100

	fmt.Printf("Po modifikaci:   %v\n", s)
}
