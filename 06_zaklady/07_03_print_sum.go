// Základy programovacího jazyka Go
//
// - deklarace funkce s parametry, ale bez návratové hodnoty
// - parametry stejného typu jsou deklarovány společně

package main

import "fmt"

// Deklarace funkce s parametry, ale bez návratové hodnoty
func printSum(x, y int) {
	sum := x + y
	fmt.Printf("%d + %d = %d\n", x, y, sum)
}

func main() {
	// zavolání funkce s předáním parametrů
	printSum(1, 2)
}
