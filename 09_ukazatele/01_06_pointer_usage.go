// Ukazatele
//
// - deklarace proměnné typu ukazatel
// - nastavení proměnné přes ukazatel

package main

import "fmt"

func main() {
	// běžná proměnná
	var x int = 42

	// ukazatel inicializovaný na nil
	var p *int

	// tisk hodnoty proměnné i ukazatele
	fmt.Println("x=", x)
	fmt.Println("p=", p)

	// ukazatel obsahuje adresu proměnnné x
	p = &x
	fmt.Println("p=", p)

	// změna proměnné x nepřímo - přes ukazatel
	*p = -1
	fmt.Println("x=", x)
}
