// Základy programovacího jazyka Go
//
// - deklarace anonymní funkce
// - přiřazení této funkce do globální proměnné
// - zavolání této funkce s předáním argumentů

package main

import (
	"fmt"
)

// Deklarace funkce s parametry, ale bez návratové hodnoty
var printSum = func(x, y int) {
	sum := x + y
	fmt.Printf("%d + %d = %d\n", x, y, sum)
}

func main() {
	// zavolání funkce s předáním argumentů
	printSum(1, 2)
}
