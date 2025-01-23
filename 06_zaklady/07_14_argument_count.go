// Základy programovacího jazyka Go
//
// - deklarace funkce s parametry, ale bez návratové hodnoty
// - parametry stejného typu jsou deklarovány společně
// - funkce je volána s nekorektními počty argumentů

package main

import "fmt"

// Deklarace funkce s parametry, ale bez návratové hodnoty
func printSum(x, y int) {
	sum := x + y
	fmt.Printf("%d + %d = %d\n", x, y, sum)
}

func main() {
	// zavolání funkce s předáním menšího počtu argumentů,
	// než se očekává z hlavičky funkce
	printSum(1)

	// zavolání funkce s předáním většího počtu argumentů,
	// než se očekává z hlavičky funkce
	printSum(1, 2, 3)

	// zavolání funkce bez předání argumentů
	printSum()
}
