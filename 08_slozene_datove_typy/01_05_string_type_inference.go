// Řetězce v programovacím jazyce Go
//
// - deklarace proměnné typu řetězec
// - využití typové inference

package main

import "fmt"

func main() {
	// deklarace proměnné typu řetězec
	// využívající typové inference
	message := "Toto je retezec!"

	// tisk obsahu této proměnné
	fmt.Println(message)
}
