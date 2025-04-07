// Pole v programovacím jazyku Go
//
// - deklarace lokální proměnné s polem
// - inicializace pole
// - počet prvků si zjistí překladač automaticky

package main

import "fmt"

func main() {
	// deklarace lokální proměnné s polem
	// inicializace prvků
	var a1 [10]int = [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// tisk pole
	fmt.Println(a1)
}
