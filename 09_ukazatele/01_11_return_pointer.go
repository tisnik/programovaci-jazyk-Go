// Ukazatele v jazyku Go
//
// - funkce vracející ukazatel na lokální proměnnou

package main

import "fmt"

func getPointer() *int {
	// lokální proměnná
	var i int = 42

	// vracíme adresu lokální proměnné
	return &i
}

func main() {
	// získání adresy lokální proměnné
	p := getPointer()

	// přístup na adresu
	fmt.Printf("%v\n", p)

	// přístup na adresu
	fmt.Printf("%d\n", *p)
}
