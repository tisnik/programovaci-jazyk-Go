// Pole v programovacím jazyku Go
//
// - deklarace lokální proměnné s polem
// - použití typové inference
// - inicializace pole

package main

import "fmt"

func main() {
	// deklarace pole s jeho inicializací
	a1 := [10]int{10, 20, 30, 40, 50}

	// tisk pole
	fmt.Println(a1)
}
