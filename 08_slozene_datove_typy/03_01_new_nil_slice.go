// Řezy v programovacím jazyku Go
//
// - konstrukce takzvaného nulového řezu (nil slice)

package main

import "fmt"

func main() {
	// konstrukce nulového řezu
	var s []int

	// tisk hodnoty a typu řezu
	fmt.Printf("Hodnota: %v\n", s)
	fmt.Printf("Typ:     %T\n", s)
	fmt.Printf("=nil?    %v\n", s == nil)
	fmt.Printf("Len:     %d\n", len(s))
	fmt.Printf("Cap:     %d\n", cap(s))
}
