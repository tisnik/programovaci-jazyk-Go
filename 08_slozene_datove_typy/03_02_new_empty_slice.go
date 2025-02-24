// Řezy v programovacím jazyku Go
//
// - konstrukce prázdného řezu

package main

import "fmt"

func main() {
	// konstrukce prázdného řezu
	s := []int{}

	// tisk hodnoty a typu řezu
	fmt.Printf("Value: %v\n", s)
	fmt.Printf("Type:  %T\n", s)
}
