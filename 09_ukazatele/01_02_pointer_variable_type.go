// Ukazatele v jazyku Go
//
// - deklarace lokální proměnné typu ukazatel
// - tisk typu této proměnné

package main

import "fmt"

func main() {
	// deklarace lokální proměnné typu ukazatel
	var p *int

	// tisk typu hodnoty uložené v proměnné
	fmt.Printf("%T\n", p)
}
