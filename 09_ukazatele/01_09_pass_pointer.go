// Ukazatele v jazyku Go
//
// - deklarace proměnné typu ukazatel
// - nastavení proměnné přes ukazatel
// - použití typové inference

package main

import "fmt"

// funkce, která modifikuje obsah proměnné přes ukazatel
func increment(ptr *int) {
	// nepřímá změna obsahu proměnné
	*ptr++
}

func main() {
	// deklarace proměnné s její inicializací
	x := 42

	// tisk původního obsahu proměnné
	fmt.Println("x=", x)

	// modifikace proměnné nepřímo přes ukazatel
	increment(&x)

	// tisk nového obsahu proměnné
	fmt.Println("x=", x)
}
