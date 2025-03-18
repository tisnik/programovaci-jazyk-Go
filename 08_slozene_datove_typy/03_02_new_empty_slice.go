// Řezy v programovacím jazyku Go
//
// - konstrukce prázdného řezu

package main

import "fmt"

func main() {
	// konstrukce prázdného řezu
	s := []int{}

	// tisk hodnoty a typu řezu
	fmt.Printf("Hodnota: %v\n", s)
	fmt.Printf("Typ:     %T\n", s)
	fmt.Printf("=nil?    %v\n", s == nil)
}
