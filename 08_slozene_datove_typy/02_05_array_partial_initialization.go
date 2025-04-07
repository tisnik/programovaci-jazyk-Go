// Pole v programovacím jazyku Go
//
// - deklarace lokální proměnné s polem
// - inicializace pole
// - prvky bez explicitně uvedené hodnoty jsou vynulovány

package main

import "fmt"

func main() {
	// deklarace pole s jeho inicializací
	var a1 [10]int = [10]int{10, 20, 30, 40, 50}

	// pole lze vytisknout
	fmt.Println(a1)
}
