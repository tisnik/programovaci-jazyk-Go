// Řezy v programovacím jazyku Go
//
// - rozdíl mezi řezem a polem

package main

import "fmt"

func main() {
	// deklarace pole
	a := [10]int{1, 2, 3}

	// deklarace řezu
	s := []int{1, 2, 3}

	fmt.Printf("Typ proměnné 'a':     %T\n", a)
	fmt.Printf("Typ proměnné 's':     %T\n", s)

	fmt.Printf("Hodnota proměnné 'a': %v\n", a)
	fmt.Printf("Hodnota proměnné 's': %v\n", s)
}
