// Ukazatele v jazyku Go
//
// - deklarace lokální proměnné typu ukazatel
// - tisk nulové hodnoty této proměnné

package main

import "fmt"

func main() {
	// deklarace lokální proměnné typu ukazatel
	var p *int

	// tisk výchozí hodnoty ukazatele
	fmt.Println(p)
}
