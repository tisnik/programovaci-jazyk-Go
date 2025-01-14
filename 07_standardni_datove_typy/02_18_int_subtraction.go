// Základní datové typy v jazyce Go
//
// - rozdíl dvou hodnot typu int

package main

import "fmt"

func main() {
	var a, b int = 1, 2

	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println()

	fmt.Println("a-b > 0:   ", a-b > 0)
	fmt.Println("a-b > 100: ", a-b > 100)
}
