// Základní datové typy v jazyce Go
//
// - celočíselná hodnota deklarovaná jako lokální konstanta
// - tisk typu této konstanty

package main

import "fmt"

func main() {
	// deklarace lokální celočíselné konstanty
	const Answer = 42

	// tisk typu konstanty
	fmt.Printf("%T\n", Answer)
}
