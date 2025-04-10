// Řezy v programovacím jazyku Go
//
// - konstrukce řezu s jeho inicializací
// - modifikace prvního prvku uloženého v řezu

package main

import "fmt"

func main() {
	// konstrukce řezu s jeho inicializací
	s := []int{1, 2, 3}

	fmt.Printf("Typ:  %T\n", s)

	// výpis obsahu řezu
	fmt.Printf("Před modifikací: %v\n", s)

	// modifikace prvku
	s[0] = 100

	// výpis obsahu řezu
	fmt.Printf("Po modifikaci:   %v\n", s)
}
