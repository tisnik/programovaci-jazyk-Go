// Řezy v programovacím jazyku Go
//
// - konstrukce prázdného řezu pomocí make
// - řez bude mít nenulovou kapacitu

package main

import "fmt"

func main() {
	// konstrukce prázdného řezu
	s := make([]int, 10)

	// tisk hodnoty a typu řezu
	fmt.Printf("Value: %v\n", s)
	fmt.Printf("Type:  %T\n", s)
}
