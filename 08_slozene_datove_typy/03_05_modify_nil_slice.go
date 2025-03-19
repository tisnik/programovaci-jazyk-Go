// Řezy v programovacím jazyku Go
package main

import "fmt"

func main() {
	var s []int

	fmt.Printf("Typ:  %T\n", s)

	fmt.Printf("Před modifikací: %v\n", s)

	s[0] = 100

	fmt.Printf("Po modifikaci:   %v\n", s)
}
