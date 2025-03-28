// Ukazatele
//
// - deklarace proměnné typu ukazatel
// - nastavení proměnné přes ukazatel
// - použití typové inference

package main

import "fmt"

func main() {
	x := 42
	p := &x

	fmt.Println("x=", x)
	fmt.Println("p=", p)

	*p = -1
	fmt.Println("x=", x)
}
