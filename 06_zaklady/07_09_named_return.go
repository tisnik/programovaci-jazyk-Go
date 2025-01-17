// Základy programovacího jazyka Go
//
// - deklarace funkce s parametry a s návratovými hodnotami
// - pojmenování návratových hodnot
// - parametry stejného typu jsou deklarovány společně

package main

import "fmt"

// Deklarace funkce s parametry a s návratovými hodnotami
func calcSumProd(x, y int) (sum, prod int) {
	sum = x + y
	prod = x * y

	// návrat
	return
}

func main() {
	// zavolání funkce s předáním parametrů
	s, p := calcSumProd(2, 3)
	fmt.Printf("sum  = %d\n", s)
	fmt.Printf("prod = %d\n", p)
}
