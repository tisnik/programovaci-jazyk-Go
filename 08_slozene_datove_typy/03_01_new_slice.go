// Řezy v programovacím jazyku Go
//
// - konstrukce takzvaného nulového řezu (nil slice)

package main

import "fmt"

func main() {
	// konstrukce nulového řezu
	var s []int

	// tisk hodnoty a typu řezu
	fmt.Printf("Value: %v\n", s)
	fmt.Printf("Type:  %T\n", s)
}
