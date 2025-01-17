// Základy programovacího jazyka Go
//
// - deklarace funkce s parametry a s návratovými hodnotami
// - parametry stejného typu jsou deklarovány společně
// - ignorování obou návratových hodnot

package main

// Deklarace funkce s parametry a s návratovými hodnotami
func calcSumProd(x, y int) (int, int) {
	sum := x + y
	prod := x * y

	// návrat s předáním dvou návratových hodnot
	return sum, prod
}

func main() {
	// zavolání funkce s předáním parametrů
	calcSumProd(2, 3)
}
