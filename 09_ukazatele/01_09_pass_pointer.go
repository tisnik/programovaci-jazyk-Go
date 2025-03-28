// Ukazatele
//
// - deklarace proměnné typu ukazatel
// - nastavení proměnné přes ukazatel
// - použití typové inference

package main

import "fmt"

func increment(ptr *int) {
	*ptr++
}

func main() {
	x := 42
	fmt.Println("x=", x)

	increment(&x)
	fmt.Println("x=", x)
}
