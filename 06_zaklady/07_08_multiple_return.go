// Základy programovacího jazyka Go
//
// - deklarace funkce s parametry a s návratovými hodnotami
// - parametry stejného typu jsou deklarovány společně

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
	s, p := calcSumProd(2, 3)
	fmt.Printf("sum  = %d\n", s)
	fmt.Printf("prod = %d\n", p)
}
