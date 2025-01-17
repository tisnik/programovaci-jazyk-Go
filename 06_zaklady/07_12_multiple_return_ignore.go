// Základy programovacího jazyka Go
//
// - deklarace funkce s parametry a s návratovými hodnotami
// - parametry stejného typu jsou deklarovány společně
// - ignorování jedné návratové hodnoty

package main

import "fmt"

// Deklarace funkce s parametry a s návratovými hodnotami
func calcSumProd(x, y int) (int, int) {
	sum := x + y
	prod := x * y

	// návrat s předáním dvou návratových hodnot
	return sum, prod
}

func main() {
	// zavolání funkce s předáním parametrů
	_, p := calcSumProd(2, 3)
	fmt.Printf("prod  = %d\n", p)
}
