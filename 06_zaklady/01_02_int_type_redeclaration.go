// Základy programovacího jazyka Go
//
// - deklarace datového typu nazvaného "int", který přepíše původní typ "int"
//
// POZOR: tento způsob není doporučený a je dobré vědět, kdy ho použít
//        a především, kdy ho naopak nepoužít.

package main

import "fmt"

// nový datový typ "int", který v tomto balíčku přepíše původní typ
// se stejným jménem
type int = float32

func main() {
	// proměnná nazvaná "pi" je sice typu "int", ovšem na rozdíl od jména
	// typu se jedná o numerickou hodnotu s plovoucí řádovou čárkou
	var pi int

	// nyní můžeme do proměnné přiřadit hodnotu s plovoucí řádovou čárkou
	// i když se (zdánlivě) jedná o proměnnou celočíselného typu
	pi = 3.14
	fmt.Println(pi)
}
