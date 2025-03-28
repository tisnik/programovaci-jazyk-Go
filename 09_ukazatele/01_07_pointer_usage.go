// Ukazatele
//
// - deklarace proměnné typu ukazatel
// - nastavení proměnné přes ukazatel

package main

import "fmt"

func main() {
	var x int = 42
	var p *int = &x

	fmt.Println("x=", x)
	fmt.Println("p=", p)

	*p = -1
	fmt.Println("x=", x)
}
