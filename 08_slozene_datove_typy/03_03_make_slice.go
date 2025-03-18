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
	fmt.Printf("Hodnota: %v\n", s)
	fmt.Printf("Typ:     %T\n", s)
	fmt.Printf("=nil?    %v\n", s == nil)
	fmt.Printf("Len:     %d\n", len(s))
	fmt.Printf("Cap:     %d\n", cap(s))
}
