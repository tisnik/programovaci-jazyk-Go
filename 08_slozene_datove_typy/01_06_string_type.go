// Řetězce v programovacím jazyce Go
//
// - deklarace proměnné typu řetězec
// - využití typové inference
// - tisk typu proměnné

package main

import "fmt"

func main() {
	// deklarace proměnné typu řetězec
	// využívající typové inference
	message := "Toto je retezec!"

	// tisk typu proměnné
	fmt.Printf("%T\n", message)
}
