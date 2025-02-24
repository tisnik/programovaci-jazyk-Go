// Řezy v programovacím jazyku Go
//
// - konstrukce prázdného řezu pomocí make
// - řez bude mít nulovou kapacitu

package main

import "fmt"

func main() {
	// konstrukce prázdného řezu
	s := make([]int, 0)

	// tisk hodnoty a typu řezu
	fmt.Printf("Value: %v\n", s)
	fmt.Printf("Type:  %T\n", s)
}
