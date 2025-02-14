// Základy programovacího jazyka Go
//
// - deklarace funkce s parametry a s návratovou hodnotou
// - parametry stejného typu jsou deklarovány společně
// - funkce nevrací žádnou hodnotu, i když by měla

package main

import "fmt"

// Deklarace funkce s parametry a s návratovou hodnotou
func calcSum(x, y int) int {
	sum := x + y
	// zde chybí příkaz return
}

func main() {
	// zavolání funkce s předáním parametrů
	result := calcSum(1, 2)
	fmt.Printf("Result = %d\n", result)
}
