// Pole v programovacím jazyku Go
//
// - deklarace lokální konstanty s polem
// - inicializace pole

package main

import "fmt"

func main() {
	var a1 [10]int = [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(a1)
}
