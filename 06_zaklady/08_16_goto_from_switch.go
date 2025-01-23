// Základy programovacího jazyka Go
//
// - výskok z příkazu switch pomocí goto.

package main

import "fmt"

func classify(x int) string {
	switch x {
	case 0:
		return "nula"
	case 2, 4, 6, 8:
		goto SudeCislo
	case 1, 3, 5, 7, 9:
		goto LicheCislo
	default:
		goto JineCislo
	}
JineCislo:
	return "mimo známý rozsah"
SudeCislo:
	return "sudé číslo"
LicheCislo:
	return "liché číslo"
}

func main() {
	for x := 0; x <= 10; x++ {
		fmt.Println(x, classify(x))
	}
}
