// Ukazatele v jazyku Go
//
// - deklarace proměnné typu ukazatel
// - nastavení proměnné přes ukazatel
// - použití typové inference

package main

import "fmt"

func main() {
	// běžná proměnná i ukazatel deklarované
	// s využitím typové inference
	x := 42
	p := &x

	fmt.Println("x=", x)
	fmt.Println("p=", p)

	// změna proměnné x nepřímo - přes ukazatel
	*p = -1

	// tisk nové hodnoty proměnné
	fmt.Println("x=", x)
}
