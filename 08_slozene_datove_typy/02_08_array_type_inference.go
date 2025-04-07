// Pole v programovacím jazyku Go
//
// - deklarace lokální proměnné s polem
// - použití typové inference
// - inicializace pole
// - odvození délky pole překladačem

package main

import "fmt"

func main() {
	// deklarace pole s jeho inicializací
	a1 := [...]int{10, 20, 30, 40, 50}

	// tisk pole
	fmt.Println(a1)
}
