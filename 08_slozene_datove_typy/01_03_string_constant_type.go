// Řetězce v programovacím jazyce Go
//
// - deklarace konstanty s řetězcem
// - tisk typu této konstanty

package main

import "fmt"

// deklarace konstanty typu řetězec
const message = "Toto je řetězec!"

func main() {
	// tisk typu konstanty
	fmt.Printf("%T\n", message)
}
