// Řezy v programovacím jazyku Go
//
// - konstrukce nulového řezu
// - pokus o modifikaci prvního prvku, který neexistuje

package main

import "fmt"

func main() {
	// konstrukce nulového řezu
	var s []int

	fmt.Printf("Typ:  %T\n", s)

	fmt.Printf("Před modifikací: %v\n", s)

	// modifikace prvku, který neexistuje
	s[0] = 100

	fmt.Printf("Po modifikaci:   %v\n", s)
}
