// Ukazatele v jazyku Go
//
// - deklarace proměnné typu ukazatel
// - nastavení proměnné přes ukazatel

package main

import "fmt"

func main() {
	// běžná proměnná
	var x int = 42

	// ukazatel inicializovaný na adresu proměnné
	var p *int = &x

	// tisk hodnoty proměnné i ukazatele
	fmt.Println("x=", x)
	fmt.Println("p=", p)

	// změna proměnné x nepřímo - přes ukazatel
	*p = -1

	// tisk nové hodnoty proměnné
	fmt.Println("x=", x)
}
