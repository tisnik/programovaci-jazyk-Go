// Základy programovacího jazyka Go
//
// - deklarace uzávěru s parametrem
// - opakované volání uzávěru

package main

import "fmt"

func main() {
	// interní proměnná použitá v uzávěru
	counter := 0

	// interní funkce - uzávěr
	nextValue := func() int {
		// přístup k vnější proměnné
		counter += 1
		return counter
	}

	// opakované zavolání uzávěru
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(nextValue())
}
