// Základy programovacího jazyka Go
//
// - deklarace uzávěru
// - využití uzávěru (jeho zavolání)

package main

import "fmt"

func main() {
	// interní proměnná použitá v uzávěru
	counter := 0

	// interní funkce - uzávěr
	nextValue := func(increment int) int {
		// přístup k vnější proměnné
		counter += increment
		return counter
	}

	// opakované zavolání uzávěru s předáním argumentu
	fmt.Println(nextValue(10))
	fmt.Println(nextValue(10))
	fmt.Println(nextValue(10))
	fmt.Println(nextValue(10))
	fmt.Println(nextValue(10))
}
