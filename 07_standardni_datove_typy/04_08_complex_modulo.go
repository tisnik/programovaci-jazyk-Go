// Základní datové typy v jazyce Go
//
// - pokus o provedení operace modulo

package main

import "fmt"

func main() {
	// zpracovávaná komplexní čísla
	var a complex64 = 1 + 2i
	var b complex64 = 2 + 2i

	// tisk původních hodnot
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	fmt.Println()

	// tisk výsledku výpočtu modulo
	// (NEKOREKTNÍ OPERACE)
	fmt.Println("a%b=", a%b)
}
